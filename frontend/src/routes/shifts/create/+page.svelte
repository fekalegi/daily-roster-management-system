<script lang="ts">
	import { apiPost } from '$lib/api';
	import BackToDashboard from '$lib/components/BackToDashboard.svelte';
	import Toast from '$lib/components/Toast.svelte';

	let date = '';
	let start_time = '';
	let end_time = '';
	let role = '';
	let location = '';
	let success = '';
	let error = '';

	async function submitShift() {
		try {
			await apiPost('/shifts', {
				date,
				start_time,
				end_time,
				role,
				location
			});
			success = '✅ Shift created successfully!';
			error = '';
			// Clear form
			date = '';
			start_time = '';
			end_time = '';
			role = '';
			location = '';
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
	<h1 class="text-3xl font-bold text-center my-8 text-white">Create a New Shift</h1>

	<form
		class="max-w-md mx-auto bg-gray-900 rounded-lg shadow divide-y divide-gray-700"
		on:submit|preventDefault={submitShift}
	>
		<div class="px-6 py-4 bg-gray-800 space-y-4">
			<div>
				<label class="block text-gray-300 font-semibold mb-1" for="date">Date:</label>
				<input
					id="date"
					type="date"
					bind:value={date}
					class="w-full rounded-md px-3 py-2 bg-gray-700 text-white border border-gray-600 focus:outline-none focus:ring-2 focus:ring-indigo-500"
					required
				/>
			</div>

			<div>
				<label class="block text-gray-300 font-semibold mb-1" for="start">Start Time:</label>
				<input
					id="start"
					type="time"
					bind:value={start_time}
					class="w-full rounded-md px-3 py-2 bg-gray-700 text-white border border-gray-600 focus:outline-none focus:ring-2 focus:ring-indigo-500"
					required
				/>
			</div>

			<div>
				<label class="block text-gray-300 font-semibold mb-1" for="end">End Time:</label>
				<input
					id="end"
					type="time"
					bind:value={end_time}
					class="w-full rounded-md px-3 py-2 bg-gray-700 text-white border border-gray-600 focus:outline-none focus:ring-2 focus:ring-indigo-500"
					required
				/>
			</div>
		</div>

		<div class="px-6 py-4 bg-gray-700 space-y-4">
			<div>
				<label class="block text-gray-300 font-semibold mb-1" for="role">Role:</label>
				<input
					id="role"
					type="text"
					bind:value={role}
					class="w-full rounded-md px-3 py-2 bg-gray-700 text-white border border-gray-600 focus:outline-none focus:ring-2 focus:ring-indigo-500"
					required
				/>
			</div>

			<div>
				<label class="block text-gray-300 font-semibold mb-1" for="location">Location:</label>
				<input
					id="location"
					type="text"
					bind:value={location}
					class="w-full rounded-md px-3 py-2 bg-gray-700 text-white border border-gray-600 focus:outline-none focus:ring-2 focus:ring-indigo-500"
					required
				/>
			</div>
		</div>

		<div class="px-6 py-4 bg-gray-800 text-center">
			<button
				type="submit"
				class="btn btn-primary px-6 py-2 rounded-md font-semibold disabled:opacity-50 disabled:cursor-not-allowed"
				disabled={!date || !start_time || !end_time || !role || !location}
			>
				Create Shift
			</button>
		</div>
	</form>
</div>
