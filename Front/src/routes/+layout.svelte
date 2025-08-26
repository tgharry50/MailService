<script lang="ts">
	import Header from '../layout/header.svelte';
	import Footer from '../layout/footer.svelte';
	import SideMenu from '../layout/sideMenu.svelte';
	import 'carbon-components-svelte/css/white.css';
</script>

<div class="page-container">
	<header>
		<Header />
	</header>
	<main>
		<SideMenu>
			<slot />
		</SideMenu>
	</main>
	<footer>
		<Footer />
	</footer>
</div>

<style>
	.page-container {
		min-height: 100vh;
		display: flex;
		flex-direction: column;
	}
	main {
		flex: 1;
	}
</style>
