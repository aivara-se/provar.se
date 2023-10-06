<script lang="ts">
	import { Button, Input, Label, Select } from 'flowbite-svelte';

	const ENVIRONMENTS = [
		{ value: 'staging', name: 'Staging' },
		{ value: 'production', name: 'Production' }
	];

	let name = '';
	let environment: string = 'production';
</script>

<div class="wrapper">
	<main>
		<form method="post">
			<div class="grid gap-6 mb-6">
				<div>
					<Label color="gray" for="name" class="mb-2">Organization name</Label>
					<Input type="text" id="name" required bind:value={name} />
				</div>
				<div>
					<Label color="gray">
						Deploy environment
						<Select class="mt-2" items={ENVIRONMENTS} required bind:value={environment} />
					</Label>
				</div>
				<div>
					<Button outline type="submit" size="sm">Create organization</Button>
				</div>
			</div>
		</form>
	</main>
</div>

<style>
	.wrapper {
		display: flex;
		flex-direction: column;
		align-items: center;
		justify-content: center;
		height: 100vh;
	}

	main {
		display: flex;
		flex-direction: column;
		align-items: center;
		justify-content: center;
		text-align: center;
	}
</style>
