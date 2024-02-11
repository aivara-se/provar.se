<script lang="ts">
	import { AuthLayout } from '$lib/client/layout';
	import type { Route } from '$lib/types';

	const route: Route = {
		getName: () => 'Someting went wrong',
		getPath: () => '/'
	};
</script>

<AuthLayout {route}>
	<div class="flex flex-col prose text-center">
		<h2>Someting went wrong</h2>
		<p>
			Sorry, something went wrong.<br />
			Please let us know if you think something is broken
		</p>
		<div>
			<a href="/" role="button" class="btn">Back to homepage</a>
		</div>
	</div>
</AuthLayout>
