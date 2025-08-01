
// Helper function for authenticated API requests
export async function makeAuthenticatedRequest(
  url: string,
  authToken: string,
  options: RequestInit = {}
): Promise<Response> {
  return fetch(url, {
    ...options,
    headers: {
      'Content-Type': 'application/json',
      'Authorization': `Bearer ${authToken}`,
      ...options.headers,
    }
  });
}
