<script lang="ts">
	import { ChevronRightIcon } from 'lucide-svelte';
	import type { Invitation } from '$lib/client/types';

	export let invitations: Invitation[] = [];
</script>

{#if invitations.length}
	<h3 class="font-medium mt-8">Invitations</h3>
	<table class="table mt-4">
		<thead>
			<tr>
				<th></th>
				<th class="hidden md:table-cell">Name</th>
				<th>Email</th>
				<th></th>
			</tr>
		</thead>
		<tbody>
			{#each invitations as invitation}
				<tr>
					<td>
						<div class="avatar">
							<div class="w-8 rounded-full">
								<img src={invitation.avatar} alt={invitation.name} />
							</div>
						</div>
					</td>
					<td class="hidden md:table-cell">{invitation.name}</td>
					<td>{invitation.email}</td>
					<td class="text-right">
						<a
							href="./members/invitations/{invitation.id}"
							role="button"
							class="btn btn-sm btn-ghost"
						>
							<ChevronRightIcon class="w-4 h-4" />
						</a>
					</td>
				</tr>
			{/each}
		</tbody>
	</table>
{/if}
