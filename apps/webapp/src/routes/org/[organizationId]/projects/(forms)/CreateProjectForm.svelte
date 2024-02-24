<script lang="ts">
	import { SelectInput, TextInput } from '$lib/client/forms';
	import { toast } from '$lib/client/toast';
	import { FeedbackType } from '$lib/types';
	import { superForm } from 'sveltekit-superforms';

	export let formData;

	const FEEDBACK_TYPE_OPTIONS: { label: string; value: FeedbackType }[] = [
		{ label: 'Simple Text Feedback', value: FeedbackType.Text },
		{ label: 'Customer Satisfaction', value: FeedbackType.CSAT },
		{ label: 'Net Promoter Score', value: FeedbackType.CNPS }
	];

	const { form, errors, constraints, enhance } = superForm(formData, {
		onResult: ({ result }) => {
			if (result.type === 'failure') {
				toast('error', 'Failed to create project');
			} else if (result.type === 'success') {
				toast('success', 'Project created successfully');
			}
		}
	});
</script>

<section class="mt-8 rounded-lg p-4 bg-gray-100 text-gray-900 dark:bg-gray-950 dark:text-gray-200">
	<h3 class="flex items-center justify-between text-lg font-medium">Create Project</h3>
	<p class="text-sm mt-2">
		Feedback collection projects are used to collect and analyze feedback for a specific purpose.
	</p>
	<form method="POST" action="?/createProject" class="flex flex-col gap-2 mt-2" use:enhance>
		<TextInput
			name="name"
			label="Name"
			className="border-none bg-black/10 dark:bg-white/10"
			errors={$errors.name}
			constraints={$constraints.name}
			bind:value={$form.name}
		/>
		<SelectInput
			name="feedbackType"
			label="Feedback type"
			options={FEEDBACK_TYPE_OPTIONS}
			className="border-none bg-black/10 dark:bg-white/10"
			errors={$errors.feedbackType}
			constraints={$constraints.feedbackType}
			bind:value={$form.feedbackType}
		/>
		<button class="btn btn-sm w-24 btn-neutral mt-4">Create</button>
	</form>
</section>
