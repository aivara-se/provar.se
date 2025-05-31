<script lang="ts">
	import logoUrl from '$lib/client/assets/aivara.png';
</script>

<footer
	class="mt-16 p-8 flex flex-col items-center gap-4 mx-auto max-w-(--breakpoint-xl) md:flex-row md:justify-between"
>
	<aside class="flex items-center gap-3">
		<img width="32" src={logoUrl} alt="Provar" />
		<p class="text-sm">© 2024 Aivara<br />Empowering your software solutions!</p>
	</aside>
	<nav class="hidden md:flex gap-4 text-sm">
		<a href="/privacy" class="link link-hover">Privacy Policy</a>
		<a href="/cookies" class="link link-hover">Cookie Policy</a>
		<a href="/terms" class="link link-hover">Terms of Service</a>
	</nav>
</footer>
