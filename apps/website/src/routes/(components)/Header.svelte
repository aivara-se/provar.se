<script lang="ts">
	import { PUBLIC_PROVAR_APP_URL } from '$env/static/public';
	import { MenuIcon, UserIcon } from 'lucide-svelte';

	const links = [
		{ name: 'Home', href: '/' },
		{ name: 'Pricing', href: '/#pricing' },
		{ name: 'Docs', href: '/docs' }
	];
</script>

<header>
	<div class="navbar lg:px-10">
		<div class="navbar-start">
			<div class="dropdown">
				<div tabindex="0" role="button" class="btn btn-ghost lg:hidden">
					<MenuIcon />
				</div>
				<ul
					class="menu menu-lg dropdown-content mt-3 z-[1] p-2 shadow bg-base-100 rounded-box w-52"
				>
					{#each links as link}
						<li><a href={link.href}>{link.name}</a></li>
					{/each}
					<li><a href="/privacy">Privacy Policy</a></li>
					<li><a href="/terms">Terms of Use</a></li>
				</ul>
			</div>
			<a href="/" class="text-lg hidden lg:flex subpixel-antialiased tracking-wider">PROVAR</a>
		</div>
		<div class="navbar-center">
			<ul class="menu menu-sm menu-horizontal px-1 hidden lg:flex">
				{#each links as link}
					<li><a href={link.href}>{link.name}</a></li>
				{/each}
			</ul>
			<a href="/" class="text-lg lg:hidden subpixel-antialiased tracking-wider">PROVAR</a>
		</div>
		<div class="navbar-end">
			<a href={PUBLIC_PROVAR_APP_URL} class="btn btn-ghost btn-sm">
				<UserIcon class="w-5 h-5 lg:w-4 lg:h-4" />
				<span class="hidden lg:inline">Sign in</span>
			</a>
		</div>
	</div>
</header>
