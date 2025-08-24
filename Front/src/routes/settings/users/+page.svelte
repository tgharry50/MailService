<script lang="ts">
	// Imports
	import Icon from '@iconify/svelte';
	import plusIcon from '@iconify-icons/fa-solid/plus';
	import Dialog from '$lib/components/Dialog.svelte';
	import type { User } from '../../../interfaces/user';
	/////
	const users = [
		{ id: 1, name: 'Alice', email: 'alice@example.com' },
		{ id: 2, name: 'Bob', email: 'bob@example.com' },
		{ id: 3, name: 'Charlie', email: 'charlie@example.com' }
	];
	let dialogOptions = {
		open: false,
		title: '',
		message: '',
		confirmText: '',
		cancelText: '',
		dialogType: 'default',
		user_uuid: ''
	};
	// Remove user
	function openRemoveDialog(user) {
		dialogOptions = {
			open: true,
			title: 'Confirm Removal',
			message: `Remove user ${user.name}?`,
			confirmText: 'Remove',
			cancelText: 'Cancel',
			dialogType: 'delete',
			user_uuid: '123'
		};
	}
	// Edit user
	function openEditDialog(user) {
		dialogOptions = {
			open: true,
			title: 'Edit User',
			message: `Edit user ${user.name}?`,
			confirmText: 'Save',
			cancelText: 'Cancel',
			dialogType: 'edit',
			user_uuid: '123'
		};
	}
	// Add user
	function openAddUserDialog() {
		dialogOptions = {
			open: true,
			title: 'Add User',
			message: 'Add a new user here.',
			confirmText: 'Add',
			cancelText: 'Cancel',
			dialogType: 'add',
			user_uuid: ''
		};
	}
</script>

<svelte:head>
	<title>Settings | Users</title>
</svelte:head>

<div>
	<table class="user-table">
		<thead>
			<tr>
				<th colspan="4" style="padding:0; background:transparent; border:none;">
					<div class="table-users-row">
						<span>Table Users</span>
						<button class="add-user-btn" on:click={openAddUserDialog} title="Add user">
							<Icon icon={plusIcon} width="1.3em" />
						</button>
					</div>
				</th>
			</tr>
			<tr>
				<th>ID</th>
				<th>Name</th>
				<th>Email</th>
				<th class="options-col">Options</th>
			</tr>
		</thead>
		<tbody>
			{#each users as user}
				<tr>
					<td>{user.id}</td>
					<td>{user.name}</td>
					<td>{user.email}</td>
					<td class="options-col">
						<button on:click={() => openEditDialog(user)}>Edit</button>
						<button on:click={() => openRemoveDialog(user)}>Remove</button>
					</td>
				</tr>
			{/each}
		</tbody>
	</table>
</div>

<Dialog {...dialogOptions} />

<style>
	.user-table {
		width: 100%;
		border-collapse: collapse;
		margin-top: 2rem;
	}
	.user-table th,
	.user-table td {
		border: 1px solid #ddd;
		padding: 0.75rem 1rem;
		text-align: left;
	}
	.user-table th {
		background: #1481ca;
		color: #fff;
	}
	.user-table tr:nth-child(even) {
		background: #f6f8fa;
	}
	.user-table th.options-col,
	.user-table td.options-col {
		width: 1%;
		white-space: nowrap;
		padding-left: 0.5rem;
		padding-right: 0.5rem;
	}
	.options-col button {
		padding: 0.25rem 0.6rem;
		font-size: 0.95rem;
		margin-right: 0.2rem;
	}
	.add-user-btn {
		background: #1481ca;
		color: #fff;
		border: none;
		border-radius: 50%;
		width: 1.5em;
		height: 1.5em;
		display: flex;
		align-items: center;
		justify-content: center;
		font-size: 1.2em;
		cursor: pointer;
		transition: background 0.18s;
		box-shadow: 0 2px 8px rgba(0, 0, 0, 0.08);
	}
	.add-user-btn:hover {
		background: #0e5d99;
	}
	.table-users-row {
		display: flex;
		align-items: center;
		justify-content: space-between;
		width: 100%;
		font-size: 1.13rem;
		font-weight: 600;
		color: #1481ca;
		padding: 0.7rem 1.2rem 0.7rem 1.2rem;
		background: #f6f8fa;
		border-radius: 10px 10px 0 0;
		box-sizing: border-box;
		min-height: 2.8em;
	}
	.table-users-row span {
		display: flex;
		align-items: center;
		height: 100%;
		letter-spacing: 0.01em;
	}
	.add-user-btn {
		background: #1481ca;
		color: #fff;
		border: none;
		border-radius: 50%;
		width: 1.7em;
		height: 1.7em;
		display: flex;
		align-items: center;
		justify-content: center;
		font-size: 1em;
		cursor: pointer;
		transition:
			background 0.18s,
			box-shadow 0.18s;
		box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
		margin-left: 0.5em;
		padding: 0;
	}
	.add-user-btn :global(svg) {
		width: 1em;
		height: 1em;
	}
	.add-user-btn:hover {
		background: #0e5d99;
		box-shadow: 0 4px 16px rgba(20, 129, 202, 0.13);
	}
</style>
