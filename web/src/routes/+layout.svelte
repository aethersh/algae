<script>
	import ServerSelector from '$lib/components/ui/server-selector.svelte';
	import '../app.css';
</script>

<div class="max-w-2xl text-white">
	<div class="flex min-w-[42rem] items-center justify-between py-2">
		<h1 class="text-4xl">Algae</h1>
		<ServerSelector />
	</div>
	<slot></slot>
</div>
