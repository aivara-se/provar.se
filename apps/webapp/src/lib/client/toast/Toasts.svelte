<script lang="ts">
	import { toasts } from './toast.store';
</script>

<div class="toast">
	{#each $toasts as toast (toast.id)}
		<div class="alert alert-{toast.type}">
			<span>{toast.content}</span>
		</div>
	{/each}
</div>
