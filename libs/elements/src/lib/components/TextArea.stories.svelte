<script module>
	import { defineMeta } from '@storybook/addon-svelte-csf';
	import TextArea from './TextArea.svelte';

	const { Story } = defineMeta({
		title: 'Components/TextArea',
		component: TextArea,
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
		name: 'default-textarea',
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
		name: 'labeled-textarea',
		value: '',
		label: 'Description',
		error: '',
		className: '',
		props: { placeholder: 'Enter a detailed description...' }
	}}
/>

<Story
	name="with error"
	args={{
		name: 'error-textarea',
		value: 'Too short',
		label: 'Comments',
		error: 'Comments must be at least 10 characters long',
		className: '',
		props: { placeholder: 'Please provide your feedback...' }
	}}
/>

<Story
	name="large size"
	args={{
		name: 'large-textarea',
		value: '',
		label: 'Large Text Area',
		error: '',
		className: 'textarea-lg h-48',
		props: { placeholder: 'This is a larger text area for longer content...' }
	}}
/>

<Story
	name="with content"
	args={{
		name: 'content-textarea',
		value: 'This is some pre-filled content in the textarea. It demonstrates how the component looks with existing text.',
		label: 'Message',
		error: '',
		className: '',
		props: {
			placeholder: 'Enter your message...',
			maxlength: 500
		}
	}}
/>
