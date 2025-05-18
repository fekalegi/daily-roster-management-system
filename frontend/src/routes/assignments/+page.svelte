<script lang="ts">
	import { onMount } from 'svelte';
	import { apiGet, apiPut } from '$lib/api';
	import BackToDashboard from '$lib/components/BackToDashboard.svelte';
	import Toast from '$lib/components/Toast.svelte';

	// your existing code...
	let allAssignments = [];
	let assignments = [];
	let filter = { worker: '', date: '' };
	let error = '';
	let success = '';
	let loading = true;

	async function loadData() {
		try {
			const res = await apiGet<{ data: any[] }>('/assignments');
			allAssignments = res.data ?? [];
			applyFilter();
		} catch (err) {
			error = err.message;
		} finally {
			loading = false;
		}
	}

	function applyFilter() {
		assignments = allAssignments.filter((req) => {
			const matchesWorker = !filter.worker || req.worker.name.toLowerCase().includes(filter.worker.toLowerCase());
			const matchesDate = !filter.date || req.shift.date === filter.date;
			return matchesWorker && matchesDate;
		});
	}

	async function handleAction(id: number, action: 'approve' | 'reject') {
		try {
			await apiPut(`/assignments/${id}/${action}`, {});
			await loadData();
		} catch (err) {
			error = err.message;
		} finally {
			success = 'Success';
		}
	}
	onMount(loadData);

	function clearError() {
		error = '';
	}

	function clearSuccess() {
		success = '';
	}
</script>
<div class="max-w-7xl mx-auto px-6 py-12">
	{#if error}
		<Toast message={error} type="error" onClose={clearError} />
	{/if}
	
	{#if success}
		<Toast message={success} type="success" onClose={clearSuccess} />
	{/if}
	<BackToDashboard />
	<h1 class="text-4xl font-bold mb-10 text-center tracking-tight text-white-900">Shift Assigned</h1>

	<div class="flex flex-wrap justify-center gap-4 mb-6">
		<input
			class="input input-bordered w-60"
			type="text"
			placeholder="Search by worker"
			bind:value={filter.worker}
			on:input={applyFilter}
		/>
		<input
			class="input input-bordered w-48"
			type="date"
			bind:value={filter.date}
			on:input={applyFilter}
		/>
	</div>

	{#if loading}
		<div class="text-center text-gray-500">Loading...</div>
	{:else if assignments.length === 0}
		<div class="text-center text-gray-400">No assigned shift match your filter.</div>
	{:else}
		<div class="overflow-x-auto rounded-xl shadow border border-gray-700 bg-gray-900">
  <table class="table table-zebra w-full">
    <thead class="bg-gray-800 text-gray-100 text-sm font-semibold">
      <tr>
        <th class="px-4 py-3">ID</th>
        <th class="px-4 py-3">Worker</th>
        <th class="px-4 py-3">Shift Date</th>
        <th class="px-4 py-3">Time</th>
        <th class="px-4 py-3">Role</th>
        <th class="px-4 py-3">Location</th>
      </tr>
    </thead>
    <tbody>
      {#each assignments as req}
        <tr class="hover:bg-gray-700">
          <td class="px-4 py-3 text-gray-300">{req.id}</td>
          <td class="px-4 py-3 text-gray-300">{req.worker.name}</td>
          <td class="px-4 py-3 text-gray-300">{req.shift.date}</td>
          <td class="px-4 py-3 text-gray-300">{req.shift.start_time} – {req.shift.end_time}</td>
          <td class="px-4 py-3 text-gray-300">{req.shift.role}</td>
          <td class="px-4 py-3 text-gray-300">{req.shift.location}</td>
        </tr>
      {/each}
    </tbody>
  </table>
</div>
	{/if}
</div>
