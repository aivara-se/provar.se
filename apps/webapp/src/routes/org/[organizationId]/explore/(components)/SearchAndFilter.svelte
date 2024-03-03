<script lang="ts">
	import { getSearchStore, getSearchValue, setSearchValue } from '$lib/client/search';
	import { parseSearch, type SearchQuery } from '$lib/shared/search';
	import type { FeedbackType } from '@provar/provar-js';
	import pick from 'lodash/pick';
	import { SearchIcon } from 'lucide-svelte';
	import { onMount } from 'svelte';

	interface Fields {
		search: string;
		query: SearchQuery;
	}

	let fields: Fields = {
		search: getSearchValue().text.join(' '),
		query: getSearchValue()
	};
	let hasBadges = false;

	$: {
		hasBadges =
			fields.query.type.length > 0 ||
			Object.keys(fields.query.tags).length > 0 ||
			Object.keys(fields.query.meta).length > 0;
	}

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
		<div class="join flex shadow">
			<input
				class="input input-sm join-item flex-1"
				placeholder="Search ..."
				bind:value={fields.search}
			/>
			<button type="submit" class="btn btn-sm join-item">
				<SearchIcon class="w-3 h-3" />
			</button>
		</div>
	</form>

	{#if hasBadges}
		<div class="my-4 flex flex-wrap gap-1">
			{#key fields.query.type}
				{#each fields.query.type as type}
					<button class="badge" on:click={() => removeTypeFilter(type)}>type: {type}</button>
				{/each}
			{/key}

			{#key fields.query.tags}
				{#each Object.keys(fields.query.tags || {}).sort() as key}
					<button class="badge text-xs" on:click={() => removeTagFilter(key)}>
						#{key}: {fields.query.tags?.[key]}
					</button>
				{/each}
			{/key}

			{#key fields.query.meta}
				{#each Object.keys(fields.query.meta || {}).sort() as key}
					<button class="badge text-xs" on:click={() => removeMetaFilter(key)}>
						${key}: {fields.query.meta?.[key]}
					</button>
				{/each}
			{/key}
		</div>
	{/if}
</section>
