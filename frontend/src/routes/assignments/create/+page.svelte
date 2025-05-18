<script lang="ts">
	import { onMount } from 'svelte';
	import { apiGet, apiPost } from '$lib/api';
	import BackToDashboard from '$lib/components/BackToDashboard.svelte';
	import Toast from '$lib/components/Toast.svelte';

	let shift_id = '';
	let worker_id = '';
	let shifts = [];
	let workers = [];

	let success = '';
	let error = '';

	onMount(async () => {
		try {
			const shiftData = await apiGet('/shifts');
			const workerData = await apiGet('/workers');
			shifts = shiftData?.data;
			workers = workerData?.data;
		} catch (err) {
			error = '❌ Failed to fetch shifts or workers';
		}
	});

	async function submitAssignment() {
		try {
			await apiPost('/assignments', {
				shift_id: Number(shift_id),
				worker_id: Number(worker_id),
			});
			success = '✅ Assignment created successfully!';
			error = '';
			shift_id = '';
			worker_id = '';
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
	<h1 class="text-3xl font-bold text-center my-8 text-white">Create Assignment</h1>

	<form
		class="max-w-md mx-auto bg-gray-900 rounded-lg shadow divide-y divide-gray-700"
		on:submit|preventDefault={submitAssignment}
	>
		<div class="px-6 py-4 bg-gray-800 space-y-4">
			<div>
				<label class="block text-gray-300 font-semibold mb-1" for="shift">Shift:</label>
				<select
					id="shift"
					bind:value={shift_id}
					required
					class="w-full rounded-md px-3 py-2 bg-gray-700 text-white border border-gray-600 focus:outline-none focus:ring-2 focus:ring-indigo-500"
				>
					<option value="" disabled selected>Select a shift</option>
					{#each shifts as shift}
						<option value={shift.id}>
							{shift.date} | {shift.start_time} - {shift.end_time} | {shift.role}
						</option>
					{/each}
				</select>
			</div>

			<div>
				<label class="block text-gray-300 font-semibold mb-1" for="worker">Worker:</label>
				<select
					id="worker"
					bind:value={worker_id}
					required
					class="w-full rounded-md px-3 py-2 bg-gray-700 text-white border border-gray-600 focus:outline-none focus:ring-2 focus:ring-indigo-500"
				>
					<option value="" disabled selected>Select a worker</option>
					{#each workers as worker}
						<option value={worker.id}>{worker.name}</option>
					{/each}
				</select>
			</div>
		</div>

		<div class="px-6 py-4 bg-gray-800 text-center">
			<button
				type="submit"
				class="btn btn-primary px-6 py-2 rounded-md font-semibold disabled:opacity-50 disabled:cursor-not-allowed"
				disabled={!shift_id || !worker_id}
			>
				Assign Worker
			</button>
		</div>
	</form>
</div>
