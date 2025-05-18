<script lang="ts">
  let username = '';
  let pass = '';
  let error = '';
  let success = '';

  import Toast from '$lib/components/Toast.svelte';
  import { apiPost } from '$lib/api';
  import { goto } from '$app/navigation';

  async function handleLogin() {
    try {
      const res = await apiPost('/auth/login', {
        username: username,
        password: pass
      });

      // Store the token
       const token = res?.data?.access_token;
       if (token) {
           localStorage.setItem('access_token', token);
       }

      success = '✅ Login success !';
      error = '';
      goto('/');
    } catch (err) {
      try {
        const parsed = JSON.parse(err.message);
        error = `❌ ${parsed.error_message}`;
      } catch {
        error = `❌ ${err.message}`;
      }
    }
  }

  function clearError() {
    error = '';
  }

  function clearSuccess() {
    success = '';
  }
</script>

	{#if error}
		<Toast message={error} type="error" onClose={clearError} />
	{/if}
	
	{#if success}
		<Toast message={success} type="success" onClose={clearSuccess} />
	{/if}

<div class="min-h-screen bg-gradient-to-r from-purple-900 via-purple-900 to-pink-900 flex flex-col justify-center items-center px-6 py-12">
  <div class="bg-white rounded-lg shadow-lg p-8 max-w-md w-full">
    <h2 class="text-3xl font-extrabold text-gray-900 mb-6 text-center">Sign in to your account</h2>
    <form class="space-y-6" on:submit|preventDefault={handleLogin}>
      <div>
        <label for="username" class="block text-sm font-medium text-gray-700 mb-1">Username</label>
        <input
          id="username"
          type="text"
          bind:value={username}
          required
          class="w-full text-gray-700 border border-gray-300 rounded-md px-3 py-2 focus:outline-none focus:ring-2 focus:ring-indigo-500"
        />
      </div>
      <div>
        <label for="pass" class="block text-sm font-medium text-gray-700 mb-1">Password</label>
        <input
          id="pass"
          type="password"
          bind:value={pass}
          required
          class="w-full border text-gray-700 border-gray-300 rounded-md px-3 py-2 focus:outline-none focus:ring-2 focus:ring-indigo-500"
        />
      </div>
      <button
        type="submit"
        class="w-full bg-indigo-600 text-white py-2 rounded-md hover:bg-indigo-700 transition"
      >
        Sign in
      </button>
    </form>
  </div>
</div>
