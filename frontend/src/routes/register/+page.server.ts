import { redirect, fail } from '@sveltejs/kit';
import type { Actions } from './$types';
import { API_BASE_URL } from '$env/static/private';


export const actions = {
	register: async ({ request, cookies }) => {
		const form = await request.formData();
		const name = form.get('name');
		const email = form.get('email');
		const password = form.get('password');

		console.log("API URL: " + API_BASE_URL)
		const res = await fetch(`${API_BASE_URL}/auth/register`, {
			method: 'POST',
			headers: { 'Content-Type': 'application/json' },
			body: JSON.stringify({ email, password, name })
		});

		if (!res.ok) {
			const err = await res.json();
			console.log(err);
			return fail(400, { error: true, message: err.status });
		}

		throw redirect(302, '/login');
	}
} satisfies Actions;
