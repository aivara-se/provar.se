<script lang="ts">
	import type { Feedback } from '$lib/types';
	import { Avatar } from 'flowbite-svelte';

	export let user: Feedback['user'];
	export let withName: boolean = false;

	$: name = user?.name || 'Anonymous';
</script>

<div class="mt-2 flex items-center gap-2">
	{#if user?.image}
		<Avatar src={user?.image} size="sm" />
	{:else}
		<Avatar size="sm" border />
	{/if}
	{#if withName}
		<cite class="text-sm font-semibold">{name}</cite>
	{/if}
</div>
