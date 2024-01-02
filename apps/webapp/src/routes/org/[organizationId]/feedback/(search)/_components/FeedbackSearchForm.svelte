<script lang="ts">
	import { getSearchStore, getSearchValue, setSearchValue } from '$lib/client/search';
	import { parseSearch, type SearchQuery } from '$lib/shared/search';
	import type { FeedbackType } from '@provar/provar-js';
	import { Badge, Button, ButtonGroup, Input } from 'flowbite-svelte';
	import { SearchOutline } from 'flowbite-svelte-icons';
	import pick from 'lodash/pick';
	import { onMount } from 'svelte';

	interface Fields {
		search: string;
		query: SearchQuery;
	}

	let fields: Fields = {
		search: getSearchValue().text.join(' '),
		query: getSearchValue()
	};

	onMount(() => {
		return getSearchStore().subscribe((value) => {
			fields.search = value.text.join(' ');
			fields.query = value;
		});
	});

	// Sync the search query param
	setSearchValue(fields.query);

	/**
	 * Update the search query param with updated input value
	 */
	function syncSearchValue() {
		const filters = pick(fields.query, ['type', 'tags', 'meta']);
		setSearchValue(filters, parseSearch(fields.search));
	}

	/**
	 * Remove a type filter
	 */
	function removeTypeFilter(type: FeedbackType) {
		const index = fields.query.type.indexOf(type);
		if (index > -1) {
			fields.query.type.splice(index, 1);
		}
		syncSearchValue();
	}

	/**
	 * Remove a tag filter
	 */
	function removeTagFilter(key: string) {
		delete fields.query.tags[key];
		syncSearchValue();
	}

	/**
	 * Remove a meta filter
	 */
	function removeMetaFilter(key: string) {
		delete fields.query.meta[key];
		syncSearchValue();
	}
</script>

<section>
	<form on:submit|self|stopPropagation|preventDefault={syncSearchValue}>
		<ButtonGroup class="w-full">
			<Input id="search" size="sm" placeholder="Search ..." bind:value={fields.search} />
			<Button type="submit" color="light">
				<SearchOutline class="w-3 h-3" />
			</Button>
		</ButtonGroup>
	</form>

	<div class="my-2 flex flex-wrap gap-1">
		{#key fields.query.type}
			{#each fields.query.type as type}
				<Badge dismissable color="dark" on:close={() => removeTypeFilter(type)}>
					type: {type}
				</Badge>
			{/each}
		{/key}

		{#key fields.query.tags}
			{#each Object.keys(fields.query.tags || {}).sort() as key}
				<Badge dismissable color="blue" on:close={() => removeTagFilter(key)}>
					{key}: {fields.query.tags?.[key]}
				</Badge>
			{/each}
		{/key}

		{#key fields.query.meta}
			{#each Object.keys(fields.query.meta || {}).sort() as key}
				<Badge dismissable color="pink" on:close={() => removeMetaFilter(key)}>
					{key}: {fields.query.meta?.[key]}
				</Badge>
			{/each}
		{/key}
	</div>
</section>
