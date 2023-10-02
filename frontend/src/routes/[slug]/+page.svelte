<script>
	import { onMount } from 'svelte';

	export let data;

	onMount(() => {
		if (!data.response.true_url.startsWith('https://')) {
			window.location.href = `https://${data.response.true_url}`;
			return;
		}

		window.location.href = `${data.response.true_url}`;
	});
</script>
