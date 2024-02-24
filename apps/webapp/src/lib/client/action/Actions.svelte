<script lang="ts">
	import { page } from '$app/stores';
	import ActionSearchInput from './(components)/ActionSearchInput.svelte';
	import ActionsListView from './(components)/ActionsListView.svelte';
	import type { Action, ActionGroup } from './action.types';
	import { getActions, performAction } from './action.utils';

	let input: HTMLInputElement;
	let value = '';
	let visible = false;

	$: groups = getActions($page.url.pathname, $page.params);
	$: result = groups
		.map((group) => (value ? filterGroupActions(group) : group))
		.filter((group) => group.actions.length > 0);

	function filterGroupActions(group: ActionGroup) {
		return {
			...group,
			actions: group.actions.filter((action) =>
				action.name.toLowerCase().includes(value.toLowerCase())
			)
		};
	}

	function triggerAction(action: Action) {
		performAction(action);
		closeModal();
	}

	function openModal() {
		visible = true;
		focusSearch();
	}

	function closeModal() {
		visible = false;
	}

	function focusSearch() {
		input.focus();
	}

	function shouldEscape(event: KeyboardEvent) {
		return event.key === 'Escape';
	}

	function shouldToggle(event: KeyboardEvent) {
		return (event.ctrlKey || event.metaKey) && (event.key === 'k' || event.key === 'p');
	}

	function onKeyDown(event: KeyboardEvent) {
		if (visible) {
			if (shouldEscape(event) || shouldToggle(event)) {
				event.preventDefault();
				return closeModal();
			}
		} else {
			if (shouldToggle(event)) {
				event.preventDefault();
				openModal();
			}
		}
	}

	function onClickOutside() {
		closeModal();
	}
</script>

<div
	role="presentation"
	class="modal p-8"
	class:modal-open={visible}
	on:click|self={onClickOutside}
>
	<div class="w-full max-w-lg rounded-lg shadow-lg p-4 bg-gray-100 dark:bg-gray-800">
		<header>
			<ActionSearchInput bind:element={input} bind:value />
		</header>
		<ActionsListView
			groups={result}
			enabled={visible}
			on:action={(event) => triggerAction(event.detail)}
			on:focus-search={() => focusSearch()}
		/>
	</div>
</div>

<svelte:window on:keydown={onKeyDown} />
