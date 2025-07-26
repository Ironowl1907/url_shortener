export async function load({ cookies }) {
  const authToken = cookies.get('JWT'); // Keep your original cookie name

  if (!authToken) {
    return {
      user: null,
      isAuthenticated: false
    };
  }

  try {
    const response = await fetch('http://localhost:8080/auth/me', {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json',
        'Authorization': `Bearer ${authToken}`
      }
    });

    if (response.ok) {
      const userData = await response.json();
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
