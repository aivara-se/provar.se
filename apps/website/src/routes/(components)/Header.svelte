<script lang="ts">
	import { PUBLIC_PROVAR_APP_URL } from '$env/static/public';
	import { MenuIcon, UserIcon } from 'lucide-svelte';

	const links = [
		{ name: 'Home', href: '/' },
		{ name: 'Pricing', href: '/#pricing' },
		{ name: 'Readme', href: '/docs' }
	];
</script>

<header>
	<div class="navbar md:px-10">
		<div class="navbar-start">
			<div class="dropdown">
				<div tabindex="0" role="button" class="btn btn-ghost md:hidden">
					<MenuIcon />
				</div>
				<ul
					class="menu menu-lg dropdown-content mt-3 z-[1] p-2 shadow bg-base-100 rounded-box w-52"
				>
					{#each links as link}
						<li><a href={link.href}>{link.name}</a></li>
					{/each}
					<li class="h-1 divider"></li>
					<li><a href="/privacy" class="link link-hover">Privacy Policy</a></li>
					<li><a href="/cookies" class="link link-hover">Cookie Policy</a></li>
					<li><a href="/terms" class="link link-hover">Terms of Service</a></li>
				</ul>
			</div>
			<a href="/" class="text-lg hidden md:flex subpixel-antialiased tracking-wider">PROVAR</a>
		</div>
		<div class="navbar-center">
			<ul class="menu menu-sm menu-horizontal px-1 hidden md:flex">
				{#each links as link}
					<li><a href={link.href}>{link.name}</a></li>
				{/each}
			</ul>
			<a href="/" class="text-lg md:hidden subpixel-antialiased tracking-wider">PROVAR</a>
		</div>
		<div class="navbar-end">
			<a href={PUBLIC_PROVAR_APP_URL} class="btn btn-ghost btn-sm">
				<UserIcon class="w-5 h-5 md:w-4 md:h-4" />
				<span class="hidden md:inline">Sign in</span>
			</a>
		</div>
	</div>
</header>
