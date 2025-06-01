<script lang="ts">
	import CloseButton from './CloseButton.svelte';

	export let isOpen = false;
	export let action: () => Promise<void>;
	export let submitText: string = 'Confirm';
</script>

<dialog class="modal" class:modal-open={isOpen}>
	<div class="modal-box flex flex-col items-center">
		<CloseButton on:click={() => (isOpen = false)} />
		<slot />
		<footer class="flex justify-around gap-4">
			<form on:submit|preventDefault={action}>
				<button class="btn" on:click={() => (isOpen = false)}>{submitText}</button>
			</form>
			<button class="btn btn-ghost" on:click={() => (isOpen = false)}> Cancel </button>
		</footer>
	</div>
</dialog>
