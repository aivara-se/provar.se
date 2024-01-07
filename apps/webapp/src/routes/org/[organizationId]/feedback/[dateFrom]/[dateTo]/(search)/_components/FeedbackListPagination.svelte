<script lang="ts">
	import { goto } from '$app/navigation';
	import { page } from '$app/stores';
	import { Pagination, type LinkType } from 'flowbite-svelte';
	import { ChevronLeftOutline, ChevronRightOutline } from 'flowbite-svelte-icons';
	import { createEventDispatcher } from 'svelte';

	export let pages: number = 1;
	export let currentPage: number = 1;

	const dispatch = createEventDispatcher();

	let paginationItems: LinkType[] = [];

	$: {
		paginationItems = [];
		const min = Math.max(1, currentPage - 2);
		const max = Math.min(pages, currentPage + 2);
		for (let i = min; i <= max; i++) {
			let name = i.toString();
			if ((i === min && min > 1) || (i === max && max < pages)) {
				name = '...';
			}
			const url = new URL($page.url.toString());
			url.searchParams.set('page', String(i));
			paginationItems.push({
				name,
				href: url.toString(),
				active: i === currentPage
			});
		}
	}

	function previousPage() {
		const url = new URL($page.url.toString());
		if (currentPage <= 1) {
			url.searchParams.set('page', String(1));
		} else {
			url.searchParams.set('page', String(currentPage + 1));
		}
		goto(url.toString());
	}

	function nextPage() {
		const url = new URL($page.url.toString());
		if (currentPage >= pages) {
			url.searchParams.set('page', String(pages));
		} else {
			url.searchParams.set('page', String(currentPage + 1));
		}
		goto(url.toString());
	}
</script>

<section class="pagination">
	<Pagination
		pages={paginationItems}
		on:previous={previousPage}
		on:next={nextPage}
		on:click={() => dispatch('change', {})}
		icon
	>
		<svelte:fragment slot="prev">
			<span class="sr-only">Previous</span>
			<ChevronLeftOutline class="w-2 h-2 outline-none" />
		</svelte:fragment>
		<svelte:fragment slot="next">
			<span class="sr-only">Next</span>
			<ChevronRightOutline class="w-2 h-2 outline-none" />
		</svelte:fragment>
	</Pagination>
</section>

<style>
	.pagination {
		display: flex;
		justify-content: center;
		padding: 20px 0 20px 0;
		box-sizing: border-box;
	}
</style>
