import type { Action, TestFileView, TestFileGraph } from '../domain/types';
import {
  addNodeToGraph,
  deleteNodeFromGraph,
  dropSynthesisedEdges,
  generateNodeId,
  getNextNodes,
  toEngineTasks,
} from '../modules/graphs';
import { FileService } from '../services/file-service';
import { GraphValidator, type DiagnosticReport } from '../domain/graph-validator';
import { projectStore } from './project-store.svelte';
import { applicationStore } from './application-store.svelte';
import { executionStore } from './execution-store.svelte';

const WRITE_DEBOUNCE_MS = 250;

function isTerminal(graph: TestFileView, id: string): boolean {
  for (const edge of graph.graph.edges) {
    if (edge.from === id && !edge.implicit) return false;
  }
  return graph.graph.nodes[id] !== undefined && id !== '__start__';
}

class EditorStore {
  // ---- File selection ----
  currentFile = $state<TestFileView | null>(null);
  selectedFilePath = $state<string | null>(null);
  selectedNodeId = $state<string | null>(null);

  // ---- Transactional Persistence & Dirty Tracking ----
  isDirty = $state(false);
  dirtyNeedsCompile = $state(false);
  isSaving = $state(false);
  private saveTimer: ReturnType<typeof setTimeout> | null = null;
  private pendingWritePromise: Promise<void> | null = null;

  // ---- Reactive Diagnostics (FOUNDATION-03) ----
  diagnostics = $derived.by<DiagnosticReport>(() => {
    return GraphValidator.validate(this.currentFile?.graph);
  });

  selectedNode = $derived.by(() => {
    if (!this.currentFile || !this.selectedNodeId) return null;
    return this.currentFile.graph.nodes[this.selectedNodeId] ?? null;
  });

  needsCompile = $derived.by(() => {
    if (!this.currentFile) return false;
    if (this.dirtyNeedsCompile) return true;
    if (!this.diagnostics.isValid) return true;
    const nodes = Object.values(this.currentFile.graph.nodes);
    if (nodes.length === 0) return false;
    return nodes.some((n) => n.id !== '__start__' && (!n.source || n.source.trim() === ''));
  });

  allPaths = $derived.by(() => {
    if (!this.currentFile) return [];
    return buildGraphPathsFromFrontend(this.currentFile.graph);
  });

  // ---- File lifecycle ----
  async loadFile(path: string, file: TestFileView) {
    await this.flushPendingSave();

    this.selectedFilePath = path;
    this.currentFile = file;
    this.selectedNodeId = null;
    this.isDirty = false;
    this.dirtyNeedsCompile = false;
    executionStore.clearStates();
  }

  async closeFile() {
    await this.flushPendingSave();
    this.selectedFilePath = null;
    this.currentFile = null;
    this.selectedNodeId = null;
    this.isDirty = false;
    this.dirtyNeedsCompile = false;
  }

  // ---- Mutations ----
  updateNode(id: string, updates: Partial<Action>): void {
    if (!this.currentFile) return;
    const node = this.currentFile.graph.nodes[id];
    if (!node) return;
    this.isDirty = true;
    this.dirtyNeedsCompile = true;
    this.currentFile = {
      ...this.currentFile,
      graph: {
        ...this.currentFile.graph,
        nodes: {
          ...this.currentFile.graph.nodes,
          [id]: { ...node, ...updates },
        },
      },
    };
    this.scheduleSave();
  }

  addNode(fromId: string | null, toId: string | null): string | null {
    if (!this.currentFile) return null;
    this.isDirty = true;
    this.dirtyNeedsCompile = true;
    const { graph, newNodeId } = addNodeToGraph(this.currentFile.graph, fromId, toId);
    this.currentFile = { ...this.currentFile, graph };
    this.selectedNodeId = newNodeId;
    void this.saveFile();
    return newNodeId;
  }

  deleteNode(id: string) {
    if (!this.currentFile) return;
    const file = this.currentFile;
    applicationStore.openConfirmModal(
      'Delete Action Node',
      'Delete this action node and everything after it?',
      () => {
        if (!file) return;
        this.isDirty = true;
        this.dirtyNeedsCompile = true;
        const after = deleteNodeFromGraph(file.graph, id);
        this.currentFile = { ...file, graph: after };
        if (this.selectedNodeId === id) this.selectedNodeId = null;
        void this.saveFile();
      },
    );
  }

  // ---- File ops ----
  async createFile(dir: string, name: string): Promise<void> {
    if (!projectStore.path) return;
    const filename = name.endsWith('.test.yml') ? name : `${name}.test.yml`;
    const relPath = joinRel(dir, filename);
    try {
      await FileService.createFile(joinAbs(projectStore.path, relPath));
    } catch (e) {
      applicationStore.showToast('error', `Could not create ${relPath}: ${errorMessage(e)}`);
      return;
    }
    await projectStore.refreshTests();
    await this.loadFileFromDisk(relPath);
  }

  async createDirectory(path: string): Promise<void> {
    if (!projectStore.path) return;
    try {
      await FileService.createDirectory(joinAbs(projectStore.path, path));
    } catch (e) {
      applicationStore.showToast('error', `Could not create ${path}: ${errorMessage(e)}`);
      return;
    }
    await projectStore.refreshTests();
  }

  async deletePath(path: string): Promise<void> {
    if (!projectStore.path) return;
    const label = path.endsWith('.test.yml') ? 'test' : 'folder';
    applicationStore.openConfirmModal(
      `Delete ${label}`,
      `Delete ${path}? This can't be undone.`,
      async () => {
        try {
          await FileService.deletePath(joinAbs(projectStore.path!, path));
        } catch (e) {
          applicationStore.showToast('error', `Could not delete ${path}: ${errorMessage(e)}`);
          return;
        }
        const closedHere =
          this.selectedFilePath === path ||
          (label === 'folder' &&
            this.selectedFilePath !== null &&
            this.selectedFilePath.startsWith(path));
        if (closedHere) await this.closeFile();
        await projectStore.refreshTests();
      },
    );
  }

  // ---- Transactional Persistence ----
  scheduleSave() {
    if (this.saveTimer !== null) clearTimeout(this.saveTimer);
    this.saveTimer = setTimeout(() => {
      this.saveTimer = null;
      void this.saveFile();
    }, WRITE_DEBOUNCE_MS);
  }

  async flushPendingSave(): Promise<void> {
    if (this.saveTimer !== null) {
      clearTimeout(this.saveTimer);
      this.saveTimer = null;
      await this.saveFile();
    } else if (this.pendingWritePromise) {
      await this.pendingWritePromise;
    }
  }

  async saveFile(): Promise<void> {
    if (this.saveTimer !== null) {
      clearTimeout(this.saveTimer);
      this.saveTimer = null;
    }
    if (!projectStore.path || !this.selectedFilePath || !this.currentFile) return;

    const view: TestFileView = {
      ...this.currentFile,
      graph: {
        ...this.currentFile.graph,
        edges: dropSynthesisedEdges(this.currentFile.graph).edges,
      },
      order: this.currentFile.order,
    };

    const writeOp = async () => {
      this.isSaving = true;
      try {
        await FileService.writeTestFile(projectStore.path!, this.selectedFilePath!, view);
        this.isDirty = false;
      } catch (e) {
        applicationStore.showToast('error', `Could not save: ${errorMessage(e)}`);
      } finally {
        this.isSaving = false;
      }
    };

    this.pendingWritePromise = writeOp();
    await this.pendingWritePromise;
    this.pendingWritePromise = null;
  }

  async loadFileFromDisk(relPath: string): Promise<void> {
    if (!projectStore.path) return;
    try {
      const view = await FileService.readTestFile(projectStore.path, relPath);
      await this.loadFile(relPath, view);
    } catch (e) {
      applicationStore.showToast('error', `Could not open ${relPath}: ${errorMessage(e)}`);
    }
  }

  // ---- Execution Delegations ----
  async runCurrent(): Promise<void> {
    if (!projectStore.path || !this.selectedFilePath) return;
    await this.flushPendingSave();
    await executionStore.runCurrent(projectStore.path, this.selectedFilePath);
  }

  async stopRun(): Promise<void> {
    await executionStore.stopRun();
  }

  async compileCurrent(): Promise<void> {
    if (!projectStore.path || !this.selectedFilePath) return;
    await this.flushPendingSave();
    await executionStore.compileCurrent(projectStore.path, this.selectedFilePath, async () => {
      this.dirtyNeedsCompile = false;
      if (this.selectedFilePath && projectStore.path) {
        try {
          const view = await FileService.readTestFile(projectStore.path, this.selectedFilePath);
          this.currentFile = { graph: view.graph, order: view.order ?? [] };
        } catch {
          // Re-read failed gracefully
        }
      }
    });
  }

  async stopCompile(): Promise<void> {
    await executionStore.stopCompile();
  }

  _generateId(): string {
    return generateNodeId();
  }

  _isTerminal(graph: TestFileView, id: string): boolean {
    return isTerminal(graph, id);
  }
}

function buildGraphPathsFromFrontend(graph: TestFileGraph): { tasks: string[] }[] {
  const startId = graph.start;
  if (!startId || !graph.nodes[startId]) return [];
  const tasks = toEngineTasks(graph);

  type Frame = { id: string; tasks: string[] };
  const stack: Frame[] = [{ id: startId, tasks: [startId] }];
  const completed: { tasks: string[] }[] = [];

  while (stack.length > 0) {
    const frame = stack.pop()!;
    const nextIds = getNextNodes(graph, frame.id);
    if (nextIds.length === 0) {
      completed.push({ tasks: frame.tasks });
      continue;
    }
    for (const nextId of nextIds) {
      if (tasks[nextId]) {
        stack.push({ id: nextId, tasks: [...frame.tasks, nextId] });
      } else {
        completed.push({ tasks: [...frame.tasks, nextId] });
      }
    }
  }

  return completed.reverse();
}

function joinRel(parent: string, child: string): string {
  if (parent === '') return child;
  return parent.endsWith('/') ? parent + child : `${parent}/${child}`;
}

function joinAbs(root: string, rel: string): string {
  if (rel.startsWith('/')) return rel;
  return `${root.replace(/\/$/, '')}/${rel}`;
}

function errorMessage(e: unknown): string {
  if (e instanceof Error) return e.message;
  if (typeof e === 'string') return e;
  try {
    return JSON.stringify(e);
  } catch {
    return String(e);
  }
}

export const editorStore = new EditorStore();
