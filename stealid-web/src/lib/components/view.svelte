<script lang="ts">
	import Eye from './eye.svelte';
	import {
		Dialog,
		DialogDescription,
		DialogOverlay,
		DialogTitle
	} from '@rgossiaux/svelte-headlessui';
	export let creds = '';
	let isOpen = false;
	const toggleIsOpen = () => {
		isOpen = true;
	};
</script>

<button
	class="px-2 py-1 bg-neutral-700 text-white rounded hover:ring ring-neutral-500 outline-none"
	on:click={toggleIsOpen}><Eye /></button
>
<Dialog open={isOpen} on:close={() => (isOpen = false)}>
	<DialogOverlay class="bg-black/60 inset-0 fixed w-screen h-screen" />

	<div
		class="shadow absolute z-40 py-2 top-32 left-1 right-1 bg-neutral-800 rounded primary-container"
	>
		<DialogTitle><h2 class="text-2xl font-bold mb-2">Loot</h2></DialogTitle>

		<DialogDescription>
			<textarea
				class="h-[300px] w-full bg-neutral-800 focus-visible:outline-none font-mono"
				readonly>{JSON.stringify(creds, null, ' ')}</textarea
			>
		</DialogDescription>
	</div>
</Dialog>
