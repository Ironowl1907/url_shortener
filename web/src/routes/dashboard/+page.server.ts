import type { PageServerLoad, Actions } from './$types';
import type { ShortenedUrl } from '$lib/types.js';
import { fail } from '@sveltejs/kit';
import { env } from '$env/dynamic/private';
import { makeAuthenticatedRequest } from '$lib/utils.js'

// Use environment variable for API base URL
const API_BASE_URL = env.API_BASE_URL || 'http://localhost:8080';

interface LoadResult {
	shortenedUrls: ShortenedUrl[] | null;
	apiBaseUrl: string;
}

interface ApiUrlResponse {
	id: number;
	original_url: string;
	short_code: string;
	title: string;
	description: string;
	owner_id: number;
	created_at: string;
	updated_at: string;
}

// Transform API response to ShortenedUrl type
function transformApiResponse(item: ApiUrlResponse): ShortenedUrl {
	return {
		id: item.id,
		OriginalURL: item.original_url,
		ShortCode: item.short_code,
		Title: item.title,
		Description: item.description,
		OwnerID: item.owner_id,
		CreatedAt: new Date(item.created_at),
		UpdatedAt: new Date(item.updated_at),
	};
}

export const load: PageServerLoad = async ({ cookies }): Promise<LoadResult> => {
	const authToken = cookies.get('JWT');

	if (!authToken) {
		return {
			shortenedUrls: null,
			apiBaseUrl: API_BASE_URL
		};
	}

	try {
		const response = await makeAuthenticatedRequest(
			`${API_BASE_URL}/urls`,
			authToken,
			{ method: 'GET' }
		);

		if (!response.ok) {
			console.error('Failed to fetch URLs:', response.status, response.statusText);
			return {
				shortenedUrls: null,
				apiBaseUrl: API_BASE_URL
			};
		}

		const rawData: ApiUrlResponse[] = await response.json();
		const shortenedUrls = rawData.map(transformApiResponse);

		return {
			shortenedUrls,
			apiBaseUrl: API_BASE_URL
		};
	} catch (error) {
		console.error('Error fetching URLs:', error);
		return {
			shortenedUrls: null,
			apiBaseUrl: API_BASE_URL
		};
	}
};

export const actions: Actions = {
	delete: async ({ request, cookies }) => {
		const authToken = cookies.get('JWT');

		if (!authToken) {
			return fail(401, { error: 'Not authenticated' });
		}

		const data = await request.formData();
		const id = data.get('id') as string;

		if (!id) {
			return fail(400, { error: 'ID is required' });
		}

		try {
			const response = await makeAuthenticatedRequest(
				`${API_BASE_URL}/urls/${id}`,
				authToken,
				{ method: 'DELETE' }
			);

			if (!response.ok) {
				console.error('Failed to delete URL:', response.status, response.statusText);
				return fail(response.status, { error: 'Failed to delete URL' });
			}

			return { success: true };
		} catch (error) {
			console.error('Error deleting URL:', error);
			return fail(500, { error: 'Server error occurred' });
		}
	}
};
