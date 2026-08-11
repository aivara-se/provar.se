<script lang="ts">
  import type { Snippet } from 'svelte';
  interface Props {
    children: Snippet;
  }
  let { children }: Props = $props();
</script>

<div
  class="shrink-0 flex items-center justify-end gap-0.5 border-t border-zinc-800/50 px-3 py-2 select-none"
>
  {@render children()}
</div>
