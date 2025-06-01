<script module>
	import { SelectInput } from '$lib/index.js';
	import { defineMeta } from '@storybook/addon-svelte-csf';

	const { Story } = defineMeta({
		title: 'Forms/SelectInput',
		component: SelectInput,
		argTypes: {
			name: { control: 'text' },
			value: { control: 'text' },
			label: { control: 'text' },
			error: { control: 'text' },
			className: { control: 'text' },
			options: { control: 'object' },
			props: { control: 'object' }
		}
	});

	const basicOptions = [
		{ label: 'Option 1', value: 'option1' },
		{ label: 'Option 2', value: 'option2' },
		{ label: 'Option 3', value: 'option3' }
	];

	const countryOptions = [
		{ label: 'United States', value: 'us' },
		{ label: 'Canada', value: 'ca' },
		{ label: 'United Kingdom', value: 'uk' },
		{ label: 'Germany', value: 'de' },
		{ label: 'France', value: 'fr' },
		{ label: 'Japan', value: 'jp' }
	];
</script>

<Story
	name="default"
	args={{
		name: 'default-select',
		value: '',
		options: basicOptions,
		label: '',
		error: '',
		className: '',
		props: {}
	}}
/>

<Story
	name="with label"
	args={{
		name: 'labeled-select',
		value: '',
		options: countryOptions,
		label: 'Country',
		error: '',
		className: '',
		props: {}
	}}
/>

<Story
	name="with error"
	args={{
		name: 'error-select',
		value: '',
		options: basicOptions,
		label: 'Required Selection',
		error: 'Please select an option',
		className: '',
		props: {}
	}}
/>

<Story
	name="pre-selected value"
	args={{
		name: 'preselected-select',
		value: 'ca',
		options: countryOptions,
		label: 'Country',
		error: '',
		className: '',
		props: {}
	}}
/>

<Story
	name="empty options"
	args={{
		name: 'empty-select',
		value: '',
		options: [],
		label: 'No Options Available',
		error: 'No options to select from',
		className: '',
		props: { disabled: true }
	}}
/>
