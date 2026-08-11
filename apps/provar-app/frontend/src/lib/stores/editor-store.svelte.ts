import type { Action, TestFileView } from '../domain/types';
import {
  addNodeToGraph,
  deleteNodeFromGraph,
  dropSynthesisedEdges,
  generateNodeId,
  getNextNodes,
  toEngineTasks,
} from '../modules/graphs';
import { File, Run, Compile } from '../services/bindings';
import { projectStore } from './project-store.svelte';
import { applicationStore } from './application-store.svelte';
import { forJob } from '../services/events';

// WRITE_DEBOUNCE_MS is the quiet-period before a typed title/info change
// is committed back to disk. Per-keystroke writes used to fire fresh
// refreshFiles sweeps on every character (see old editor-store docs);
// debouncing keeps the click stream — and the round-trip cost — to one
// write per typing burst.
const WRITE_DEBOUNCE_MS = 250;

// Edge.Synthesised distance for "auto-show End on the canvas". The
// engine never persists an End node — it's rendered virtually based on
// graph structure — so this lives in the store, not on the wire type.
function isTerminal(graph: TestFileView, id: string): boolean {
  for (const edge of graph.graph.edges) {
    if (edge.from === id && !edge.implicit) return false;
  }
  return graph.graph.nodes[id] !== undefined && id !== '__start__';
}

type RunState = 'idle' | 'running' | 'success' | 'failed';
type CompileState = 'idle' | 'compiling' | 'compiled' | 'failed';

class EditorStore {
  // ---- File selection ----
  currentFile = $state<TestFileView | null>(null);
  selectedFilePath = $state<string | null>(null);
  selectedNodeId = $state<string | null>(null);

  selectedNode = $derived.by(() => {
    if (!this.currentFile || !this.selectedNodeId) return null;
    return this.currentFile.graph.nodes[this.selectedNodeId] ?? null;
  });

  // ---- Run / compile state ----
  // runState and compileState aggregate across nodes for compact display
  // in the toolbar; taskStates holds the per-node resolution so the
  // canvas can colour individual shapes.
  isRunning = $state(false);
  activeRunId = $state<string | null>(null);
  isCompiling = $state(false);
  activeCompileId = $state<string | null>(null);
  dirtyNeedsCompile = $state(false);

  taskStates = $state<Record<string, RunState>>({});
  compileStates = $state<Record<string, CompileState>>({});

  needsCompile = $derived.by(() => {
    if (!this.currentFile) return false;
    if (this.dirtyNeedsCompile) return true;
    const nodes = Object.values(this.currentFile.graph.nodes);
    if (nodes.length === 0) return false;
    return nodes.some((n) => n.id !== '__start__' && (!n.source || n.source.trim() === ''));
  });

  allPaths = $derived.by(() => {
    if (!this.currentFile) return [];
    return buildGraphPathsFromFrontend(this.currentFile.graph);
  });

  // ---- File lifecycle ----
  loadFile(path: string, file: TestFileView) {
    this.selectedFilePath = path;
    this.currentFile = file;
    this.selectedNodeId = null;
    this.dirtyNeedsCompile = false;
    // Reset transient state so a stale compile border doesn't linger from
    // a different file the user opened earlier in the session.
    this.taskStates = {};
    this.compileStates = {};
  }

  closeFile() {
    this.selectedFilePath = null;
    this.currentFile = null;
    this.selectedNodeId = null;
    this.dirtyNeedsCompile = false;
  }

  // ---- Mutations ----
  /**
   * updateNode merges the given updates into the node. The change is
   * persisted via a debounced saveFile so a typing burst of keystrokes
   * only lands a single disk write. The pending timer is captured so a
   * subsequent mutation (or the panel's blur flush) can force an
   * immediate commit.
   */
  updateNode(id: string, updates: Partial<Action>): void {
    if (!this.currentFile) return;
    const node = this.currentFile.graph.nodes[id];
    if (!node) return;
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

  /**
   * addNode splices a new task between from and to (either may be null)
   * and selects the new node. Persists once via saveFile; the canvas
   * rebuild picks up the change through the GraphRenderer's reactive
   * bindings.
   */
  addNode(fromId: string | null, toId: string | null): string | null {
    if (!this.currentFile) return null;
    this.dirtyNeedsCompile = true;
    const { graph, newNodeId } = addNodeToGraph(this.currentFile.graph, fromId, toId);
    this.currentFile = { ...this.currentFile, graph };
    this.selectedNodeId = newNodeId;
    void this.saveFile();
    return newNodeId;
  }

  /**
   * deleteNode opens the global confirm modal and removes the node (and
   * its downstream descendants) on confirm. The confirm wiring lives
   * in applicationStore.openConfirmModal — we just hand it a callback that does
   * the mutation.
   */
  deleteNode(id: string) {
    if (!this.currentFile) return;
    const file = this.currentFile;
    applicationStore.openConfirmModal(
      'Delete Action Node',
      'Delete this action node and everything after it?',
      () => {
        if (!file) return;
        this.dirtyNeedsCompile = true;
        // `file` is captured by closure but Svelte 5's $state proxies
        // are reactive in place — so the closure reflects current values.
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
      await File.CreateFile(joinAbs(projectStore.path, relPath));
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
      await File.CreateDirectory(joinAbs(projectStore.path, path));
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
          await File.DeletePath(joinAbs(projectStore.path!, path));
        } catch (e) {
          applicationStore.showToast('error', `Could not delete ${path}: ${errorMessage(e)}`);
          return;
        }
        const closedHere =
          this.selectedFilePath === path ||
          (label === 'folder' &&
            this.selectedFilePath !== null &&
            this.selectedFilePath.startsWith(path));
        if (closedHere) this.closeFile();
        await projectStore.refreshTests();
      },
    );
  }

  // ---- Persistence ----
  private saveTimer: ReturnType<typeof setTimeout> | null = null;

  scheduleSave() {
    if (this.saveTimer !== null) clearTimeout(this.saveTimer);
    this.saveTimer = setTimeout(() => {
      this.saveTimer = null;
      void this.saveFile();
    }, WRITE_DEBOUNCE_MS);
  }

  /** saveFile writes currentFile to disk. Drops synthesised edges so the
   * YAML stays clean — the FromActions synthesis happens again on next
   * load, naturally.
   */
  async saveFile(): Promise<void> {
    if (this.saveTimer !== null) {
      clearTimeout(this.saveTimer);
      this.saveTimer = null;
    }
    if (!projectStore.path) return;
    if (!this.selectedFilePath || !this.currentFile) return;
    const view: TestFileView = {
      ...this.currentFile,
      graph: {
        ...this.currentFile.graph,
        edges: dropSynthesisedEdges(this.currentFile.graph).edges,
      },
      order: this.currentFile.order,
    };
    try {
      await File.WriteTestFile(projectStore.path, this.selectedFilePath, view);
    } catch (e) {
      applicationStore.showToast('error', `Could not save: ${errorMessage(e)}`);
    }
  }

  private async loadFileFromDisk(relPath: string): Promise<void> {
    if (!projectStore.path) return;
    try {
      const view = await File.ReadTestFile(projectStore.path, relPath);
      this.loadFile(relPath, {
        graph: view.graph,
        order: view.order ?? [],
      });
    } catch (e) {
      applicationStore.showToast('error', `Could not open ${relPath}: ${errorMessage(e)}`);
    }
  }

  // ---- Run / compile streams ----
  // The two streams share the lifecycle shape — emit a start, then a
  // sequence of per-node events, then a terminal. EditorStore keeps
  // each stream's job-id so stale events from a previous run don't
  // poison the current one (the engine's Stop races the cancel, and
  // sometimes a late event slips through after we cancel).

  async runCurrent(): Promise<void> {
    if (!projectStore.path || !this.selectedFilePath) return;
    if (this.isRunning) return;
    this.taskStates = {};
    try {
      const jobId = await Run.Start(projectStore.path, this.selectedFilePath, true, '');
      this.activeRunId = jobId;
      this.isRunning = true;
      void this.consumeRunEvents(jobId);
    } catch (e) {
      applicationStore.showToast('error', `Could not start run: ${errorMessage(e)}`);
    }
  }

  async stopRun(): Promise<void> {
    const jobId = this.activeRunId;
    if (!jobId) return;
    try {
      await Run.Cancel(jobId);
    } catch {
      // Already finished — fine.
    }
  }

  private async consumeRunEvents(jobId: string): Promise<void> {
    for await (const event of forJob<{ jobId: string; type: string; data?: unknown }>(jobId, 'job:event')) {
      if (event.type === 'task-started' || event.type === 'task-finished' || event.type === 'task-failed') {
        const payload = event.data as { actionId?: string } | undefined;
        const actionId = payload?.actionId;
        if (!actionId) continue;
        const state: RunState =
          event.type === 'task-started'
            ? 'running'
            : event.type === 'task-finished'
              ? 'success'
              : 'failed';
        this.taskStates = { ...this.taskStates, [actionId]: state };
      }
      if (event.type === 'run-finished') {
        this.isRunning = false;
        this.activeRunId = null;
      }
    }
  }

  async compileCurrent(): Promise<void> {
    if (!projectStore.path || !this.selectedFilePath) return;
    if (this.isCompiling) return;
    this.compileStates = {};
    try {
      const jobId = await Compile.Start(projectStore.path, this.selectedFilePath);
      this.activeCompileId = jobId;
      this.isCompiling = true;
      void this.consumeCompileEvents(jobId);
    } catch (e) {
      applicationStore.showToast('error', `Could not start compile: ${errorMessage(e)}`);
    }
  }

  async stopCompile(): Promise<void> {
    const jobId = this.activeCompileId;
    if (!jobId) return;
    try {
      await Compile.Cancel(jobId);
    } catch {
      // Already finished — fine.
    }
  }

  private async consumeCompileEvents(jobId: string): Promise<void> {
    for await (const event of forJob<{ jobId: string; type: string; data?: unknown }>(jobId, 'job:event')) {
      if (
        event.type === 'action-started' ||
        event.type === 'action-finished' ||
        event.type === 'action-failed'
      ) {
        const payload = event.data as { ActionID?: string; actionId?: string } | undefined;
        const actionId = payload?.ActionID ?? payload?.actionId;
        if (!actionId) continue;
        const state: CompileState =
          event.type === 'action-started'
            ? 'compiling'
            : event.type === 'action-finished'
              ? 'compiled'
              : 'failed';
        this.compileStates = { ...this.compileStates, [actionId]: state };
      }
      if (event.type === 'compile-finished') {
        this.isCompiling = false;
        this.activeCompileId = null;
        this.dirtyNeedsCompile = false;
        // Compile finished → the .test.lua on disk now reflects the
        // latest yaml. Re-read the file so node.source picks up the
        // freshly emitted bodies; without this the side panel's
        // generated-code view would lag until a manual reload.
        const path = this.selectedFilePath;
        if (path && projectStore.path) {
          try {
            const view = await File.ReadTestFile(projectStore.path, path);
            this.currentFile = { graph: view.graph, order: view.order ?? [] };
          } catch {
            // Compile may have failed before writing the lua — surface
            // the failure via the existing per-node red border.
          }
        }
      }
    }
  }

  // Exposed for tests that need a synthetic snapshot.
  _generateId(): string {
    return generateNodeId();
  }

  _isTerminal(graph: TestFileView, id: string): boolean {
    return isTerminal(graph, id);
  }
}

/**
 * buildGraphPathsFromFrontend is a JS port of libs/engine's buildGraphPaths
 * — it walks the graph and produces the per-start-node paths the canvas
 * needs to count executions and find which path contains a given task.
 * We don't import from @libs/engine here because the editor bundle is a
 * browser context and that module pulls domain Go-only plumbing.
 */
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

  // Preserve source order — pop gives reverse DFS so reverse once.
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
