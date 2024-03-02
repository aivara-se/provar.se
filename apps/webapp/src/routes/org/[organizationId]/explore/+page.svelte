<script lang="ts">
	import { DashLayout } from '$lib/client/layout';
	import { addDays, endOfToday, startOfDay } from 'date-fns';
	import ActionsDropdown from './(components)/ActionsDropdown.svelte';
	import DurationPicker from './(components)/DurationPicker.svelte';
	import FeedbackResults from './(components)/FeedbackResults.svelte';
	import SearchAndFilter from './(components)/SearchAndFilter.svelte';
	import route from './route.meta';

	let range = {
		beg: startOfDay(addDays(endOfToday(), -29)),
		end: endOfToday()
	};
</script>

<DashLayout {route}>
	<section slot="actions" class="flex gap-2 items-center">
		<DurationPicker bind:range />
		<ActionsDropdown />
	</section>

	<SearchAndFilter />
	<FeedbackResults />
</DashLayout>
