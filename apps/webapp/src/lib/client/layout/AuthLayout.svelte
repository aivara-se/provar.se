<script lang="ts">
	import { page } from '$app/stores';
	import type { Route } from '$lib/types';

	export let route: Route;

	$: name = route.getName($page.params);
	$: path = route.getName($page.params);
</script>

<div class="flex flex-1 items-center justify-center min-h-screen w-full h-full">
	<slot />
</div>
