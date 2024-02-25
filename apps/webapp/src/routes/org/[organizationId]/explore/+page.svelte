<script lang="ts">
	import { DashLayout } from '$lib/client/layout';
	import DurationPicker from './(components)/DurationPicker.svelte';
	import ActionsDropdown from './(components)/ActionsDropdown.svelte';
	import route from './route.meta';
</script>

<DashLayout {route}>
	<section slot="actions" class="flex gap-1 items-center">
		<DurationPicker />
		<ActionsDropdown />
	</section>

	<p>[explore]</p>
</DashLayout>
