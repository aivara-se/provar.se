<script lang="ts">
	import exploreDesktopPng from '$lib/client/assets/screenshots/explore-desktop.png';
	import feedbackCSATDesktopPng from '$lib/client/assets/screenshots/feedback-csat-desktop.png';
	import feedbackCNPSDesktopPng from '$lib/client/assets/screenshots/feedback-cnps-desktop.png';
	import feedbackIllustration from '$lib/client/assets/illustrations/feedback.svg';
	import analyticsIllustration from '$lib/client/assets/illustrations/analytics.svg';
	import Header from './(components)/Header.svelte';
	import Footer from './(components)/Footer.svelte';
	import Carousel from './(components)/Carousel.svelte';

	const MD_IMAGES = [exploreDesktopPng, feedbackCSATDesktopPng, feedbackCNPSDesktopPng];
</script>

<Header />

<main id="hero" class="px-4 mx-auto max-w-screen-xl md:p-10">
	<section class="hero py-24">
		<div class="hero-content text-center subpixel-antialiased">
			<div class="max-w-xl">
				<h1 class="text-5xl font-bold">Want happier customers?</h1>
				<p class="py-6 text-lg">
					Gain precision, clarity, and actionable data to drive customer satisfaction
				</p>
				<button class="btn btn-wide btn-info">Get Started in Minutes!</button>
			</div>
		</div>
	</section>
	<section class="hidden md:flex">
		<Carousel images={MD_IMAGES} />
	</section>
</main>

<main id="features" class="flex flex-col gap-16 p-4 mx-auto max-w-screen-xl md:p-10 md:gap-24">
	<section class="flex flex-col gap-6 md:flex-row-reverse md:gap-8">
		<aside class="md:flex-auto md:w-32 md:mt-12">
			<div class="bg-gray-800 rounded-xl p-4">
				<img src={feedbackIllustration} alt="Collect feedback illustration" />
			</div>
		</aside>
		<article class="md:flex-auto md:w-64 flex flex-col gap-4">
			<h2 class="text-xl md:text-2xl font-semibold text-center md:text-left">
				Effortless Feedback Collection
			</h2>
			<div class="border border-dashed border-slate-700 rounded-xl p-4">
				<h3 class="text-lg font-semibold">Many Feedback Channels</h3>
				<p class="text-sm mt-2">
					Provar allows you to collect feedback through various channels, including surveys, in-app
					feedback forms and email requests.
				</p>
			</div>
			<div class="border border-dashed border-slate-700 rounded-xl p-4">
				<h3 class="text-lg font-semibold">NPS & CSAT Automation</h3>
				<p class="text-sm mt-2">
					Effortlessly automate the calculation of key customer satisfaction metrics like Net
					Promoter Score (NPS) and Customer Satisfaction Score (CSAT)
				</p>
			</div>
			<div class="border border-dashed border-slate-700 rounded-xl p-4">
				<h3 class="text-lg font-semibold">Provar Feedback Collector</h3>
				<p class="text-sm mt-2">
					Provar extends its reach to physical locations with the Provar Feedback Collector,
					allowing for seamless feedback gathering on-site.
				</p>
			</div>
		</article>
	</section>

	<section class="flex flex-col gap-6 md:flex-row-reverse md:gap-8">
		<aside class="md:flex-auto md:w-32 md:mt-12">
			<div class="bg-gray-800 rounded-xl p-4">
				<img src={analyticsIllustration} alt="Analyze feedback illustration" />
			</div>
		</aside>
		<article class="md:flex-auto md:w-64 flex flex-col gap-4">
			<h2 class="text-xl md:text-2xl font-semibold text-center md:text-left">
				Powerful Feedback Analysis
			</h2>
			<div class="border border-dashed border-slate-700 rounded-xl p-4">
				<h3 class="text-lg font-semibold">Real-time Analytics Dashboard</h3>
				<p class="text-sm mt-2">
					Gain real-time insights into your customer feedback with a user-friendly dashboard that
					displays key metrics and trends.
				</p>
			</div>
			<div class="border border-dashed border-slate-700 rounded-xl p-4">
				<h3 class="text-lg font-semibold">Advanced Sentiment Analysis</h3>
				<p class="text-sm mt-2">
					Go beyond simple text by leveraging sentiment analysis to understand the emotional tone
					behind user feedback.
				</p>
			</div>
			<div class="border border-dashed border-slate-700 rounded-xl p-4">
				<h3 class="text-lg font-semibold">Clear and Actionable Reporting</h3>
				<p class="text-sm mt-2">
					Generate customizable reports that surface key themes and actionable insights from your
					customer feedback data.
				</p>
			</div>
		</article>
	</section>
</main>

<Footer />
