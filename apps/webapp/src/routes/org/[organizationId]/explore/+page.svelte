<script lang="ts">
	import { DashLayout } from '$lib/client/layout';
	import route from './route.meta';
</script>

<DashLayout {route}>
	<p slot="actions">[actions]</p>
	<p>[explore]</p>
</DashLayout>
