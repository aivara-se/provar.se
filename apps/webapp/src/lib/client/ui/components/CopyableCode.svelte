<script lang="ts">
	import { ClipboardCheckOutline, ClipboardOutline } from 'flowbite-svelte-icons';
	import { fade } from 'svelte/transition';

	export let text = '';

	let copied = false;

	function resetIcon() {
		copied = false;
	}

	function copyText() {
		try {
			navigator.clipboard.writeText(text);
			copied = true;
			setTimeout(resetIcon, 2000);
		} catch {
			console.warn('Failed to copy text');
		}
	}
</script>

<section class="relative">
	<pre
		class="block w-full px-2 py-1 overflow-hidden resize-none bg-gray-100 border-0 break-words text-xs text-gray-900 leading-6 tracking-wide rounded">{text.trim()}</pre>
	{#if copied}
		<div transition:fade={{ delay: 100 }} class="absolute cursor-pointer top-2 right-2">
			<ClipboardCheckOutline class="w-3.5 h-3.5 text-green-700" />
		</div>
	{:else}
		<div
			transition:fade={{ delay: 100 }}
			class="absolute cursor-pointer top-2 right-2 text-gray-700 opacity-75 hover:opacity-100"
		>
			<ClipboardOutline class="w-3.5 h-3.5" on:click={copyText} />
		</div>
	{/if}
</section>
