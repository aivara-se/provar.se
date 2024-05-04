<script lang="ts">
	import exploreDesktopPng from '$lib/client/assets/screenshots/explore-desktop.png';
	import feedbackCSATDesktopPng from '$lib/client/assets/screenshots/feedback-csat-desktop.png';
	import feedbackCNPSDesktopPng from '$lib/client/assets/screenshots/feedback-cnps-desktop.png';
	import exploreMobilePng from '$lib/client/assets/screenshots/explore-mobile.png';
	import feedbackCSATMobilePng from '$lib/client/assets/screenshots/feedback-csat-mobile.png';
	import feedbackCNPSMobilePng from '$lib/client/assets/screenshots/feedback-cnps-mobile.png';
	import Header from './(components)/Header.svelte';
	import Footer from './(components)/Footer.svelte';
	import Carousel from './(components)/Carousel.svelte';

	const desktopImages = [exploreDesktopPng, feedbackCSATDesktopPng, feedbackCNPSDesktopPng];
	const mobileImages = [exploreMobilePng, feedbackCSATMobilePng, feedbackCNPSMobilePng];
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

	<section class="hidden md:flex">
		<Carousel images={desktopImages} />
	</section>
	<section class="md:hidden">
		<Carousel images={mobileImages} />
	</section>
</main>

<Footer />
