<script lang="ts">
	import '../../app.css';

	interface Props {
		children: () => any;
		isOpen?: boolean;
		action?: () => Promise<void>;
		submitText?: string;
	}

	let {
		children,
		isOpen = false,
		action = () => Promise.resolve(),
		submitText = 'Confirm'
	}: Props = $props();
</script>

<dialog class="modal" class:modal-open={isOpen}>
	<div class="modal-box">
		<form method="dialog">
			<button class="btn btn-sm btn-circle btn-ghost absolute right-2 top-2" onclick={() => (isOpen = false)}>✕</button>
		</form>

		{@render children()}

		<div class="modal-action">
			<form method="dialog" onsubmit={action}>
				<button class="btn btn-primary">{submitText}</button>
			</form>
			<button class="btn btn-ghost" onclick={() => (isOpen = false)}>Cancel</button>
		</div>
	</div>
</dialog>
