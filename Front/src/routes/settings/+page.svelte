<script lang="ts">
	import Icon from '@iconify/svelte';
	import userIcon from '@iconify-icons/fa-solid/user';
	import cogIcon from '@iconify-icons/fa-solid/cog';
	import chartArea from '@iconify-icons/fa-solid/chart-area';
	import { goto } from '$app/navigation';
</script>

<svelte:head>
	<title>Create | Mail Service</title>
</svelte:head>

<div class="card-container">
	<div class="card card1" on:click={() => goto('/settings/users')}>
		<div class="icon-wrapper">
			<Icon icon={userIcon} width="2em" color="white" />
		</div>
		<div class="card-label">Users</div>
	</div>
	<div class="card card2" on:click={() => goto('/settings/connection')}>
		<div class="icon-wrapper">
			<Icon icon={cogIcon} width="2em" color="white" />
		</div>
		<div class="card-label">Connection settings</div>
	</div>
	<div class="card card4 disabled">
		<div class="icon-wrapper">
			<Icon icon={chartArea} width="2em" color="white" />
		</div>
		<div class="card-label">Placeholder</div>
	</div>
	<div class="card card3 disabled">
		<div class="icon-wrapper">
			<Icon icon={chartArea} width="2em" color="white" />
		</div>
		<div class="card-label">Placeholder</div>
	</div>
</div>

<style>
	.card-container {
		display: flex;
		gap: 1.5rem;
		justify-content: center;
		margin-top: 2rem;
		flex-wrap: wrap;
	}
	.card {
		background: #f0f0f0;
		border-radius: 10px;
		box-shadow: 0 2px 8px rgba(0, 0, 0, 0.08);
		padding: 2rem 1.5rem;
		min-width: 200px;
		min-height: 200px;
		display: flex;
		flex-direction: column;
		align-items: center;
		justify-content: center;
		font-size: 1.2rem;
		font-weight: 500;
	}
	.card:hover {
		cursor: pointer;
	}
	.icon-wrapper {
		display: flex;
		justify-content: center;
		align-items: center;
		margin-bottom: 1rem;
	}
	.card-label {
		text-align: center;
		font-size: 1.1rem;
	}
	.card1 {
		background: rgb(99, 224, 99);
	}
	.card2 {
		background: red;
	}
	.card3 {
		background: blue;
	}
	.card4 {
		background: rgb(0, 255, 221);
	}
	.disabled {
		opacity: 0.5;
		pointer-events: none;
		cursor: not-allowed;
		filter: grayscale(0.3);
	}
</style>
