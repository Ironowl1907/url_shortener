<script lang="ts">
	let { children, data } = $props();
</script>

<nav class="navbar">
	<div class="brand">
		<a href="/">URL Shortener</a>
	</div>

	<div class="nav-links">
		<a href="/">Home</a>
		{#if data.isAuthenticated}
			<a href="/dashboard">Dashboard</a>
			<a href="/profile">Profile</a>
		{:else}
			<a href="/login">Login</a>
			<a href="/register">Register</a>
		{/if}
	</div>
</nav>

{@render children()}

<style>
	.navbar {
		color: white;
		padding: 1rem 1.5rem;
		display: flex;
		justify-content: space-between;
		align-items: center;
	}

	.brand {
		font-size: 1.25rem;
		font-weight: bold;
	}

	.nav-links a {
		margin-left: 1rem;
		text-decoration: none;
		color: black;
	}

	.nav-links a:hover {
		text-decoration: underline;
	}
</style>
