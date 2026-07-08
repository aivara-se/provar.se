<script lang="ts">
  import { ChevronDown, File, Folder, Search } from 'lucide-svelte';
  import { projectStore } from '../stores/project-store.svelte';
  import { editorStore } from '../stores/editor-store.svelte';
  import { uiStore } from '../stores/ui-store.svelte';
  import { File as FileApi } from '../api';
  import type { TestFileView } from '../types';

  type TreeNode = {
    type: 'folder' | 'file';
    name: string;
    path: string;
    children?: TreeNode[];
  };

  type ContextMenuKind = 'file' | 'folder';
  let contextMenu = $state<{
    x: number;
    y: number;
    path: string;
    kind: ContextMenuKind;
  } | null>(null);

  let query = $state('');
  let closedFolders = $state<Set<string>>(new Set());

  let tree = $derived.by(() => {
    const filtered = query
      ? projectStore.tests.filter((p) =>
          p.toLowerCase().includes(query.toLowerCase()),
        )
      : projectStore.tests;
    const root: TreeNode[] = [];
    for (const file of filtered) {
      const parts = file.split('/');
      let level = root;
      let currentPath = '';
      for (let i = 0; i < parts.length; i++) {
        const part = parts[i];
        currentPath = currentPath ? `${currentPath}/${part}` : part;
        const isFile = i === parts.length - 1;
        let node = level.find((n) => n.name === part);
        if (!node) {
          node = {
            type: isFile ? 'file' : 'folder',
            name: part,
            path: currentPath,
            ...(isFile ? {} : { children: [] }),
          };
          level.push(node);
        }
        if (!isFile) level = node.children!;
      }
    }
    const sortNodes = (nodes: TreeNode[]) => {
      nodes.sort((a, b) => {
        if (a.type !== b.type) return a.type === 'folder' ? -1 : 1;
        return a.name.localeCompare(b.name);
      });
      for (const n of nodes) if (n.children) sortNodes(n.children);
    };
    sortNodes(root);
    return root;
  });

  function toggleFolder(path: string) {
    const next = new Set(closedFolders);
    if (next.has(path)) next.delete(path);
    else next.add(path);
    closedFolders = next;
  }

  async function selectFile(path: string) {
    if (!projectStore.path) return;
    try {
      const view = await FileApi.ReadTestFile(projectStore.path, path);
      editorStore.loadFile(path, { graph: view.graph, order: view.order ?? [] });
    } catch (e) {
      console.error('TestExplorer: failed to load file', path, e);
    }
  }

  function handleContextMenu(
    e: MouseEvent,
    path: string,
    kind: ContextMenuKind,
  ) {
    e.preventDefault();
    e.stopPropagation();
    contextMenu = { x: e.clientX, y: e.clientY, path, kind };
  }

  function closeContextMenu() {
    contextMenu = null;
  }

  function newFileInFolder(parent: string) {
    uiStore.openInputModal(
      'New Test',
      'Test name (without .test.yml)',
      (name) => {
        closeContextMenu();
        if (!name) return;
        void editorStore.createFile(parent, name);
      },
    );
  }

  function newFolderIn(parent: string) {
    uiStore.openInputModal(
      'New Folder',
      'Folder name',
      (name) => {
        closeContextMenu();
        if (!name) return;
        void editorStore.createDirectory(`${parent}/${name}`);
      },
    );
  }

  function deleteItem(path: string, kind: ContextMenuKind) {
    const label = kind === 'file' ? 'test' : 'folder';
    uiStore.openConfirmModal(
      `Delete ${label}`,
      `Delete ${path}? This can't be undone.`,
      () => {
        closeContextMenu();
        void editorStore.deletePath(path);
      },
    );
  }
</script>

<svelte:window onclick={closeContextMenu} />

{#snippet treeNode(node: TreeNode, depth: number)}
  {#if node.type === 'folder'}
    <button
      type="button"
      class="mx-2 flex w-[calc(100%-1rem)] cursor-pointer items-center rounded py-1 pr-2 text-left text-xs text-zinc-300 hover:bg-[#21262d]"
      style="padding-left: {depth * 14 + 8}px"
      onclick={() => toggleFolder(node.path)}
      oncontextmenu={(e) => handleContextMenu(e, node.path, 'folder')}
    >
      <ChevronDown
        class="mr-1 h-3.5 w-3.5 text-zinc-500 transition-transform {closedFolders.has(
          node.path,
        )
          ? '-rotate-90'
          : ''}"
      />
      <Folder class="mr-2 h-3.5 w-3.5 fill-zinc-400/20 text-zinc-400" />
      <span>{node.name}</span>
    </button>
    {#if !closedFolders.has(node.path)}
      {#each node.children! as child}
        {@render treeNode(child, depth + 1)}
      {/each}
    {/if}
  {:else}
    <button
      type="button"
      class="mx-2 flex w-[calc(100%-1rem)] cursor-pointer items-center gap-2 rounded py-1 pr-2 text-left text-xs hover:text-zinc-200 {editorStore.selectedFilePath ===
      node.path
        ? 'bg-[#21262d] text-zinc-200'
        : 'text-zinc-400'}"
      style="padding-left: {depth * 14 + 26}px"
      onclick={() => selectFile(node.path)}
      oncontextmenu={(e) => handleContextMenu(e, node.path, 'file')}
    >
      <File
        class="h-3.5 w-3.5 shrink-0 {editorStore.selectedFilePath === node.path
          ? 'text-blue-400'
          : 'text-zinc-500'}"
      />
      <span class="truncate">{node.name}</span>
    </button>
  {/if}
{/snippet}

{#if uiStore.isSidebarOpen}
  <aside
    class="absolute top-0 bottom-0 left-0 z-20 flex w-[260px] flex-col border-r border-zinc-800 bg-[#161b22]/50 pt-[36px] backdrop-blur-md"
  >
    <div class="px-3 pt-3 pb-2">
      <div class="relative">
        <Search class="absolute top-2 right-2.5 h-3.5 w-3.5 text-zinc-500" />
        <input
          type="text"
          placeholder="Search files"
          bind:value={query}
          class="w-full rounded-full border border-zinc-700/30 bg-[#21262d] py-1.5 pr-3 pl-3 text-xs text-zinc-300 placeholder-zinc-500 focus:border-zinc-600 focus:ring-1 focus:ring-zinc-600 focus:outline-none"
        />
      </div>
    </div>

    <div class="flex-1 overflow-y-auto">
      <div class="mt-1 pb-4">
        {#if tree.length === 0}
          <p class="px-4 pt-4 text-xs text-zinc-500">
            {projectStore.tests.length === 0
              ? 'No test files in this project yet. Right-click to create one.'
              : 'No files match the search.'}
          </p>
        {/if}
        {#each tree as node}
          {@render treeNode(node, 0)}
        {/each}
      </div>
    </div>
  </aside>
{/if}

<!--
  Context menu floats above the tree. Positioned at the click coordinate
  via fixed top/left so we don't inherit the aside's scrolling.
-->
{#if contextMenu}
  <div
    class="fixed z-[100] min-w-[160px] rounded-lg border border-zinc-800 bg-[#161b22] p-1 shadow-xl"
    style="top: {contextMenu.y}px; left: {contextMenu.x}px"
  >
    {#if contextMenu.kind === 'folder'}
      <button
        type="button"
        class="flex w-full rounded px-4 py-2 text-left text-xs text-zinc-300 hover:bg-[#21262d]"
        onclick={() => newFileInFolder(contextMenu!.path)}
      >
        New test
      </button>
      <button
        type="button"
        class="flex w-full rounded px-4 py-2 text-left text-xs text-zinc-300 hover:bg-[#21262d]"
        onclick={() => newFolderIn(contextMenu!.path)}
      >
        New folder
      </button>
      <div class="my-1 border-t border-zinc-800"></div>
      <button
        type="button"
        class="flex w-full rounded px-4 py-2 text-left text-xs text-red-400 hover:bg-[#21262d]"
        onclick={() => deleteItem(contextMenu!.path, 'folder')}
      >
        Delete folder
      </button>
    {:else}
      <button
        type="button"
        class="flex w-full rounded px-4 py-2 text-left text-xs text-red-400 hover:bg-[#21262d]"
        onclick={() => deleteItem(contextMenu!.path, 'file')}
      >
        Delete test
      </button>
    {/if}
  </div>
{/if}
