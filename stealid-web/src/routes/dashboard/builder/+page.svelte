<script lang="ts">
	import Box from '$lib/components/box.svelte';
	import type { PageData } from './$types';
	import { format } from 'date-fns';
	import {
		createSvelteTable,
		getCoreRowModel,
		flexRender,
		type ColumnDef,
		type TableOptions
	} from '@tanstack/svelte-table';
	import type { Build } from './+page.server';
	import Country from '$lib/components/country.svelte';
	import View from '$lib/components/view.svelte';

	const defaultColumns: ColumnDef<Build>[] = [
		{
			accessorKey: 'id',
			header: 'ID',
			cell: (info) => info.getValue()
		}
	];

	export let data: PageData;
	const options: TableOptions<Bot> = {
		data: data.builds,
		columns: defaultColumns,
		getCoreRowModel: getCoreRowModel()
	};

	const table = createSvelteTable(options);
</script>

<div class="primary-container">
	<h2 class="text-3xl font-bold py-6">Builder</h2>

	<section>
		<div class="flex justify-between">
			<div class="flex items-center">
				<input
					type="text"
					id="table-search"
					class="bg-neutral-800 text-white focus-visible:outline-none focus:ring ring-neutral-500 rounded px-2 py-1"
					placeholder="Search for items"
				/>
			</div>
			<button
				type="submit"
				class="bg-neutral-800 text-white rounded py-1 px-2 font-bold hover:ring ring-neutral-500 flex gap-2"
			>
				<Box />
				<span>Create Build</span>
			</button>
		</div>

		<table class="w-full bg-neutral-800 rounded mt-3 font-mono">
			<thead>
				{#each $table.getHeaderGroups() as headerGroup}
					<tr>
						{#each headerGroup.headers as header}
							<th class="uppercase font-normal text-sm text-left p-2">
								{#if !header.isPlaceholder}
									<svelte:component
										this={flexRender(header.column.columnDef.header, header.getContext())}
									/>
								{/if}
							</th>
						{/each}
					</tr>
				{/each}
			</thead>
			<tbody>
				{#each $table.getRowModel().rows as row}
					<tr>
						{#each row.getVisibleCells() as cell}
							<td class="p-2 text-base">
								<svelte:component
									this={flexRender(cell.column.columnDef.cell, cell.getContext())}
								/>
							</td>
						{/each}
					</tr>
				{/each}
			</tbody>
		</table>
	</section>
</div>
