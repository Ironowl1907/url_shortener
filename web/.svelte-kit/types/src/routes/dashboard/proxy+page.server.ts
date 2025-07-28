// @ts-nocheck
import type { PageServerLoad } from './$types';
import type { ShortenedUrl } from '$lib/types.js';

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
        ...item,
        OriginalURL: item.original_url,
        CreatedAt: new Date(item.CreatedAt),
        UpdatedAt: new Date(item.UpdatedAt),
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
