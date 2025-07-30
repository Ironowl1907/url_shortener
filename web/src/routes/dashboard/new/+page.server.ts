import { fail } from '@sveltejs/kit';
import type { PageServerLoad, Actions } from './$types';

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
      const response = await fetch(`http://localhost:8080/urls`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
          'Authorization': `Bearer ${authToken}`,
          body: JSON.stringify({ url, description, title, ignore_response })
        }
      });
      if (response.ok) {
        return { success: true };
      }

      var res = await response.json()
      console.error('Failed to craete URL:', response.status, res);
      return fail(response.status, { error: 'Failed to create URL' });

    } catch (error) {
      console.error('Error Creating URL:', error);
      return fail(500, { error: 'Server error occurred' });
    }
  },
} satisfies Actions;
