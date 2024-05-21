<script lang="ts">
	import { setConfig } from '$lib/client/store';
	import { type FeedbackType } from '$lib/client/types';
	import { goto } from '$app/navigation';

	interface Option {
		value: FeedbackType;
		title: string;
		active: boolean;
	}

	let options: Option[] = [
		{ value: 'cnps', title: 'Would customers recommend you?', active: false },
		{ value: 'csat', title: 'How satisfied are your customers?', active: false }
	];

	$: configured = options.filter((option) => option.active).length > 0;

	function select(option: Option) {
		option.active = !option.active;
		options = [...options];
	}

	function next() {
		const feedbackTypes = options.filter((o) => o.active).map((o) => o.value);
		setConfig({ feedbackTypes });
		goto('/admin');
	}
</script>

<main class="flex-1 flex items-center justify-center h-dvh overflow-hidden">
	<content class="flex flex-col items-center w-full p-4">
		<p class="text-3xl text-center leading-relaxed subpixel-antialiased">
			Select the type of feedback you would like to gather
		</p>
		<div class="flex flex-row mt-8 gap-4">
			{#each options as option}
				<button
					class="btn btn-lg"
					class:btn-success={option.active}
					on:click={() => select(option)}
				>
					{option.title}
				</button>
			{/each}
		</div>
	</content>
</main>

<footer class="fixed bottom-0 right-0 p-4">
	<button class="btn px-8" disabled={!configured} class:btn-info={configured} on:click={next}>
		Next
	</button>
</footer>
