<script lang="ts">
	import type { Route } from '$lib/client/routes';

	export let route: Route;
</script>

<div
	class="flex flex-1 items-center justify-center min-h-screen w-full h-full bg-gray-100 dark:bg-gray-900"
>
	<slot />
</div>
