<script lang="ts">
	import { DownloadIcon, MoreHorizontalIcon, UploadIcon } from 'lucide-svelte';

	let element: HTMLDetailsElement;

	function closeDropdown() {
		element.removeAttribute('open');
	}

	function onKeyDown(event: KeyboardEvent) {
		if (event.key !== 'Escape') {
			return;
		}
		event.preventDefault();
		closeDropdown();
		if (document.activeElement instanceof HTMLElement) {
			document.activeElement.blur();
		}
	}
</script>

<details bind:this={element} class="dropdown dropdown-end">
	<summary class="btn btn-sm btn-ghost">
		<MoreHorizontalIcon class="w-3.5 h-3.5" />
	</summary>

	<ul class="p-2 shadow menu dropdown-content z-[1] bg-base-100 rounded-box w-32">
		<li>
			<a href="./import" class="flex justify-between">
				Import
				<UploadIcon class="w-3 h-3" />
			</a>
		</li>
		<li>
			<a href="./export" class="flex justify-between">
				Export
				<DownloadIcon class="w-3 h-3" />
			</a>
		</li>
	</ul>
</details>

<svelte:window on:keydown={onKeyDown} on:click={closeDropdown} />
