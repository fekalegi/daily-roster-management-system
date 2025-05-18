<script lang="ts">
	import { onMount } from 'svelte';
	import { apiGet } from '$lib/api';
	import BackToDashboard from '$lib/components/BackToDashboard.svelte';

	let shifts = [];

	onMount(async () => {
		const res = await apiGet<{ data: any[] }>('/shifts');
		shifts = res.data;
	});
</script>
<div class="max-w-7xl mx-auto px-6 py-12">
	<BackToDashboard />

<h1 class="text-3xl font-bold text-center my-8 text-white">Shifts</h1>

{#if shifts.length === 0}
	<p class="text-center text-gray-400">No shifts available.</p>
{:else}
	<ul class="max-w-md mx-auto bg-gray-900 rounded-lg shadow divide-y divide-gray-700">
		{#each shifts as shift}
			<li class="px-6 py-4 text-white hover:bg-gray-800 cursor-pointer transition-colors">
				<span class="font-semibold">{shift.date}</span> — {shift.role} at <span class="italic">{shift.location}</span>
			</li>
		{/each}
	</ul>
{/if}
</div>
