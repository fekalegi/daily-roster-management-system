<script lang="ts">
	import { onMount } from 'svelte';
	import { apiGet } from '$lib/api';
	import BackToDashboard from '$lib/components/BackToDashboard.svelte';

	let shifts = [];
	let error = '';
	let loading = true;

	onMount(async () => {
		try {
			const res = await apiGet<{ data: any[] }>('/shifts/unassigned');
			shifts = res.data ?? [];
		} catch (err) {
			error = err.message;
		} finally {
			loading = false;
		}
	});
</script>

<div class="max-w-7xl mx-auto px-6 py-12">
	<BackToDashboard />
	<h1 class="text-3xl font-bold text-center text-white mb-8">Available (Unassigned) Shifts</h1>

	{#if loading}
		<p class="text-center text-gray-400">Loading shifts...</p>
	{:else if error}
		<p class="text-center text-red-500">{error}</p>
	{:else if shifts.length === 0}
		<p class="text-center text-gray-400">No available shifts at the moment.</p>
	{:else}
		<div class="overflow-x-auto shadow rounded-lg border border-gray-700 bg-gray-900">
			<table class="table w-full text-sm table-zebra">
				<thead class="bg-gray-800 text-white uppercase">
					<tr>
						<th class="px-4 py-3">Date</th>
						<th class="px-4 py-3">Time</th>
						<th class="px-4 py-3">Role</th>
						<th class="px-4 py-3">Location</th>
					</tr>
				</thead>
				<tbody>
					{#each shifts as shift}
						<tr class="hover:bg-gray-700">
							<td class="px-4 py-3 text-gray-300">{shift.date}</td>
							<td class="px-4 py-3 text-gray-300">{shift.start_time} – {shift.end_time}</td>
							<td class="px-4 py-3 text-gray-300">{shift.role}</td>
							<td class="px-4 py-3 text-gray-300">{shift.location}</td>
						</tr>
					{/each}
				</tbody>
			</table>
		</div>
	{/if}
</div>
