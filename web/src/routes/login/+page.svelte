<script>
	export let form;
</script>

<h1>Login</h1>

<form method="POST" action="?/login">
	<input name="email" placeholder="Email" required />
	<input name="password" type="password" placeholder="Password" required />
	<button type="submit">Login</button>
	{#if form?.error}
		<p class="error">{form.error}</p>
	{/if}
</form>
