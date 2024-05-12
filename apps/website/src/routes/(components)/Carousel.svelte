<script lang="ts">
	export let images: string[] = [];

	let idx = 0;

	function setIdx(i: number) {
		idx = i;
	}
</script>

<section class="relative my-4 md:my-8 mx-auto max-w-screen-xl">
	<img src={images[0]} alt="" aria-hidden class="w-full invisible" />
	{#each images as image, i}
		<img
			src={image}
			alt="Dashboard"
			class="w-full rounded-xl absolute top-0 shadow-xl {i === idx ? 'opacity-100' : 'opacity-0'}"
		/>
	{/each}
	<div class="flex gap-4 mt-4 justify-center">
		{#each images as _, i}
			<button
				class="btn btn-xs btn-circle {idx === i ? 'btn-active' : ''}"
				on:click={() => setIdx(i)}
			>
				{i + 1}
			</button>
		{/each}
	</div>
</section>

<style>
	img {
		transition: opacity 0.5s;
	}
</style>
