import { redirect, fail } from '@sveltejs/kit';

export const actions = {
	register: async ({ request, cookies }) => {
    const form = await request.formData();
    const name = form.get('name');
    const email = form.get('email');
    const password = form.get('password');

    const res = await fetch('http://localhost:8080/auth/register', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ email, password, name })
    })
    if (!res.ok){
      const err = await res.json();
      return fail(401, { error: err.message || 'Login failed' });
    }

    throw redirect(302, '/login');
	}
} satisfies Actions;



