import { env } from '$env/dynamic/private';
import { redirect, fail } from '@sveltejs/kit';


const API_BASE_URL = env.API_BASE_URL || 'http://localhost:8080';

export const actions = {
	login: async ({ request, cookies }) => {
		const form = await request.formData();
		const email = form.get('email');
		const password = form.get('password');

		const res = await fetch(`${API_BASE_URL}/auth/login`, {
			method: 'POST',
			headers: { 'Content-Type': 'application/json' },
			body: JSON.stringify({ email, password })
		});

		if (!res.ok) {
			const err = await res.json();
			return fail(401, { status: err.status, error: true });
		}

		const { auth_token } = await res.json();

		cookies.set('JWT', auth_token, {
			httpOnly: true,
			path: '/',
			sameSite: 'lax',
			secure: process.env.NODE_ENV === 'production',
			maxAge: 60 * 60 * 24 * 7 // 1 week
		});

		throw redirect(302, '/dashboard');
	},
};
