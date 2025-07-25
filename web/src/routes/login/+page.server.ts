import { redirect, fail } from '@sveltejs/kit';

export const actions = {
  login: async ({ request, cookies }) => {
    const form = await request.formData();
    const email = form.get('email');
    const password = form.get('password');

    const res = await fetch('http://localhost:8080/auth/login', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ email, password })
    });

    if (!res.ok) {
      const err = await res.json();
      return fail(401, { error: err.message || 'Login failed' });
    }

    const { token } = await res.json();

    // Set secure, HTTP-only cookie
    cookies.set('auth_token', token, {
      httpOnly: true,
      path: '/',
      sameSite: 'lax',
      secure: process.env.NODE_ENV === 'production',
      maxAge: 60 * 60 * 24 * 7 // 1 week
    });

    throw redirect(302, '/dashboard');
  },
} satisfies Actions;
