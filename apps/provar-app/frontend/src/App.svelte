<script lang="ts">
  import { projectStore } from './lib/stores/project-store.svelte';
  import { settingsStore } from './lib/stores/settings-store.svelte';
  import { historyStore } from './lib/stores/history-store.svelte';
  import { editorStore } from './lib/stores/editor-store.svelte';
  import { applicationStore } from './lib/stores/application-store.svelte';
  import { ProjectService } from './lib/services/project-service';
  import { subscribe } from './lib/services/events';
  import Welcome from './lib/components/views/Welcome.svelte';
  import SetupWizard from './lib/components/views/SetupWizard.svelte';
  import Toolbar from './lib/components/toolbar/Toolbar.svelte';
  import TestExplorer from './lib/components/explorer/TestExplorer.svelte';
  import Canvas from './lib/components/Canvas.svelte';
  import RightSidebar from './lib/components/panels/RightSidebar.svelte';
  import AppModals from './lib/components/modals/AppModals.svelte';

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
