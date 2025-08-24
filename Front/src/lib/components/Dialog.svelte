<!-- svelte-ignore a11y_click_events_have_key_events -->
<!-- svelte-ignore a11y_no_noninteractive_element_interactions -->
<script lang="ts">
	// Props
	let { open, title, message, confirmText, cancelText, dialogType, user_uuid } = $props();
	// Imports
	import { Axios } from 'axios';
	import { globalValues } from '../../interfaces/global';
	import type { UserAdd, UserEdit } from '../../interfaces/user';
	// Ensure userAdd is always defined for form binding
	let userAdd: UserAdd = {
		firstname: '',
		lastname: '',
		username: '',
		password: '',
		email: '',
		phone: ''
	};
	let userEdit: UserEdit = {
		firstname: '',
		lastname: '',
		username: '',
		password: '',
		email: '',
		phone: '',
		uuid: '',
		active: false
	};

	let userDelete = $state(false);
	let userDeleteText = '';

	function handleConfirm() {
		const form = new FormData();
		if (dialogType === 'add') {
			form.append('firstname', userAdd.firstname);
			form.append('lastname', userAdd.lastname);
			form.append('username', userAdd.username);
			form.append('password', userAdd.password);
			form.append('email', userAdd.email);
			form.forEach((element) => {
				console.log(element);
			});
			console.log(globalValues.apiUrl);
		} else if (dialogType === 'edit') {
			form.append('uuid', user_uuid as string);
			form.append('firstname', userEdit?.firstname as string);
			form.append('lastname', userEdit?.lastname as string);
			form.append('username', userEdit?.username as string);
			form.append('password', userEdit?.password as string);
			form.append('email', userEdit?.email as string);
			form.append('active', userEdit?.active !== undefined ? String(userEdit.active) : '');
			form.forEach((element) => {
				console.log(element);
			});
		} else if (dialogType === 'delete') {
			form.append('uuid', user_uuid as string);
			form.forEach((element) => {
				console.log(element);
			});
		}
	}

	function handleCancel() {
		open = false;
	}
</script>

{#if open}
	<div
		class="dialog-backdrop"
		role="dialog"
		tabindex="-1"
		aria-modal="true"
		onclick={handleCancel}
		onkeydown={(e) => {
			if (e.key === 'Escape') handleCancel();
		}}
	>
		<div class="dialog" role="document" onclick={(e) => e.stopPropagation()}>
			<h2>{title}</h2>
			<p>{message}</p>
			{#if dialogType === 'add'}
				<form
					class="dialog-form"
					onsubmit={(e) => {
						e.preventDefault();
						handleConfirm();
					}}
				>
					<div class="form-group">
						<label for="firstname">First Name</label>
						<input id="firstname" type="text" bind:value={userAdd.firstname} required />
					</div>
					<div class="form-group">
						<label for="lastname">Last Name</label>
						<input id="lastname" type="text" bind:value={userAdd.lastname} required />
					</div>
					<div class="form-group">
						<label for="username">Username</label>
						<input id="username" type="text" bind:value={userAdd.username} required />
					</div>
					<div class="form-group">
						<label for="email">Email</label>
						<input id="email" type="email" bind:value={userAdd.email} required />
					</div>
					<div class="form-group">
						<label for="password">Password</label>
						<input id="password" type="password" bind:value={userAdd.password} required />
					</div>
				</form>
			{:else if dialogType === 'edit'}
				<form
					class="dialog-form"
					onsubmit={(e) => {
						e.preventDefault();
						handleConfirm();
					}}
				>
					<div class="form-group">
						<label for="edit-firstname">First Name</label>
						<input id="edit-firstname" type="text" bind:value={userEdit.firstname} required />
					</div>
					<div class="form-group">
						<label for="edit-lastname">Last Name</label>
						<input id="edit-lastname" type="text" bind:value={userEdit.lastname} required />
					</div>
					<div class="form-group">
						<label for="edit-username">Username</label>
						<input id="edit-username" type="text" bind:value={userEdit.username} required />
					</div>
					<div class="form-group">
						<label for="edit-email">Email</label>
						<input id="edit-email" type="email" bind:value={userEdit.email} required />
					</div>
					<div class="form-group">
						<label for="edit-password">Password</label>
						<input id="edit-password" type="password" bind:value={userEdit.password} required />
					</div>
					<div class="form-group">
						<label for="edit-active">Active</label>
						<input id="edit-active" type="checkbox" bind:checked={userEdit.active} />
					</div>
				</form>
			{:else if dialogType === 'delete'}
				<div class="form-group" style="flex-direction: row; align-items: center; gap: 0.5rem;">
					<input id="delete-confirm" type="checkbox" bind:checked={userDelete} />
					<label for="delete-confirm" style="margin: 0; font-weight: 400;">
						Yes I want to delete
					</label>
				</div>
			{:else}
				<p>Default</p>
			{/if}
			<div class="dialog-actions">
				<button class="cancel" type="button" onclick={handleCancel}>{cancelText}</button>
				<button class="confirm" type="button" onclick={handleConfirm}>{confirmText}</button>
			</div>
		</div>
	</div>
{/if}

<style scoped>
	.dialog-backdrop {
		position: fixed;
		top: 0;
		left: 0;
		width: 100vw;
		height: 100vh;
		background: rgba(0, 0, 0, 0.35);
		display: flex;
		align-items: center;
		justify-content: center;
		z-index: 1000;
		transition: background 0.2s;
	}
	.dialog {
		background: #fff;
		padding: 2.5rem 2rem 2rem 2rem;
		border-radius: 12px;
		box-shadow: 0 8px 32px rgba(0, 0, 0, 0.18);
		min-width: 340px;
		max-width: 95vw;
		animation: dialog-pop 0.18s cubic-bezier(0.4, 1.4, 0.6, 1) both;
	}
	@keyframes dialog-pop {
		0% {
			transform: scale(0.95);
			opacity: 0;
		}
		100% {
			transform: scale(1);
			opacity: 1;
		}
	}
	.dialog h2 {
		margin: 0 0 1rem 0;
		font-size: 1.35rem;
		color: #1481ca;
		font-weight: 600;
		letter-spacing: 0.01em;
	}
	.dialog p {
		margin: 0 0 1.5rem 0;
		color: #333;
		font-size: 1.05rem;
	}
	.dialog-actions {
		display: flex;
		justify-content: flex-end;
		gap: 1rem;
		margin-top: 1.5rem;
	}
	button {
		padding: 0.5rem 1.4rem;
		border: none;
		border-radius: 5px;
		font-size: 1rem;
		cursor: pointer;
		font-weight: 500;
		transition:
			background 0.18s,
			color 0.18s;
		outline: none;
	}
	button.cancel {
		background: #f3f3f3;
		color: #444;
		border: 1px solid #ddd;
	}
	button.confirm {
		background: #1481ca;
		color: #fff;
		border: 1px solid #1481ca;
	}
	button.cancel:hover {
		background: #e0e0e0;
		color: #222;
	}
	button.confirm:hover {
		background: #0e5d99;
		border-color: #0e5d99;
	}

	/* Form styles for add dialog */
	.dialog-form {
		display: flex;
		flex-direction: column;
		gap: 1rem;
	}
	.form-group {
		display: flex;
		flex-direction: column;
		gap: 0.3rem;
	}
	.form-group label {
		font-size: 0.98rem;
		color: #1481ca;
		font-weight: 500;
	}
	.form-group input {
		padding: 0.5rem 0.8rem;
		border: 1px solid #ccc;
		border-radius: 4px;
		font-size: 1rem;
	}
</style>
