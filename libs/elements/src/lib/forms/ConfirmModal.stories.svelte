<script module>
	import { ConfirmModal } from '$lib/index.js';
	import { defineMeta } from '@storybook/addon-svelte-csf';

	const { Story } = defineMeta({
		title: 'Forms/ConfirmModal',
		component: ConfirmModal,
		argTypes: {
			isOpen: { control: 'boolean' },
			submitText: { control: 'text' },
			action: { action: 'confirmed' }
		}
	});

	// Mock action function for stories
	const mockAction = async () => {
		console.log('Action confirmed!');
		await new Promise(resolve => setTimeout(resolve, 1000));
	};
</script>

<Story
	name="closed"
	args={{
		isOpen: false,
		submitText: 'Confirm',
		action: mockAction
	}}
>
	Are you sure you want to proceed?
</Story>

<Story
	name="open"
	args={{
		isOpen: true,
		submitText: 'Confirm',
		action: mockAction
	}}
>
	Are you sure you want to proceed?
</Story>

<Story
	name="delete confirmation"
	args={{
		isOpen: true,
		submitText: 'Delete',
		action: mockAction
	}}
>
	<div class="text-center">
		<svg class="w-16 h-16 mx-auto text-red-500 mb-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
			<path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16"></path>
		</svg>
		<h3 class="text-lg font-bold mb-2">Delete Item</h3>
		<p>This action cannot be undone. Are you sure you want to delete this item?</p>
	</div>
</Story>

<Story
	name="save changes"
	args={{
		isOpen: true,
		submitText: 'Save Changes',
		action: mockAction
	}}
>
	<div class="text-center">
		<svg class="w-16 h-16 mx-auto text-blue-500 mb-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
			<path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 7H5a2 2 0 00-2 2v9a2 2 0 002 2h14a2 2 0 002-2V9a2 2 0 00-2-2h-3m-1 4l-3 3m0 0l-3-3m3 3V4"></path>
		</svg>
		<h3 class="text-lg font-bold mb-2">Save Changes</h3>
		<p>Do you want to save your changes before leaving?</p>
	</div>
</Story>

<Story
	name="interactive example">
	{#snippet template(args)}
		<div class="p-4">
			<button
				class="btn btn-primary"
				on:click={() => args.isOpen = true}
			>
				Open Modal
			</button>

			<ConfirmModal
				isOpen={args.isOpen}
				submitText="Yes, Continue"
				action={mockAction}
			>
				<div class="text-center">
					<h3 class="text-lg font-bold mb-2">Confirm Action</h3>
					<p>Click the button above to open this confirmation modal.</p>
				</div>
			</ConfirmModal>
		</div>
	{/snippet}
</Story>
