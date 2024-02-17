<script lang="ts">
	import { DashLayout } from '$lib/client/layout';
	import ProjectsTable from './(components)/ProjectsTable.svelte';
	import CreateProjectForm from './(forms)/CreateProjectForm.svelte';
	import route from './route.meta';

	export let data;
</script>

<DashLayout {route}>
	<ProjectsTable projects={data.projects} />
	<CreateProjectForm formData={data.CreateProjectForm} />
</DashLayout>
