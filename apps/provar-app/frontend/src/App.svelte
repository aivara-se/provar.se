<script lang="ts">
  import { projectStore } from './lib/stores/project-store.svelte';
  import { settingsStore } from './lib/stores/settings-store.svelte';
  import { historyStore } from './lib/stores/history-store.svelte';
  import { editorStore } from './lib/stores/editor-store.svelte';
  import { applicationStore } from './lib/stores/application-store.svelte';
  import { ProjectService } from './lib/services/project-service';
  import { subscribe } from './lib/services/events';
  import { AppCommands } from './lib/services/app-commands';
  import { PanelLeft } from 'lucide-svelte';
  import Welcome from './lib/components/views/Welcome.svelte';
  import SetupWizard from './lib/components/views/SetupWizard.svelte';
  import Toolbar from './lib/components/toolbar/Toolbar.svelte';
  import TestExplorer from './lib/components/explorer/TestExplorer.svelte';
  import Canvas from './lib/components/Canvas.svelte';
  import RightSidebar from './lib/components/panels/RightSidebar.svelte';
  import LogConsoleDrawer from './lib/components/panels/LogConsoleDrawer.svelte';
  import AppModals from './lib/components/modals/AppModals.svelte';

  $effect(() => {
    const cleanup = AppCommands.init();
    return () => {
      cleanup();
    };
  });

  $effect(() => {
    if (!projectStore.path) {
      settingsStore.load();
      historyStore.load();
    }
  });

  $effect(() => {
    const path = projectStore.path;
    if (path) {
      void ProjectService.watchProject(path).catch((e) => {
        console.error('ProjectService.watchProject failed:', e);
      });
    }
  });

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

  $effect(() => {
    if (editorStore.selectedNodeId !== null) {
      applicationStore.isRightSidebarOpen = true;
    }
  });

  $effect(() => {
    applicationStore.isSidebarOpen = editorStore.selectedFilePath === null;
  });

  function handleKeydown(e: KeyboardEvent) {
    if (!projectStore.path) return;
    if ((e.metaKey || e.ctrlKey) && e.key.toLowerCase() === 'b') {
      e.preventDefault();
      applicationStore.toggleSidebar();
    } else if ((e.metaKey || e.ctrlKey) && e.key.toLowerCase() === 'j') {
      e.preventDefault();
      applicationStore.toggleConsole();
    }
  }
</script>

<svelte:window onkeydown={handleKeydown} />

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
    <LogConsoleDrawer />

    {#if !applicationStore.isSidebarOpen}
      <button
        type="button"
        onclick={() => applicationStore.toggleSidebar()}
        class="absolute bottom-3 left-3 z-30 flex h-6 w-6 cursor-pointer items-center justify-center text-zinc-500 transition-colors duration-200 hover:text-zinc-200 focus:outline-none"
        title="Show Test Explorer"
      >
        <PanelLeft class="h-4 w-4" />
      </button>
    {/if}
  {/if}

  <AppModals />
</div>
