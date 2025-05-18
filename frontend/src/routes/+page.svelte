<script lang="ts">
  import { onMount } from 'svelte';
  import { apiGetAuth } from '$lib/api';
  import Toast from '$lib/components/Toast.svelte';
  import { goto } from '$app/navigation';

  let error = '';
  let user: any = null;
  let isAdmin = false;

  onMount(async () => {
    try {
      await apiGetAuth('/auth/verify');
      const role = localStorage.getItem('role_id');
     isAdmin = role === '1';
    } catch (err) {
		console.log(err);
		error = '❌ Unauthorized';
    localStorage.clear();       
		goto('/login');
	}
  });

  function clearError() {
    error = '';
  }

  function logout() {
    localStorage.clear();       
    goto('/login');             
  }
</script>

{#if error}
  <Toast message={error} type="error" onClose={clearError} />
{/if}

<h1 class="text-4xl font-bold text-center my-10 text-white">Welcome to Daily Worker Roster UI</h1>

<nav class="max-w-md mx-auto rounded-lg overflow-hidden shadow-lg">
  <ul>
    <li class="bg-gray-800">
      <a href="/shifts" class="block px-6 py-4 text-white font-semibold hover:bg-indigo-600 transition-colors">View Shifts</a>
    </li>
    <li class="bg-gray-700">
      <a href="/requests" class="block px-6 py-4 text-white font-semibold hover:bg-indigo-600 transition-colors">Request a Shift</a>
    </li>
    <li class="bg-gray-800">
      <a href="/requests/pending" class="block px-6 py-4 text-white font-semibold hover:bg-indigo-600 transition-colors">View Pending Requests</a>
    </li>
    <li class="bg-gray-700">
      <a href="/assignments" class="block px-6 py-4 text-white font-semibold hover:bg-indigo-600 transition-colors">View Assigned Shifts</a>
    </li>
    {#if isAdmin}
  <li class="bg-gray-800">
    <a href="/shifts/create" class="block px-6 py-4 text-white font-semibold hover:bg-indigo-600 transition-colors">
      Create Shifts
    </a>
  </li>
{/if}
  </ul>
</nav>

<div class="text-center mt-6">
  <button
    class="btn btn-sm btn-outline btn-error"
    on:click={logout}
  >
    Logout
  </button>
</div>

