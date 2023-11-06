<div class="wrapper">
	<aside class="sidebar">
		<section class="sidebar-top"><slot name="sidebar-top" /></section>
		<section class="sidebar-bottom"><slot name="sidebar-bottom" /></section>
	</aside>
	<main class="content"><slot /></main>
</div>

<style>
	.wrapper {
		display: flex;
		flex-direction: row;
		align-items: center;
		justify-content: center;
		height: 100vh;
	}

	.sidebar {
		display: flex;
		padding: 20px;
		width: 200px;
		height: 100%;
		background: #f0f0f0;
		flex-direction: column;
		box-sizing: border-box;
		border-right: solid 1px #e0e0e0;
	}

	.sidebar-top {
		display: flex;
		flex: 1;
		flex-direction: column;
	}

	.sidebar-bottom {
		display: flex;
		align-items: center;
		flex-direction: column;
	}

	.content {
		display: flex;
		flex: 1;
		height: 100%;
		flex-direction: column;
		box-sizing: border-box;
		padding: 20px;
	}
</style>
