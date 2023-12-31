<script lang="ts">
	import { goto } from '$app/navigation';
	import { page } from '$app/stores';
	import { Button, ButtonGroup, Input } from 'flowbite-svelte';
	import { SearchOutline } from 'flowbite-svelte-icons';

	const SEARCH_QUERY_PARAM = 'search';

	interface Fields {
		search: string;
	}

	let fields: Fields = {
		search: new URL($page.url.toString()).searchParams.get(SEARCH_QUERY_PARAM) ?? ''
	};

	function handler() {
		const url = new URL($page.url.toString());
		url.searchParams.set(SEARCH_QUERY_PARAM, String(fields.search));
		goto(url.toString());
	}
</script>

<section>
	<form on:submit|self|stopPropagation|preventDefault={handler}>
		<ButtonGroup class="w-full">
			<Input id="search" size="sm" placeholder="Search ..." bind:value={fields.search} />
			<Button type="submit" color="light">
				<SearchOutline class="w-3 h-3" />
			</Button>
		</ButtonGroup>
	</form>
</section>
