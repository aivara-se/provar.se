<script lang="ts">
	import { page } from '$app/stores';
	import { signOut } from '$lib/client/auth';
	import type { Organization, User } from '$lib/client/types';
	import { LogOutIcon } from 'lucide-svelte';

	export let user: User | null;
	export let organizations: Organization[];

	let element: HTMLDetailsElement;

	$: items = organizations.map((organization) => ({
		id: organization.id,
		name: organization.name,
		slug: organization.slug,
		isActive: $page.url.pathname.includes(`/org/${organization.slug}`)
	}));

	function closeDropdown() {
		element.removeAttribute('open');
	}

	function onKeyDown(event: KeyboardEvent) {
		if (event.key !== 'Escape') {
			return;
		}
		event.preventDefault();
		closeDropdown();
		if (document.activeElement instanceof HTMLElement) {
			document.activeElement.blur();
		}
	}
</script>

<details bind:this={element} class="dropdown dropdown-end">
	<summary class="flex list-none">
		<div class="avatar">
			<div class="w-6 h-6 rounded-full shadow-lg">
				<img src={user?.avatar} alt={user?.name} />
			</div>
		</div>
	</summary>
	<section class="shadow menu dropdown-content z-1 bg-base-100 rounded-box w-40">
		<h3 class="py-2 text-xs text-center font-semibold cursor-default">Organizations</h3>

		<ul class="mt-2">
			{#each items as item (item.id)}
				<li>
					<a class="flex justify-between" class:opacity-50={!item.isActive} href="/org/{item.slug}">
						{item.name}
					</a>
				</li>
			{/each}
		</ul>

		<hr class="mt-2 border-t border-gray-900" />

		<ul class="mt-2">
			<li>
				<a
					href="/auth/logout"
					class="flex justify-between"
					on:click|preventDefault={() => signOut()}
				>
					Logout
					<LogOutIcon class="w-3 h-3" />
				</a>
			</li>
		</ul>
	</section>
</details>

<svelte:window on:keydown={onKeyDown} on:click={closeDropdown} />
