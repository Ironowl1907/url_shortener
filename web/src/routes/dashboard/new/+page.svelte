<script>
	import { enhance } from '$app/forms';

	let loading = false;
</script>

<form
	method="POST"
	action="?/newUrl"
	use:enhance={() => {
		loading = true;
		return async ({ result, update }) => {
			loading = false;
			await update();
		};
	}}
>
	<label>
		Original URL
		<input name="url" type="url" required placeholder="https://example.com" disabled={loading} />
	</label>

	<label>
		Title
		<input
			name="title"
			type="text"
			placeholder="Optional title for your shortened URL"
			disabled={loading}
		/>
	</label>

	<label>
		Description
		<textarea name="description" placeholder="Optional description" disabled={loading}></textarea>
	</label>

	<label>
		<input name="ignore_response" type="checkbox" disabled={loading} />
		Ignore Response
	</label>

	<button type="submit" disabled={loading}>
		{loading ? 'Creating...' : 'Shorten URL'}
	</button>
</form>

<style>
	form {
		max-width: 500px;
		margin: 0 auto;
		padding: 2rem;
	}

	label {
		display: block;
		margin-bottom: 1rem;
		font-weight: bold;
	}

	input,
	textarea {
		display: block;
		width: 100%;
		padding: 0.5rem;
		margin-top: 0.25rem;
		border: 1px solid #ccc;
		border-radius: 4px;
		font-size: 1rem;
	}

	input[type='checkbox'] {
		display: inline;
		width: auto;
		margin-right: 0.5rem;
	}

	textarea {
		height: 100px;
		resize: vertical;
	}

	button {
		background-color: #007bff;
		color: white;
		padding: 0.75rem 1.5rem;
		border: none;
		border-radius: 4px;
		cursor: pointer;
		font-size: 1rem;
		width: 100%;
	}

	button:hover:not(:disabled) {
		background-color: #0056b3;
	}

	button:disabled {
		background-color: #6c757d;
		cursor: not-allowed;
	}

	input:disabled,
	textarea:disabled {
		background-color: #f8f9fa;
		cursor: not-allowed;
	}
</style>
