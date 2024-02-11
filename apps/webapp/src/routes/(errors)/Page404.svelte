<script lang="ts">
	import { AuthLayout } from '$lib/client/layout';
	import type { Route } from '$lib/types';

	const route: Route = {
		getName: () => 'Page not found',
		getPath: () => '/'
	};
</script>

<AuthLayout {route}>
	<div class="flex flex-col prose text-center">
		<h2>Page not found</h2>
		<p>
			Sorry, we can't find that page.<br />
			Please let us know if you think something is broken
		</p>
		<div>
			<a href="/" role="button" class="btn">Back to homepage</a>
		</div>
	</div>
</AuthLayout>
