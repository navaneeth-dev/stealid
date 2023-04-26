<script lang="ts">
	import { superForm } from 'sveltekit-superforms/client';
	import type { PageData } from './$types';

	export let data: PageData;
	const { form, errors, constraints, message } = superForm(data.form);
</script>

<section class="primary-container">
	<div class="flex flex-col items-center justify-center h-[calc(100vh-60px)]">
		<div class="w-1/2 text-center">
			<h2 class="text-3xl font-bold">Login</h2>
			<form method="POST" class="mt-8 flex flex-col gap-6">
				<div class="flex flex-col gap-1 text-left">
					<label for="username">Username:</label>
					<input
						type="text"
						id="username"
						class="border border-neutral-400 focus:border-white bg-transparent rounded p-2 outline-none"
						autocomplete="off"
						placeholder="abc"
						name="username"
						bind:value={$form.username}
						{...$constraints.username}
					/>
					{#if $errors.username}
						<small class="text-red-500">{$errors.username}</small>
					{/if}
				</div>
				<div class="flex flex-col gap-1 text-left">
					<label for="password">Password:</label>
					<input
						type="password"
						id="password"
						class="border border-neutral-400 focus:border-white bg-transparent rounded p-2 outline-none"
						autocomplete="off"
						placeholder="***"
						name="password"
						bind:value={$form.password}
						{...$constraints.password}
					/>
					{#if $errors.password}
						<small class="text-red-500">{$errors.password}</small>
					{/if}
				</div>
				<!-- <div class="flex flex-col gap-1 text-left">
					<label for="code">Invitation Code:</label>
					<input
						type="text"
						id="code"
						class="border border-white bg-transparent rounded p-2 outline-none"
						autocomplete="off"
						placeholder="1234-5678-1234"
					/>
				</div> -->
				<button
					type="submit"
					class="bg-neutral-800 text-white rounded p-2 hover:bg-neutral-600 font-bold">Login</button
				>
			</form>
			{#if $message}
				<p class="text-red-500 mt-2">{$message}</p>
			{/if}
		</div>
	</div>
</section>
