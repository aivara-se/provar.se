<script lang="ts">
	import { page } from '$app/stores';
	import { DashLayout } from '$lib/client/layout';
	import route from './route.meta';
</script>

<DashLayout {route}>
	<slot />
</DashLayout>
