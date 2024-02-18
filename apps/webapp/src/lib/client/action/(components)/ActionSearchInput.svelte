<script lang="ts">
	import { SearchIcon } from 'lucide-svelte';

	export let element: HTMLInputElement;
	export let value = '';
</script>

<div class="relative">
	<input
		bind:value
		bind:this={element}
		type="text"
		class="input w-full"
		placeholder="Search for commands"
	/>
	<SearchIcon class="w-4 h-4 absolute top-4 right-4 pointer-events-none opacity-50" />
</div>
