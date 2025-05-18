<script lang="ts">
	import { goto } from '$app/navigation';

	export let label = '← Back to Dashboard';
	export let href = '/';

	function handleClick(event: MouseEvent) {
		event.preventDefault();
		goto(href);
	}
</script>

<button
	type="button"
	class="btn btn-sm mb-4"
	on:click={handleClick}
	aria-label="Back to Dashboard"
>
	{label}
</button>
