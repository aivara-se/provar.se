<script lang="ts">
	import { beforeUpdate } from 'svelte';
	import '../app.css';
	import './routes';
	import 'overlayscrollbars/overlayscrollbars.css';
	import { canAccessRoute } from '$lib/client/auth';
	import { page } from '$app/stores';
	import { goto } from '$app/navigation';

	beforeUpdate(() => {
		const whitelist = ['/auth'];
		if (!canAccessRoute($page.url.pathname, whitelist)) {
			goto('/auth/login');
		}
	});
</script>

<slot />
