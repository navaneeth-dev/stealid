<script lang="ts">
	export let buildID = '';
</script>

<form method="POST" action="/dashboard/builder?/download">
	<input type="hidden" name="buildID" value={buildID} />
	<button
		class="px-2 py-1 bg-neutral-700 text-white rounded hover:ring ring-neutral-500 outline-none inline-block"
	>
		<svg
			xmlns="http://www.w3.org/2000/svg"
			width="24"
			height="24"
			viewBox="0 0 24 24"
			fill="none"
			stroke="currentColor"
			stroke-width="2"
			stroke-linecap="round"
			stroke-linejoin="round"
			class="feather feather-download"
			><path d="M21 15v4a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2v-4" /><polyline
				points="7 10 12 15 17 10"
			/><line x1="12" y1="15" x2="12" y2="3" /></svg
		>
	</button>
</form>
