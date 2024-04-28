<script lang="ts">
	import exploreDesktopPng from '$lib/client/assets/screenshots/explore-desktop.png';
	import feedbackCSATDesktopPng from '$lib/client/assets/screenshots/feedback-csat-desktop.png';
	import feedbackCNPSDesktopPng from '$lib/client/assets/screenshots/feedback-cnps-desktop.png';
	import Header from './(components)/Header.svelte';
	import Footer from './(components)/Footer.svelte';
</script>

<Header />

<main class="min-h-screen">
	<section class="hero py-24">
		<div class="hero-content text-center subpixel-antialiased">
			<div class="max-w-xl">
				<h1 class="text-5xl font-bold">Maximize customer satisfaction & dialogue</h1>
				<p class="py-6 text-lg">
					Precision, clarity, and actionable data. One streamlined platform.
				</p>
				<button class="btn btn-success">Get Started!</button>
			</div>
		</div>
	</section>

	<section class="hidden lg:flex items-center justify-center">
		<div class="carousel w-5/6">
			<div id="slide1" class="carousel-item relative w-full">
				<img src={exploreDesktopPng} alt="Dashboard" class="w-full rounded-xl shadow" />
				<div
					class="absolute flex justify-between transform -translate-y-1/2 left-5 right-5 top-1/2"
				>
					<a href="#slide3" class="btn btn-circle btn-outline opacity-50">❮</a>
					<a href="#slide2" class="btn btn-circle btn-outline opacity-50">❯</a>
				</div>
			</div>
			<div id="slide2" class="carousel-item relative w-full">
				<img src={feedbackCSATDesktopPng} alt="Dashboard" class="w-full rounded-xl shadow" />
				<div
					class="absolute flex justify-between transform -translate-y-1/2 left-5 right-5 top-1/2"
				>
					<a href="#slide1" class="btn btn-circle btn-outline opacity-50">❮</a>
					<a href="#slide3" class="btn btn-circle btn-outline opacity-50">❯</a>
				</div>
			</div>
			<div id="slide3" class="carousel-item relative w-full">
				<img src={feedbackCNPSDesktopPng} alt="Dashboard" class="w-full rounded-xl shadow" />
				<div
					class="absolute flex justify-between transform -translate-y-1/2 left-5 right-5 top-1/2"
				>
					<a href="#slide2" class="btn btn-circle btn-outline opacity-50">❮</a>
					<a href="#slide1" class="btn btn-circle btn-outline opacity-50">❯</a>
				</div>
			</div>
		</div>
	</section>
</main>

<Footer />

<style>
	:global(html) {
		background: url(data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAAcAAAAHCAIAAABLMMCEAAAAJklEQVQIW2NgYGAQggEGOEDmINgQlqioKBZRTDY2c9GAOBhgFwUAK5UD3J7tt1sAAAAASUVORK5CYII=);
	}
</style>
