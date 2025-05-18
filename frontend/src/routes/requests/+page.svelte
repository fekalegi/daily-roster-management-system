<script lang="ts">
	import { onMount } from 'svelte';
	import { apiGet, apiPost } from '$lib/api';
	import BackToDashboard from '$lib/components/BackToDashboard.svelte';
	import Toast from '$lib/components/Toast.svelte';

	let shifts = [];
	let selectedShiftId: number | null = null;
	let workerId = 0;
	let success = '';
	let error = '';

	onMount(async () => {
		const res = await apiGet<{ data: any[] }>('/shifts');
		shifts = res.data;
		workerId = localStorage.getItem('worker_id');
	});

	async function submitRequest() {
		try {
			await apiPost('/requests', {
				worker_id: workerId,
				shift_id: selectedShiftId
			});
			success = '✅ Request submitted!';
			error = '';
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

<div class="max-w-7xl mx-auto px-6 py-12">
	<BackToDashboard />
<h1 class="text-3xl font-bold text-center my-8 text-white">Request a Shift</h1>

<form
	class="max-w-md mx-auto bg-gray-900 rounded-lg shadow divide-y divide-gray-700"
	on:submit|preventDefault={submitRequest}
>
	<div class="px-6 py-4 bg-gray-800">
		<label class="block text-gray-300 font-semibold mb-1" for="workerId">Your Worker ID:</label>
		<input
			id="workerId"
			type="number"
			bind:value={workerId}
			min="1"
			disabled
			class="w-full rounded-md px-3 py-2 bg-gray-700 text-white border border-gray-600 focus:outline-none focus:ring-2 focus:ring-indigo-500"
		/>
	</div>

	<div class="px-6 py-4 bg-gray-700">
		<label class="block text-gray-300 font-semibold mb-1" for="shiftSelect">Select Shift:</label>
		<select
			id="shiftSelect"
			bind:value={selectedShiftId}
			class="w-full rounded-md px-3 py-2 bg-gray-700 text-white border border-gray-600 focus:outline-none focus:ring-2 focus:ring-indigo-500"
		>
			<option value={null} disabled selected>-- Choose a shift --</option>
			{#each shifts as shift, i}
				<option
					value={shift.id}
					class={i % 2 === 0 ? 'bg-gray-800' : 'bg-gray-700'}
				>
					{shift.date} | {shift.start_time}–{shift.end_time} | {shift.role}
				</option>
			{/each}
		</select>
	</div>

	<div class="px-6 py-4 bg-gray-800 text-center">
		<button
			type="submit"
			class="btn btn-primary px-6 py-2 rounded-md font-semibold disabled:opacity-50 disabled:cursor-not-allowed"
			disabled={!selectedShiftId}
		>
			Submit Request
		</button>
	</div>
</form>
</div>
