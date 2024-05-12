<script lang="ts">
	import { page } from '$app/stores';
	import { parseDateRange } from '$lib/client/dates';
	import { DashLayout } from '$lib/client/layout';
	import { resetSearchValue } from '$lib/client/search';
	import { onMount } from 'svelte';
	import ActionsDropdown from './(components)/ActionsDropdown.svelte';
	import DurationPicker from './(components)/DurationPicker.svelte';
	import FeedbackResults from './(components)/FeedbackResults.svelte';
	import SearchAndFilter from './(components)/SearchAndFilter.svelte';
	import route from './route.meta';

	export let data;

	$: range = parseDateRange({
		from: $page.url.searchParams.get('from') || undefined,
		to: $page.url.searchParams.get('to') || undefined
	});

	onMount(() => {
		resetSearchValue();
	});
</script>

<DashLayout {route} user={data.user} organizations={data.organizations}>
	<section slot="actions" class="flex gap-2 items-center">
		<DurationPicker bind:range />
		<ActionsDropdown />
	</section>

	<SearchAndFilter />
	<FeedbackResults feedbacks={data.feedbacks} />
</DashLayout>
