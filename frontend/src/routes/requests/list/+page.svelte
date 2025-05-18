<script lang="ts">
	import { onMount } from 'svelte';
	import { apiGet, apiPut } from '$lib/api';
	import BackToDashboard from '$lib/components/BackToDashboard.svelte';
	import Toast from '$lib/components/Toast.svelte';

	// your existing code...
	let allRequests = [];
	let requests = [];
	let filter = { worker: '', date: '', status: '' };
	let error = '';
	let success = '';
	let loading = true;

	async function loadData() {
		try {
		let query = '?';
		if (filter.status) query += `status=${filter.status}`;
		const res = await apiGet<{ data: any[] }>(`/requests${query}`);
		allRequests = res.data ?? [];
		applyFilter();
	} catch (err) {
		error = err.message;
	} finally {
		loading = false;
	}
	}

	function applyFilter() {
		requests = allRequests.filter((req) => {
			const matchesWorker = !filter.worker || req.worker.name.toLowerCase().includes(filter.worker.toLowerCase());
			const matchesDate = !filter.date || req.shift.date === filter.date;
			return matchesWorker && matchesDate;
		});
	}

	async function handleAction(id: number, action: 'approve' | 'reject') {
		try {
			await apiPut(`/requests/${id}/${action}`, {});
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
	<h1 class="text-4xl font-bold mb-10 text-center tracking-tight text-white-900">Pending Shift Requests</h1>

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
		<select
			class="select select-bordered w-48"
			bind:value={filter.status}
			on:change={loadData}
		>
			<option value="">All Statuses</option>
			<option value="pending">Pending</option>
			<option value="approved">Approved</option>
			<option value="rejected">Rejected</option>
		</select>
	</div>

	{#if loading}
		<div class="text-center text-gray-500">Loading...</div>
	{:else if requests.length === 0}
		<div class="text-center text-gray-400">No pending requests match your filter.</div>
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
        <th class="px-4 py-3">Status</th>
        <th class="px-4 py-3 text-center">Actions</th>
      </tr>
    </thead>
    <tbody>
      {#each requests as req}
        <tr class="hover:bg-gray-700">
          <td class="px-4 py-3 text-gray-300">{req.id}</td>
          <td class="px-4 py-3 text-gray-300">{req.worker.name}</td>
          <td class="px-4 py-3 text-gray-300">{req.shift.date}</td>
          <td class="px-4 py-3 text-gray-300">{req.shift.start_time} – {req.shift.end_time}</td>
          <td class="px-4 py-3 text-gray-300">{req.shift.role}</td>
          <td class="px-4 py-3 text-gray-300">{req.shift.location}</td>
          <td class="px-4 py-3">
            <span
				class="badge font-mono px-3 py-1 rounded text-sm"
				class:bg-yellow-500={req.status === 'pending'}
				class:text-gray-800={req.status === 'pending'}
				class:bg-green-500={req.status === 'approved'}
				class:text-white={req.status === 'approved'}
				class:bg-red-500={req.status === 'rejected'}
			>
				{req.status}
			</span>

          </td>
          <td class="px-4 py-3">
			{#if req.status === 'pending'}
				<div class="flex justify-center gap-2">
				<button
					class="btn btn-sm btn-outline border-green-500 text-green-400 hover:bg-green-600 hover:text-white"
					on:click={() => handleAction(req.id, 'approve')}
				>
					Approve
				</button>
				<button
					class="btn btn-sm btn-outline border-red-500 text-red-400 hover:bg-red-600 hover:text-white"
					on:click={() => handleAction(req.id, 'reject')}
				>
					Reject
				</button>
				</div>
			{/if}
			</td>
        </tr>
      {/each}
    </tbody>
  </table>
</div>
	{/if}
</div>
