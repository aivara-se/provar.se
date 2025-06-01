<script lang="ts">
	import { MenuIcon, ChevronRightIcon, ChartColumnBigIcon, SettingsIcon } from 'lucide-svelte';

	const items = [
		{ name: 'Dashboard', href: '/admin', icon: ChartColumnBigIcon },
		{ name: 'Preferences', href: '/admin/settings', icon: SettingsIcon }
	];
</script>

<div class="navbar bg-base-100">
	<div class="navbar-start">
		<div class="dropdown">
			<div tabindex="0" role="button" class="btn btn-ghost btn-circle">
				<MenuIcon class="h-5 w-5" />
			</div>
			<ul class="menu dropdown-content mt-3 z-1 p-2 shadow bg-base-100 rounded-box w-40">
				{#each items as item}
					<li>
						<a href={item.href} class="active:bg-transparent">
							{#if item.icon}
								<svelte:component this={item.icon} class="w-4 h-4" />
							{/if}
							<span class="antialiased">
								{item.name}
							</span>
						</a>
					</li>
				{/each}
			</ul>
		</div>
	</div>
	<div class="navbar-center">
		<a href="/admin" class="btn btn-ghost text-xl">Provar</a>
	</div>
	<div class="navbar-end">
		<a href='/client' class="btn btn-ghost btn-circle">
			<ChevronRightIcon class="h-5 w-5" />
		</a>
	</div>
</div>
