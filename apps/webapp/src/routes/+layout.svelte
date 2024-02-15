<script>
	import '../app.css';

	import './auth/login/route.meta';
	import './auth/logout/route.meta';
	import './auth/verify/route.meta';
	import './org/[organizationId]/feedback/route.meta';
	import './org/[organizationId]/projects/route.meta';
	import './org/[organizationId]/settings/credentials/route.meta';
	import './org/[organizationId]/settings/members/route.meta';
	import './org/[organizationId]/settings/organization/route.meta';
	import './org/[organizationId]/settings/route.meta';
	import './org/create/route.meta';
</script>

<slot />
