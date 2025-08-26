<script lang="ts">
	// Imports
	import Icon from '@iconify/svelte';
	import saveIcon from '@iconify-icons/fa-solid/save';
	import userIcon from '@iconify-icons/fa-solid/user-alt';
	import {
		DataTable,
		Toolbar,
		ToolbarContent,
		ToolbarSearch,
		ToolbarMenu,
		ToolbarMenuItem,
		ToolbarBatchActions,
		Button
	} from 'carbon-components-svelte';
	import type { User } from '../../../interfaces/user';
	/////
	let selectedRowIds: number[] = [];
	let search = '';
	const users: User[] = [
		{
			id: 1,
			active: true,
			email: 'tessstuser@email.com',
			firstname: 'Tescxzt',
			lastname: 'Userasda',
			username: 'te21stuser',
			phone: '+123456789',
			uuid: '1234',
			password: '1234'
		},
		{
			id: 2,
			active: false,
			email: 'testuser@email.com',
			firstname: 'Test',
			lastname: 'User',
			username: 'testuser',
			phone: '+123456789',
			uuid: '12345',
			password: '1234'
		}
	];
	$: {
		console.log('Selectedlog', selectedRowIds);
		if (selectedRowIds.length === 1) {
			const selectedUser = users.find((user) => user.id === selectedRowIds[0]);
			if (selectedUser) {
				console.log('Selected user:', selectedUser.uuid);
			} else {
				console.log('Selected user: not found');
			}
		}
	}
	$: filteredUsers = users.filter(
		(user) =>
			search === '' ||
			user.firstname.toLowerCase().includes(search.toLowerCase()) ||
			user.lastname.toLowerCase().includes(search.toLowerCase()) ||
			user.username.toLowerCase().includes(search.toLowerCase()) ||
			user.email.toLowerCase().includes(search.toLowerCase())
	);
</script>

<svelte:head>
	<title>Settings | Users</title>
</svelte:head>

<div>
	<DataTable
		rows={filteredUsers}
		title="Users"
		headers={[
			{ key: 'id', value: 'ID' },
			{ key: 'firstname', value: 'First Name' },
			{ key: 'lastname', value: 'Last Name' },
			{ key: 'username', value: 'Username' },
			{ key: 'email', value: 'Email' },
			{ key: 'phone', value: 'Phone' },
			{ key: 'active', value: 'Active' }
		]}
		radio
		bind:selectedRowIds
		sortable
		zebra
	>
		<Toolbar>
			<ToolbarBatchActions>
				<Button>
					<pre>Save </pre>
					<Icon icon={saveIcon} />
				</Button>
			</ToolbarBatchActions>
			<ToolbarContent>
				<ToolbarSearch bind:value={search} />
				<ToolbarMenu>
					<ToolbarMenuItem primaryFocus>Restart all</ToolbarMenuItem>
					<ToolbarMenuItem>API documentation</ToolbarMenuItem>
					<ToolbarMenuItem hasDivider danger>Stop all</ToolbarMenuItem>
				</ToolbarMenu>
				<Button>Create balancer</Button>
			</ToolbarContent>
		</Toolbar>
		<svelte:fragment slot="cell" let:cell>
			{#if cell.key === 'active' && cell.value === true}
				<Icon icon={userIcon} color="green" />
			{:else if cell.key === 'active' && cell.value === false}
				<Icon icon={userIcon} color="red" />
			{:else}
				{cell.value}
			{/if}
		</svelte:fragment>
	</DataTable>
</div>

<style>
</style>
