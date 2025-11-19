import { fail, json } from '@sveltejs/kit';
import type { PageServerLoad, Actions } from './$types';
import { makeAuthenticatedRequest } from '$lib/utils';
import { API_BASE_URL } from '$env/static/private';


export const actions = {
	newUrl: async ({ cookies, request }) => {
		const data = await request.formData();
		const url = data.get('url');
		const description = data.get('description');
		const title = data.get('title');
		const ignore_response: boolean = data.get('ignore_response')?.toString() === "on" ? true : false;

		// Get auth token
		const authToken = cookies.get('JWT');
		if (!authToken) {
			return fail(401, { error: 'Not authenticated' });
		}

		// Try post petition
		try {
			const response = await makeAuthenticatedRequest(`${API_BASE_URL}/urls`, authToken, {
				method: "POST", body: JSON.stringify({ url, description, title, ignore_response })
			})
			if (response.ok) {
				return { success: true };
			}

			var res = await response.json()
			console.error('Failed to craete URL:', response.status, res);
			if (response.status == 400) {
				return fail(response.status, { noReachable: true });
			}
			return fail(response.status, { error: 'Failed to create URL' });

		} catch (error) {
			console.error('Error Creating URL:', error);
			return fail(500, { error: 'Server error occurred' });
		}
	},
} satisfies Actions;
