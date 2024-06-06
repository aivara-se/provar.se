<script lang="ts">
	import { goto } from '$app/navigation';
	import {
		getConfig,
		logoUrl,
		resetLogoUrl,
		setConfig,
		setLogoUrl,
		resetConfig
	} from '$lib/client/store';
	import { CNPS_FEEDBACK_TYPE, CSAT_FEEDBACK_TYPE } from '$lib/client/types';

	const LANGUAGES = [
		{ value: 'en', label: 'English' },
		{ value: 'sv', label: 'Svenska' }
	];

	let config = { ...getConfig() };
	let logoFiles: FileList;
	let feedbackTypes = [
		{ value: CNPS_FEEDBACK_TYPE, label: 'Net Promoter Score' },
		{ value: CSAT_FEEDBACK_TYPE, label: 'Customer Satisfaction Score' }
	].map((item) => ({
		...item,
		checked: config.feedbackTypes.includes(item.value)
	}));

	$: {
		config.feedbackTypes = feedbackTypes.filter((item) => item.checked).map((item) => item.value);
		setConfig({
			feedbackTypes: config.feedbackTypes,
			defaultLanguage: config.defaultLanguage,
			brandingColors: config.brandingColors
		});
	}

	async function fileToDataUrl(file: File): Promise<string> {
		return new Promise((resolve, reject) => {
			const reader = new FileReader();
			reader.readAsDataURL(file);
			reader.onerror = (e) => reject(e);
			reader.onload = () => resolve(reader.result as string);
		});
	}

	async function handleLogoChange() {
		const dataUrl = await fileToDataUrl(logoFiles[0]);
		setLogoUrl(dataUrl);
	}

	function handleFactoryReset() {
		resetConfig();
		resetLogoUrl();
		goto('/');
	}
</script>

<div class="max-w-4xl mx-auto space-y-8">
	<section class="rounded-box p-4 bg-black/5 shadow">
		<h2 class="text-xl font-medium mb-2">General Settings</h2>
		<p class="text-gray-600 text-sm">Configure Provar default settings here.</p>
		<div class="flex flex-col gap-4 mt-4">
			<div>
				<label for="language" class="font-medium mb-2">Feedback Types</label>
				<div class="flex gap-4 mt-4">
					{#each feedbackTypes as item}
						<div class="flex-1 form-control">
							<label class="label cursor-pointer">
								<span class="label-text">{item.label}</span>
								<input type="checkbox" class="checkbox" bind:checked={item.checked} />
							</label>
						</div>
					{/each}
				</div>
			</div>
			<div>
				<label for="language" class="font-medium mb-2">Default Language</label>
				<div class="flex gap-4 mt-4">
					{#each LANGUAGES as item}
						<div class="flex-1 form-control">
							<label class="label cursor-pointer">
								<span class="label-text">{item.label}</span>
								<input
									type="radio"
									bind:group={config.defaultLanguage}
									value={item.value}
									class="radio"
								/>
							</label>
						</div>
					{/each}
				</div>
			</div>
		</div>
	</section>

	<section class="rounded-box p-4 bg-black/5 shadow">
		<h2 class="text-xl font-medium mb-2">Branding</h2>
		<p class="text-gray-600 text-sm">
			Upload your logo and set your brand colors to customize the look of your feedback form.
		</p>
		<div class="flex flex-col gap-4 mt-4">
			<div class="flex flex-row gap-4">
				<div class="flex flex-col flex-1">
					<label for="logo" class="font-medium mb-2">Logo</label>
					<input
						type="file"
						id="logo"
						class="file-input w-full"
						accept="image/*"
						bind:files={logoFiles}
						on:change={handleLogoChange}
					/>
					<p class="text-xs text-gray-600 mt-1">
						Upload a PNG file with a transparent background and dimensions 128x128 or higher.
					</p>
				</div>
				<div class="flex flex-col">
					{#if $logoUrl}
						<img src={$logoUrl} alt="Logo" class="w-20 h-20" />
						<button class="btn btn-xs btn-ghost mt-2" on:click={resetLogoUrl}> Reset </button>
					{/if}
				</div>
			</div>
			<div class="flex flex-row gap-4">
				<label class="flex-1 flex-col">
					<span class="font-medium">Color 1</span>
					<div class="flex gap-2 mt-2 relative">
						<input
							type="color"
							class="absolute right-2 top-2"
							bind:value={config.brandingColors.primary}
						/>
						<input
							type="text"
							class="flex-1 input w-full"
							bind:value={config.brandingColors.primary}
						/>
					</div>
				</label>
				<label class="flex-1 flex-col">
					<span class="font-medium">Color 2</span>
					<div class="flex gap-2 mt-2 relative">
						<input
							type="color"
							class="absolute right-2 top-2"
							bind:value={config.brandingColors.secondary}
						/>
						<input
							type="text"
							class="flex-1 input w-full"
							bind:value={config.brandingColors.secondary}
						/>
					</div>
				</label>
			</div>
		</div>
	</section>

	<section class="rounded-box p-4 bg-black/5 shadow">
		<h2 class="text-xl font-medium mb-2">Export Data</h2>
		<p class="text-gray-600 text-sm">
			Download all collected feedback data as a CSV file for further analysis.
		</p>
		<div class="flex flex-col gap-4 mt-4">
			<button type="button" class="btn w-full"> Export Data (CSV) </button>
		</div>
	</section>

	<section class="rounded-box p-4 bg-black/5 shadow">
		<h2 class="text-xl font-medium mb-2">Factory Reset</h2>
		<p class="text-gray-600 text-sm">
			This action will permanently delete all collected feedback data and reset all settings to
			their default values.
		</p>
		<div class="flex flex-col gap-4 mt-4">
			<button
				type="button"
				class="btn w-full text-red-500 border border-red-900 hover:border-red-700"
				on:click={handleFactoryReset}
			>
				Delete All Data
			</button>
		</div>
	</section>
</div>

<style>
	input[type='color'] {
		background: transparent;
		appearance: none;
		border: none;
		width: 32px;
		height: 32px;
	}
	input[type='color']::-webkit-color-swatch-wrapper {
		padding: 0;
	}
	input[type='color']::-webkit-color-swatch {
		border: none;
		border-radius: 8px;
	}
</style>
