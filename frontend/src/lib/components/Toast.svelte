<script lang="ts">
  export let message: string;
  export let type: 'success' | 'error' | 'info' = 'info';
  export let onClose: () => void;

  let bgClass = 'bg-gray-800';
  let textClass = 'text-white';

  $: if (type === 'error') {
    bgClass = 'bg-red-600';
  } else if (type === 'success') {
    bgClass = 'bg-green-600';
  } else if (type === 'info') {
    bgClass = 'bg-blue-600';
  }
</script>

<div
  class={`fixed top-5 right-5 z-50 px-4 py-3 rounded-lg shadow-lg flex items-center gap-2 animate-fade-in ${bgClass} ${textClass}`}
  role="alert"
>
  <span class="font-semibold capitalize">{type}:</span>
  <span class="break-all">{message}</span>
  <button
    class="ml-4 hover:text-gray-200 font-bold"
    on:click={onClose}
    aria-label="Close"
  >
    &times;
  </button>
</div>

<style>
  @keyframes fade-in {
    from {
      opacity: 0;
      transform: translateY(-10px);
    }
    to {
      opacity: 1;
      transform: translateY(0);
    }
  }
  .animate-fade-in {
    animation: fade-in 0.3s ease-out;
  }
</style>
