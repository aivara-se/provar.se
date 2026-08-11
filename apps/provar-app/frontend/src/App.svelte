<script lang="ts">
  import { projectStore } from './lib/stores/project-store.svelte';
  import { settingsStore } from './lib/stores/settings-store.svelte';
  import { historyStore } from './lib/stores/history-store.svelte';
  import { editorStore } from './lib/stores/editor-store.svelte';
  import { applicationStore } from './lib/stores/application-store.svelte';
  import { Watcher } from './lib/api';
  import { subscribe } from './lib/events';
  import Welcome from './lib/components/Welcome.svelte';
  import SetupWizard from './lib/components/SetupWizard.svelte';
  import Toolbar from './lib/components/Toolbar.svelte';
  import TestExplorer from './lib/components/TestExplorer.svelte';
  import Canvas from './lib/components/Canvas.svelte';
  import RightSidebar from './lib/components/RightSidebar.svelte';
  import AppModals from './lib/components/AppModals.svelte';

  // Both stores load on first mount; gated on "no project open yet"
  // so the load is skipped once a project is in flight. historyStore.load
  // is also what drives settingsStore.showSetupWizard (a missing
  // history file = first launch = show the wizard).
  $effect(() => {
    if (!projectStore.path) {
      settingsStore.load();
      historyStore.load();
    }
  });

  // Re-arm the file watcher whenever the project path flips. The
  // binding disposes its previous watcher before installing the new
  // one, so we don't leak goroutines on project open/close cycles.
  $effect(() => {
    const path = projectStore.path;
    if (path) {
      void Watcher.Watch(path).catch((e) => {
        console.error('Watcher.Watch failed:', e);
      });
    }
  });

  // The fsnotify events flow through this subscription. We refresh
  // the test list only — the open file is re-read by editorStore.loadFile
  // because the canvas re-reads on every effect run that opens a file,
  // but the user's currently-open file is stable across fs events; we
  // only rescan the directory.
  $effect(() => {
    if (!projectStore.path) return;
    let cancelled = false;
    void (async () => {
      for await (const _ of subscribe('project:changed')) {
        if (cancelled) return;
        await projectStore.refreshTests();
      }
    })();
    return () => {
      cancelled = true;
    };
  });

  // Workspace effects — kept at the root so they live for the lifetime
  // of the app. Each rule is intentionally narrow:
  //
  // 1. Selecting a node opens the right sidebar on the node panel.
  //    The user almost always wants to see the node side panel when
  //    they pick something; if they really don't, they can click the
  //    sidebar toggle to dismiss it.
  $effect(() => {
    if (editorStore.selectedNodeId !== null) {
      applicationStore.isRightSidebarOpen = true;
    }
  });

  // 2. Opening a file auto-hides the test explorer. The toolbar's
  //    sidebar toggle brings it back. Closing the file restores it
  //    so the workspace is back in "browse" mode.
  $effect(() => {
    applicationStore.isSidebarOpen = editorStore.selectedFilePath === null;
  });
</script>

<div
  class="relative h-screen w-full overflow-hidden overscroll-none bg-[rgba(14,17,22,0.7)] font-sans text-zinc-300"
>
  <div
    class="absolute top-0 right-0 left-0 z-40 h-[56px]"
    style="--wails-draggable:drag"
  ></div>

  {#if !projectStore.path}
    {#if settingsStore.showSetupWizard}
      <SetupWizard />
    {:else}
      <Welcome
        homeDir={settingsStore.homeDir}
        recentProjects={historyStore.recent}
        onOpen={(path) => {
          projectStore.openProject(path);
        }}
      />
    {/if}
  {:else}
    <Toolbar />
    <TestExplorer />
    <Canvas />
    <RightSidebar />
  {/if}

  <AppModals />
</div>
