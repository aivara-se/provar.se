<script lang="ts">
  import { Button } from 'flowbite-svelte';
	import { page } from '$app/stores';

	$: createLink = `/org/${$page.params.organizationId}/feedback/project/create`
</script>

<Button href={createLink}>
	Create Project
</Button>
