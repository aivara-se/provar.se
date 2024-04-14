<script lang="ts">
	import { browser } from '$app/environment';
	import { goto } from '$app/navigation';
	import { page } from '$app/stores';
	import { DashLayout } from '$lib/client/layout';
	import { resetSearchValue } from '$lib/client/search';
	import { endOfDay, endOfToday, format, startOfDay, startOfToday } from 'date-fns';
	import { onMount } from 'svelte';
	import ActionsDropdown from './(components)/ActionsDropdown.svelte';
	import DurationPicker from './(components)/DurationPicker.svelte';
	import FeedbackResults from './(components)/FeedbackResults.svelte';
	import SearchAndFilter from './(components)/SearchAndFilter.svelte';
	import route from './route.meta';

	export let data;

	let range = {
		beg: startOfToday(),
		end: endOfToday()
	};

	$: {
		if (browser) {
			const begDateParam = $page.url.searchParams.get('beg');
			const endDateParam = $page.url.searchParams.get('end');
			if (!begDateParam || !endDateParam) {
				const url = new URL($page.url);
				url.searchParams.set('beg', format(range.beg, 'yyyy-MM-dd'));
				url.searchParams.set('end', format(range.end, 'yyyy-MM-dd'));
				goto(url.toString());
			}
		}
	}

	$: {
		const begDateParam = $page.url.searchParams.get('beg');
		const endDateParam = $page.url.searchParams.get('end');
		if (begDateParam && endDateParam) {
			range = {
				beg: startOfDay(new Date(begDateParam)),
				end: endOfDay(new Date(endDateParam))
			};
		}
	}

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
