<script lang="ts">
  import { onDestroy, onMount } from 'svelte';
  import { editorStore } from '../../stores/editor-store.svelte';
  import { applicationStore } from '../../stores/application-store.svelte';
  import { InfiniteCanvas } from './infinite-canvas';

  import { executionStore } from '../../stores/execution-store.svelte';

  let container: HTMLDivElement;
  let canvas: InfiniteCanvas | null = null;

  onMount(async () => {
    await new Promise((r) => requestAnimationFrame(() => requestAnimationFrame(r)));

    canvas = new InfiniteCanvas();
    canvas.onNodeSelect = (id) => {
      editorStore.selectedNodeId = id;
      applicationStore.isSidebarOpen = false;
      if (id === null) {
        applicationStore.closeSidebars();
      }
    };
    try {
      await canvas.init(container);
      if (editorStore.currentFile) {
        canvas.renderGraph(editorStore.currentFile, executionStore.actionStates, editorStore.diagnostics);
      }
    } catch (e) {
      console.error("Canvas init failed:", e);
    }
  });

  $effect(() => {
    const file = editorStore.currentFile;
    const diagnostics = editorStore.diagnostics;
    if (canvas && file) {
      canvas.renderGraph(file, executionStore.actionStates, diagnostics);
    }
  });

  $effect(() => {
    const states = executionStore.actionStates;
    const path = executionStore.runningPathNodeIds;
    const compile = executionStore.compilationStates;
    const diagnostics = editorStore.diagnostics;
    if (canvas) {
      canvas.updateGraphState(states, path, compile, diagnostics);
    }
  });

  onDestroy(() => {
    canvas?.destroy();
  });
</script>

<div bind:this={container} class="absolute inset-0"></div>