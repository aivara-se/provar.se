<script lang="ts">
	import { goto } from '$app/navigation';
	import { page } from '$app/stores';
	import { onMount } from 'svelte';

	interface PaginationItem {
		name: string;
		href: string;
		active: boolean;
	}

	export let pages: number = 1;

	let activePage: number = 1;
	let listItems: PaginationItem[] = [];

	onMount(() => {
		const pageParam = $page.url.searchParams.get('page');
		if (!pageParam) {
			selectPage(1);
		}
	});

	$: {
		const pageParam = $page.url.searchParams.get('page');
		if (pageParam) {
			activePage = parseInt(pageParam);
		}
	}

	$: {
		listItems = [];
		const min = Math.max(1, activePage - 2);
		const max = Math.min(pages, activePage + 2);
		for (let i = min; i <= max; i++) {
			let name = i.toString();
			if ((i === min && min > 1) || (i === max && max < pages)) {
				name = '...';
			}
			const url = new URL($page.url.toString());
			url.searchParams.set('page', String(i));
			listItems.push({
				name,
				href: url.toString(),
				active: i === activePage
			});
		}
	}

	function selectPage(page: number) {
		const url = new URL($page.url.toString());
		url.searchParams.set('page', String(page));
		goto(url.toString());
	}
</script>

<section class="flex items-center justify-center p-4">
	<div class="join">
		{#each listItems as item}
			<a href={item.href} class="join-item btn btn-sm" class:btn-active={item.active}>
				{item.name}
			</a>
		{/each}
	</div>

	<span class="ml-4 text-xs">{pages} Pages</span>
</section>
