<script lang="ts">
	import Home from '../../lib/components/home.svelte';
	import Tool from '../../lib/components/tool.svelte';
	import Globe from '../../lib/components/globe.svelte';
</script>

<div class="h-[50px] bg-neutral-800 w-full text-white">
	<div class="primary-container flex justify-between h-full">
		<ul class="flex gap-[16px]">
			<li class="border-white hover:border-b flex justify-center">
				<a href="/dashboard" class="flex py-[13px] gap-[10px]"><Home /> Home</a>
			</li>
			<li class="border-white hover:border-b flex justify-center">
				<a href="/dashboard/builder" class="flex py-[13px] gap-[10px]"><Tool /> Builder</a>
			</li>
			<li class="border-white hover:border-b flex justify-center">
				<a href="/dashboard/botnet" class="flex py-[13px] gap-[10px]"><Globe /> Botnet</a>
			</li>
		</ul>
	</div>
</div>

<slot />
