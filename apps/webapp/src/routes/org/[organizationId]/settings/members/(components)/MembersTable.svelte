<script lang="ts">
	import { ArrowBigRightIcon } from 'lucide-svelte';
	import type { User } from '$lib/types';

	export let members: User[] = [];
</script>

{#if members.length}
	<table class="table">
		<thead>
			<tr>
				<th></th>
				<th class="hidden md:table-cell">Name</th>
				<th>Email</th>
				<th></th>
			</tr>
		</thead>
		<tbody>
			{#each members as member}
				<tr>
					<td>
						<div class="avatar">
							<div class="w-8 rounded-full">
								<img src={member.image} alt={member.name} />
							</div>
						</div>
					</td>
					<td class="hidden md:table-cell">{member.name}</td>
					<td>{member.email}</td>
					<td class="text-right">
						<a href="./members/{member.id}" role="button" class="btn btn-sm btn-ghost">
							<ArrowBigRightIcon class="w-4 h-4" />
						</a>
					</td>
				</tr>
			{/each}
		</tbody>
	</table>
{/if}
