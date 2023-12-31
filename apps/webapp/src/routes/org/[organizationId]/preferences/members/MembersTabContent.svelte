<script lang="ts">
	import { page } from '$app/stores';
	import type { Invitation, User } from '$lib/types';
	import {
		Avatar,
		Button,
		Heading,
		P,
		Table,
		TableBody,
		TableBodyCell,
		TableBodyRow,
		TableHead,
		TableHeadCell
	} from 'flowbite-svelte';
	import CreateInviteModal from './_components/CreateInviteModal.svelte';
	import InvitationDetailsModal from './_components/InvitationDetailsModal.svelte';
	import MemberDetailsModal from './_components/MemberDetailsModal.svelte';

	$: members = $page.data.members || [];
	$: invitations = $page.data.invitations || [];

	let isInviteMemberModalOpen = false;
	let isMemberDetailsModalOpen = false;
	let isInviteDetailsModalOpen = false;

	let selectedMember: User;
	let selectedInvitation: Invitation;

	function selectMember(member: User) {
		selectedMember = member;
		isMemberDetailsModalOpen = true;
	}

	function selectInvitation(invitation: Invitation) {
		selectedInvitation = invitation;
		isInviteDetailsModalOpen = true;
	}

	function getTimeString(ts: number) {
		const date = new Date(ts);
		return `${date.toLocaleDateString()} ${date.toLocaleTimeString()}`;
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
				<TableHeadCell>Created At</TableHeadCell>
				<TableHeadCell />
			</TableHead>
			<TableBody>
				{#each invitations as item}
					<TableBodyRow>
						<TableBodyCell>{item.name}</TableBodyCell>
						<TableBodyCell>{item.email}</TableBodyCell>
						<TableBodyCell>{getTimeString(item.createdAt)}</TableBodyCell>
						<TableBodyCell>
							<Button pill size="xs" color="light" on:click={() => selectInvitation(item)}>
								Details
							</Button>
						</TableBodyCell>
					</TableBodyRow>
				{/each}
			</TableBody>
		</Table>
	</section>
{/if}

<section class="mt-6">
	<Button color="light" size="xs" on:click={() => (isInviteMemberModalOpen = true)}>
		+ Invite member
	</Button>
</section>

<CreateInviteModal bind:isOpen={isInviteMemberModalOpen} />
<MemberDetailsModal member={selectedMember} bind:isOpen={isMemberDetailsModalOpen} />
<InvitationDetailsModal invitation={selectedInvitation} bind:isOpen={isInviteDetailsModalOpen} />
