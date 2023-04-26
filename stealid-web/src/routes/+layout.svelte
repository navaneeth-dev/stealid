<script lang="ts">
	import '../app.css';
	import Navbar from '../lib/components/navbar.svelte';
	import type { LayoutData } from './$types';
	export let data: LayoutData;
	const { user } = data;
</script>

<Navbar {user} />

<slot />
