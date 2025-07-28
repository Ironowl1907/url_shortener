<script lang="ts">
	import type { ShortenedUrl } from '$lib/types';
	import { enhance } from '$app/forms';

	let allProps = $props();
	const user = $derived(allProps.data.user);
	const urls: ShortenedUrl[] = $derived(allProps.data.shortenedUrls);
</script>

<div class="urls-container">
	<h2>Your Shortened URLs</h2>
	{#if urls && urls.length > 0}
		<div class="urls-grid">
			{#each urls as url}
				{console.log(url)}
				<div class="url-card">
					<div class="url-info">
						<h3 class="short-url">
							<a href="http://localhost:8080/{url.ShortCode}" target="_blank">
								localhost:8080/{url.ShortCode}
							</a>
						</h3>
						<p class="original-url">
							<strong>Original:</strong>
							<a href={url.OriginalURL} target="_blank" title={url.OriginalURL}>
								{url.OriginalURL.length > 50
									? url.OriginalURL.substring(0, 50) + '...'
									: url.OriginalURL}
							</a>
						</p>
						<div class="url-meta">
							<span class="created-date">
								Created: {url.CreatedAt.toLocaleDateString()}
							</span>
							<span class="click-count">
								<!-- Clicks: {url.click_count || 0} -->
							</span>
						</div>
					</div>
					<div class="url-actions">
						<button
							class="copy-btn"
							onclick={() =>
								navigator.clipboard.writeText(`http://localhost:8080/${url.ShortCode}`)}
						>
							Copy
						</button>
						<form method="POST" action="?/delete" use:enhance>
							<input type="hidden" name="id" value={url.id} />
							<button type="submit" class="delete-btn"> Delete </button>
						</form>
					</div>
				</div>
			{/each}
		</div>
	{:else if urls === null}
		<div class="no-urls">
			<p>Unable to load URLs. Please check your authentication.</p>
		</div>
	{:else}
		<div class="no-urls">
			<p>You haven't created any shortened URLs yet.</p>
			<button class="create-btn">Create Your First URL</button>
		</div>
	{/if}
</div>

<style>
	.urls-container {
		max-width: 1200px;
		margin: 2rem auto;
		padding: 0 1rem;
	}
	.urls-container h2 {
		color: #333;
		margin-bottom: 1.5rem;
		border-bottom: 2px solid #e0e0e0;
		padding-bottom: 0.5rem;
	}
	.urls-grid {
		display: grid;
		gap: 1.5rem;
		grid-template-columns: repeat(auto-fill, minmax(400px, 1fr));
	}
	.url-card {
		background: white;
		border: 1px solid #e0e0e0;
		border-radius: 8px;
		padding: 1.5rem;
		box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
		transition: box-shadow 0.2s ease;
	}
	.url-card:hover {
		box-shadow: 0 4px 8px rgba(0, 0, 0, 0.15);
	}
	.url-info {
		margin-bottom: 1rem;
	}
	.short-url {
		margin: 0 0 0.5rem 0;
		font-size: 1.1rem;
	}
	.short-url a {
		color: #2563eb;
		text-decoration: none;
		font-weight: 600;
	}
	.short-url a:hover {
		text-decoration: underline;
	}
	.original-url {
		margin: 0.5rem 0;
		color: #666;
		font-size: 0.9rem;
	}
	.original-url a {
		color: #666;
		text-decoration: none;
	}
	.original-url a:hover {
		color: #2563eb;
		text-decoration: underline;
	}
	.url-meta {
		display: flex;
		justify-content: space-between;
		margin-top: 0.5rem;
		font-size: 0.8rem;
		color: #888;
	}
	.url-actions {
		display: flex;
		gap: 0.5rem;
		align-items: flex-start;
	}
	.url-actions form {
		margin: 0;
	}
	.copy-btn,
	.delete-btn,
	.create-btn {
		padding: 0.5rem 1rem;
		border: none;
		border-radius: 4px;
		cursor: pointer;
		font-size: 0.9rem;
		transition: background-color 0.2s ease;
	}
	.copy-btn {
		background-color: #10b981;
		color: white;
	}
	.copy-btn:hover {
		background-color: #059669;
	}
	.delete-btn {
		background-color: #ef4444;
		color: white;
	}
	.delete-btn:hover {
		background-color: #dc2626;
	}
	.create-btn {
		background-color: #2563eb;
		color: white;
		margin-top: 1rem;
	}
	.create-btn:hover {
		background-color: #1d4ed8;
	}
	.no-urls {
		text-align: center;
		padding: 3rem;
		color: #666;
	}
	.no-urls p {
		font-size: 1.1rem;
		margin-bottom: 1rem;
	}
	@media (max-width: 768px) {
		.urls-grid {
			grid-template-columns: 1fr;
		}
		.url-meta {
			flex-direction: column;
			gap: 0.25rem;
		}
		.url-actions {
			flex-direction: column;
		}
	}
</style>
