<script lang="ts">
	import Box from '$lib/components/box.svelte';
</script>

<div class="primary-container">
	<h2 class="text-3xl font-bold py-6">Builder</h2>
	<form method="POST">
		<button
			type="submit"
			class="bg-neutral-800 text-white rounded p-2 font-bold hover:ring ring-neutral-500 flex gap-2"
		>
			<Box />
			<span>Create Build</span>
		</button>
	</form>
</div>
