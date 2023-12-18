<script lang="ts">
	import { applyAction, deserialize } from '$app/forms';
	import { invalidateAll } from '$app/navigation';
	import { page } from '$app/stores';
	import type { User } from '$lib/types';
	import type { ActionResult } from '@sveltejs/kit';
	import {
		Avatar,
		Button,
		Heading,
		Input,
		Label,
		Modal,
		P,
		Table,
		TableBody,
		TableBodyCell,
		TableBodyRow,
		TableHead,
		TableHeadCell
	} from 'flowbite-svelte';

	$: members = $page.data.members || [];
	$: invitations = $page.data.invitations || [];

	let isInviteModalOpen = false;
	let isDetailsModalOpen = false;

	let invitedUserName = '';
	let invitedUserEmail = '';

	let selectedMember: User;

	function selectMember(member: User) {
		selectedMember = member;
		isDetailsModalOpen = true;
	}

	async function inviteMember() {
		const data = new FormData();
		data.set('name', invitedUserName);
		data.set('email', invitedUserEmail);
		const response = await fetch('?/inviteMember', {
			method: 'post',
			body: data
		});
		const result: ActionResult = deserialize(await response.text());
		if (result.type === 'success') {
			await invalidateAll();
		}
		applyAction(result);
	}

	async function revokeMembership(id: string) {
		const data = new FormData();
		data.set('id', id);
		const response = await fetch('?/revokeMembership', {
			method: 'post',
			body: data
		});
		const result: ActionResult = deserialize(await response.text());
		if (result.type === 'success') {
			await invalidateAll();
		}
		applyAction(result);
	}
</script>

<section>
	<Heading customSize="text-xl font-semibold">Organization members</Heading>
	<P class="mt-6">Below is the list of individuals currently associated with this organization.</P>

	{#if members.length > 0}
		<section class="current-credentials mt-6">
			<Table hoverable>
				<TableHead>
					<TableHeadCell>Avatar</TableHeadCell>
					<TableHeadCell>Name</TableHeadCell>
					<TableHeadCell>Email</TableHeadCell>
					<TableHeadCell></TableHeadCell>
				</TableHead>
				<TableBody>
					{#each members as item}
						<TableBodyRow>
							<TableBodyCell tdClass="w-24">
								<div class="flex justify-center rounded">
									<Avatar src={item.image} size="sm" rounded />
								</div>
							</TableBodyCell>
							<TableBodyCell>{item.name || ''}</TableBodyCell>
							<TableBodyCell>{item.email}</TableBodyCell>
							<TableBodyCell>
								<Button pill size="xs" color="light" on:click={() => selectMember(item)}>
									Details
								</Button>
							</TableBodyCell>
						</TableBodyRow>
					{/each}
				</TableBody>
			</Table>
		</section>
	{/if}
</section>

{#if invitations.length > 0}
	<section class="mt-6">
		<Heading customSize="font-semibold">Invited members</Heading>
		<P class="mt-6">
			Below is the list of individuals we have sent invitations to join this organization.
		</P>
	</section>

	<section class="current-credentials mt-6">
		<Table hoverable>
			<TableHead>
				<TableHeadCell>Name</TableHeadCell>
				<TableHeadCell>Email</TableHeadCell>
				<TableHeadCell />
			</TableHead>
			<TableBody>
				{#each invitations as item}
					<TableBodyRow>
						<TableBodyCell>{item.name}</TableBodyCell>
						<TableBodyCell>{item.email}</TableBodyCell>
						<TableBodyCell>
							<Button pill size="xs" color="light" on:click={() => null}>Details</Button>
						</TableBodyCell>
					</TableBodyRow>
				{/each}
			</TableBody>
		</Table>
	</section>
{/if}

<section class="mt-6">
	<Button color="light" size="xs" on:click={() => (isInviteModalOpen = true)}>
		+ Invite member
	</Button>
</section>

<Modal title="Invite member" bind:open={isInviteModalOpen} autoclose>
	<Label for="name" class="block mb-1">Name:</Label>
	<Input id="name" required bind:value={invitedUserName} />
	<Label for="email" class="block mb-1">Email:</Label>
	<Input id="email" required bind:value={invitedUserEmail} />
	<svelte:fragment slot="footer">
		<Button size="sm" color="primary" on:click={inviteMember}>Invite</Button>
	</svelte:fragment>
</Modal>

<Modal title={selectedMember?.name || ' '} bind:open={isDetailsModalOpen} autoclose>
	<P class="mb-2">
		<span class="font-semibold">Email:</span>
		<code class="block break-words text-xs mt-2 bg-gray-100 p-2 rounded">
			{selectedMember?.email}
		</code>
	</P>
	<svelte:fragment slot="footer">
		<Button size="sm" color="red" on:click={() => revokeMembership(selectedMember.id)}>
			Revoke
		</Button>
	</svelte:fragment>
</Modal>
