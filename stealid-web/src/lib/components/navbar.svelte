<script lang="ts">
	import type { User } from '../../app';
	import UserIcon from './user.svelte';
	import { Popover, PopoverButton, PopoverPanel } from '@rgossiaux/svelte-headlessui';
	import { clsx } from 'clsx';

	export let user: User;
</script>

<header class="h-[60px] bg-neutral-800 w-full text-white border-b border-white">
	<div class="primary-container flex justify-between items-center h-full">
		<a href="/" class="text-2xl font-bold">Steal ID</a>
		<nav class="flex gap-2 items-center">
			{#if user}
				<p>{user.username}</p>
			{:else}
				<p>Not Logged In</p>
			{/if}
			<small>|</small>
			<Popover class="relative text-left">
				<PopoverButton class="flex items-center"><UserIcon /></PopoverButton>
				{#if user}
					<PopoverPanel
						class="absolute right-0 mt-2 w-56 origin-top-right rounded-md bg-neutral-800 border shadow-lg focus:outline-none py-2 divide-y"
					>
						<div class="py-1">
							<p class="font-bold px-2">{user.username}</p>
							<small class="px-2 text-neutral-300">ID: {user.id}</small>
						</div>
						<form class="hover:bg-neutral-700" action="/logout" method="POST">
							<button type="submit" class={clsx('px-2 py-1 block w-full text-left')}>Logout</button>
						</form>
					</PopoverPanel>
				{/if}
			</Popover>
		</nav>
	</div>
</header>
