<script lang="ts">
	export let country_code: string;
</script>

<img src={'/' + country_code + '.svg'} alt={country_code} class="w-8 h-8 inline" />
<span class="uppercase">
	{country_code}
</span>
