<script lang="ts">
	import { DashLayout } from '$lib/client/layout';
	import ActionsDropdown from './(components)/ActionsDropdown.svelte';
	import DurationPicker from './(components)/DurationPicker.svelte';
	import FeedbackResults from './(components)/FeedbackResults.svelte';
	import SearchAndFilter from './(components)/SearchAndFilter.svelte';
	import route from './route.meta';
</script>

<DashLayout {route}>
	<section slot="actions" class="flex gap-2 items-center">
		<DurationPicker />
		<ActionsDropdown />
	</section>

	<SearchAndFilter />
	<FeedbackResults />
</DashLayout>
