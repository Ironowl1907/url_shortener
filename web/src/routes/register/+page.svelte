<script>
	export let form;
</script>

<h1>Register</h1>

<div class="body">
	<form method="POST" action="?/register" class="Form">
		<input name="email" placeholder="Email" required />
		<input name="name" placeholder="Name" required />
		<input name="password" type="password" placeholder="Password" required />
		<button type="submit">Login</button>
		{#if form?.error}
			<p class="error">{form.error}</p>
		{/if}
	</form>
</div>

<style>
	.body {
		background: white;
		padding: 2rem;
		border-radius: 8px;
		box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
		width: 100%;
		max-width: 400px;
		margin: 0 auto;
	}

	/* Heading */
	h1 {
		text-align: center;
		color: #333;
		margin-bottom: 1.5rem;
		font-size: 1.8rem;
		font-weight: 600;
	}

	/* Input fields */
	input {
		width: 100%;
		padding: 0.75rem;
		margin-bottom: 1rem;
		border: 2px solid #e1e5e9;
		border-radius: 4px;
		font-size: 1rem;
		transition: border-color 0.3s ease;
		box-sizing: border-box;
	}

	input:focus {
		outline: none;
		border-color: #667eea;
	}

	input::placeholder {
		color: #9ca3af;
	}

	button {
		width: 100%;
		padding: 0.75rem;
		background: purple;
		color: white;
		border: none;
		border-radius: 4px;
		font-size: 1rem;
		font-weight: 600;
		cursor: pointer;
		transition:
			transform 0.2s ease,
			box-shadow 0.2s ease;
	}

	button:hover {
		transform: translateY(-1px);
		box-shadow: 0 4px 12px rgba(102, 126, 234, 0.4);
	}

	button:active {
		transform: translateY(0);
	}

	.error {
		color: #ef4444;
		background: #fef2f2;
		border: 1px solid #fecaca;
		padding: 0.75rem;
		border-radius: 4px;
		margin-top: 1rem;
		text-align: center;
		font-size: 0.9rem;
	}

	.footer {
		padding-top: 1.5rem;
		color: #667eea;
		text-decoration: none;
		font-weight: 500;
		transition: color 0.3s ease;
	}

	.footer a:hover {
		color: #764ba2;
		text-decoration: underline;
	}

	@media (max-width: 480px) {
		form {
			margin: 1rem;
			padding: 1.5rem;
		}

		h1 {
			font-size: 1.5rem;
		}
	}
</style>
