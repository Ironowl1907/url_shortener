// @ts-nocheck
import type { PageServerLoad, Actions } from './$types';
import type { ShortenedUrl } from '$lib/types.js';
import { fail, redirect } from '@sveltejs/kit';

interface LoadResult {
  shortenedUrls: ShortenedUrl[] | null;
}

export const load = async ({ cookies }: Parameters<PageServerLoad>[0]): Promise<LoadResult> => {
  const authToken = cookies.get('JWT');
  if (!authToken) {
    return {
      shortenedUrls: null
    };
  }

  try {
    const response = await fetch('http://localhost:8080/urls', {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json',
        'Authorization': `Bearer ${authToken}`
      }
    });

    if (response.ok) {
      const rawData = await response.json();
      const shortenedUrls: ShortenedUrl[] = rawData.map((item: any) => ({
        id: item.id,
        OriginalURL: item.original_url,
        ShortCode: item.short_code,
        Title: item.title,
        Description: item.description,
        OwnerID: item.owner_id,
        CreatedAt: new Date(item.created_at),
        UpdatedAt: new Date(item.updated_at),
      }));

      return {
        shortenedUrls
      };
    } else {
      console.error('Failed to fetch URLs:', response.status, response.statusText);
      return {
        shortenedUrls: null
      };
    }
  } catch (error) {
    console.error('Error fetching URLs:', error);
    return {
      shortenedUrls: null
    };
  }
};

export const actions = {
  delete: async ({ request, cookies }: import('./$types').RequestEvent) => {
    const authToken = cookies.get('JWT');
    if (!authToken) {
      return fail(401, { error: 'Not authenticated' });
    }

    const data = await request.formData();
    const id = data.get('id') as string;
    console.log(id)

    if (!id) {
      return fail(400, { error: 'Short code is required' });
    }

    try {
      const response = await fetch(`http://localhost:8080/urls/${id}`, {
        method: 'DELETE',
        headers: {
          'Content-Type': 'application/json',
          'Authorization': `Bearer ${authToken}`
        }
      });

      if (response.ok) {
        return { success: true };
      } else {
        console.error('Failed to delete URL:', response.status, response.statusText);
        return fail(response.status, { error: 'Failed to delete URL' });
      }
    } catch (error) {
      console.error('Error deleting URL:', error);
      return fail(500, { error: 'Server error occurred' });
    }
  }
};
;null as any as Actions;