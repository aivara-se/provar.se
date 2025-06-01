<script lang="ts">
	import '../../app.css';

	let { children } = $props();
</script>

<div class="flex items-center justify-center h-screen w-full">
	{@render children()}
</div>
