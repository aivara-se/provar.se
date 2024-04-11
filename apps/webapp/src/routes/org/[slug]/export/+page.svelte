<script lang="ts">
	import { DashLayout } from '$lib/client/layout';
	import route from './route.meta';

	export let data;
</script>

<DashLayout {route} user={data.user} organizations={data.organizations}>
	<p>[export]</p>
</DashLayout>
