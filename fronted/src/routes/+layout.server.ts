import { API_BASE_URL } from '$env/static/private';
import type { User, AuthData } from '$lib/types.js'


export async function load({ cookies }): Promise<AuthData> {
	const authToken = cookies.get('JWT'); // Keep your original cookie name

	if (!authToken) {
		return {
			user: null,
			isAuthenticated: false
		};
	}

	try {
		const response = await fetch(`${API_BASE_URL}/auth/me`, {
			method: 'GET',
			headers: {
				'Content-Type': 'application/json',
				'Authorization': `Bearer ${authToken}`
			}
		});

		if (response.ok) {
			const rawData = await response.json();
			const userData: User = {
				id: rawData.id.toString(),
				user_name: rawData.user_name,
				user_email: rawData.user_email
			};
			return {
				user: userData,
				isAuthenticated: true
			};
		} else {
			console.log("Invalid token, deleting cookie");
			cookies.delete('auth_token', { path: '/' });
			return {
				user: null,
				isAuthenticated: false
			};
		}
	} catch (error) {
		console.error('Auth verification failed:', error);
		cookies.delete('auth_token', { path: '/' });
		return {
			user: null,
			isAuthenticated: false
		};
	}
}
