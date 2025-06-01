<script lang="ts">
	import '../../app.css';

	interface Props {
		name: string;
		value: string;
		label?: string;
		error?: string;
		props?: object;
		className?: string;
	}

	let {
		name,
		value = $bindable(),
		label = '',
		error = '',
		props = {},
		className = ''
	}: Props = $props();
</script>

<fieldset class="fieldset">
	{#if label}
		<legend class="fieldset-legend">{label}</legend>
	{/if}

	<textarea
		{name}
		bind:value
		class="textarea h-24 {error ? 'textarea-error' : ''} {className}"
		aria-invalid={error ? 'true' : undefined}
		{...props}
	></textarea>

	{#if error}
		<p class="label text-error">{error}</p>
	{/if}
</fieldset>
