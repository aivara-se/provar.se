<script module>
	import { defineMeta } from '@storybook/addon-svelte-csf';
	import TextInput from './TextInput.svelte';

	const { Story } = defineMeta({
		title: 'Components/TextInput',
		component: TextInput,
		argTypes: {
			name: { control: 'text' },
			value: { control: 'text' },
			label: { control: 'text' },
			error: { control: 'text' },
			className: { control: 'text' },
			props: { control: 'object' }
		}
	});
</script>

<Story
	name="default"
	args={{
		name: 'default-input',
		value: '',
		label: '',
		error: '',
		className: '',
		props: {}
	}}
/>

<Story
	name="with label"
	args={{
		name: 'labeled-input',
		value: '',
		label: 'Enter your name',
		error: '',
		className: '',
		props: { placeholder: 'John Doe' }
	}}
/>

<Story
	name="with error"
	args={{
		name: 'error-input',
		value: 'invalid@',
		label: 'Email Address',
		error: 'Please enter a valid email address',
		className: '',
		props: { placeholder: 'user@example.com' }
	}}
/>

<Story
	name="large size"
	args={{
		name: 'large-input',
		value: '',
		label: 'Large Input',
		error: '',
		className: 'input-lg',
		props: { placeholder: 'This is a larger input' }
	}}
/>

<Story
	name="with custom props"
	args={{
		name: 'custom-input',
		value: '',
		label: 'Custom Input',
		error: '',
		className: '',
		props: {
			placeholder: 'Type something...',
			maxlength: 50,
			'data-testid': 'custom-text-input'
		}
	}}
/>
