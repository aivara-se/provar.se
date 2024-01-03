<script lang="ts">
	export let text = '';
</script>

<div class="separator">
	<span class="bg-white px-2 text-sm text-gray-500">{text}</span>
</div>

<style>
	.separator {
		display: flex;
		align-items: center;
		justify-content: center;
		margin: 25px 0;
		width: 160px;
		height: 1px;
		background: #d0d0d0;
	}
</style>
