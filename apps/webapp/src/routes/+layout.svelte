<script lang="ts">
	import '../app.css';
	import './routes';
	import 'overlayscrollbars/overlayscrollbars.css';
	import { Toasts } from '$lib/client/toast';
	import { Actions } from '$lib/client/action';
	import { setProfile } from '$lib/client/profile';

	export let data;

	/**
	 * Set user session data
	 */
	setProfile({
		...data.profile!,
		organizations: data.organizations
	});
</script>

<slot />

<Toasts />
<Actions />
