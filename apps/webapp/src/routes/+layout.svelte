<script lang="ts">
	import '../app.css';
	import './routes';
	import 'overlayscrollbars/overlayscrollbars.css';
	import { user, organizations } from '$lib/client/stores';

	export let data;

	$: user.set(data.user);
	$: organizations.set(data.organizations);
</script>

<slot />
