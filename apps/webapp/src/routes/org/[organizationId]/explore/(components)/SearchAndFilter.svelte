<script lang="ts">
	import { SearchIcon } from 'lucide-svelte';
</script>

<div class="join flex shadow">
	<input class="input input-sm join-item flex-1" placeholder="Search ..." />
	<button class="btn btn-sm join-item">
		<SearchIcon class="w-3 h-3" />
	</button>
</div>
