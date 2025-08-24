<script lang="ts">
	import { page } from '$app/stores';
	$: currentPage = $page.url.pathname;
</script>

<div class="side-menu-container">
	<nav>
		<ul>
			<li><a href="/" class:selected={currentPage === '/'}>Home</a></li>
			<li><a href="/jobs" class:selected={currentPage.startsWith('/jobs')}>Jobs</a></li>
			<li><a href="/mails" class:selected={currentPage.startsWith('/mails')}>Mails</a></li>
			<li><a href="/profile" class:selected={currentPage.startsWith('/profile')}>Profile</a></li>
			<li><a href="/settings" class:selected={currentPage.startsWith('/create')}>Settings</a></li>
			<li><a href="/contact">Login</a></li>
			<li><a href="/contact">Logout</a></li>
		</ul>
	</nav>
	<div class="content">
		<slot />
	</div>
</div>

<style scoped>
	:global(body) {
		margin: 0;
		padding: 0;
		background: #fff;
	} /* This shit actually covers the whole nav */
	.side-menu-container {
		display: flex;
		height: 100%;
		min-height: 100vh;
	}
	nav {
		width: 200px;
		background: #f0f0f0;
		padding: 1rem 0 1rem 1rem;
		min-height: 100vh;
		box-sizing: border-box;
	}
	.content {
		flex: 1;
		padding: 1rem;
		background: #fff;
	}
	ul {
		list-style: none;
		padding: 0;
		margin: 0;
	}
	li {
		margin-bottom: 1rem;
	}
	a {
		text-decoration: none;
		color: #333;
		font-weight: 900;
		font-size: larger;
	}
	a.selected {
		color: #ff0000;
	}
</style>
