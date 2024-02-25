<script lang="ts">
	import { DashLayout } from '$lib/client/layout';
	import ProjectsList from './(components)/ProjectsList.svelte';
	import CreateProjectForm from './(forms)/CreateProjectForm.svelte';
	import route from './route.meta';

	export let data;
</script>

<DashLayout {route}>
	<ProjectsList projects={data.projects} />
	<CreateProjectForm formData={data.CreateProjectForm} />
</DashLayout>
