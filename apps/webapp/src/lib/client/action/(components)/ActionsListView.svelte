<script lang="ts">
	import { createEventDispatcher } from 'svelte';
	import { OverlayScrollbarsComponent } from 'overlayscrollbars-svelte';
	import type { Action, ActionGroup } from '../action.types';

	const dispatch = createEventDispatcher();

	export let groups: ActionGroup[] = [];
	export let enabled = false;

	let actions: Action[] = [];
	let focused: Action | null = null;

	$: {
		actions = groups.flatMap((group) => group.actions);
		focused = actions[0] ?? null;
	}

	function focusAction(action?: Action) {
		focused = action ?? null;
	}

	function focusSearch() {
		dispatch('focus-search');
	}

	function performAction(action: Action) {
		dispatch('action', action);
	}

	function focusPrevAction() {
		const currentIndex = actions.indexOf(focused!);
		const prevIndex = currentIndex - 1;
		if (prevIndex < 0) {
			focusSearch();
		} else {
			focusAction(actions[prevIndex]);
		}
	}

	function focusNextAction() {
		const currentIndex = actions.indexOf(focused!);
		const nextIndex = currentIndex + 1;
		if (nextIndex < actions.length) {
			focusAction(actions[nextIndex]);
		} else {
			focusAction(actions[actions.length - 1]);
		}
	}

	function onKeyDown(event: KeyboardEvent) {
		if (!enabled) {
			return;
		}
		const validFocus = Boolean(focused && actions.find((a) => a.id === focused!.id));
		if (event.key === 'Enter' && validFocus) {
			performAction(focused!);
			return;
		}
		if (event.key === 'ArrowDown' || event.key === 'ArrowUp') {
			if (!validFocus) {
				focusAction(actions[0]);
			} else if (event.key === 'ArrowDown') {
				focusNextAction();
			} else if (event.key === 'ArrowUp') {
				focusPrevAction();
			}
		}
	}
</script>

<OverlayScrollbarsComponent defer class="flex mt-4 min-h-24 max-h-48">
	<div class="flex flex-col min-h-24 gap-2">
		{#if !groups.length}
			<div class="flex flex-1 items-center justify-center text-sm opacity-50 pointer-events-none">
				No results found
			</div>
		{/if}

		{#each groups as group (group.name)}
			{#if group.name}
				<div class="text-xs font-medium">
					{group.name}
				</div>
			{/if}

			{#each group.actions as action (action.id)}
				<div
					tabindex="0"
					role="button"
					class="flex px-3 py-2 items-center rounded justify-between text-sm cursor-pointer hover:bg-gray-200 hover:dark:bg-gray-700 transition-all {focused ===
					action
						? 'bg-gray-200 dark:bg-gray-700'
						: ''}"
					on:mouseover={() => focusAction(action)}
					on:focus={() => focusAction(action)}
				>
					{action.name}
				</div>
			{/each}
		{/each}
	</div>
</OverlayScrollbarsComponent>

<svelte:window on:keydown={onKeyDown} />
