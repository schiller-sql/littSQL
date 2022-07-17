<script lang="ts">
  import type { QueryExecResult } from "../sql-js/sql-wasm";
  import SqlResult from "./SqlResult.svelte";

  export let result: QueryExecResult[] | string;
</script>

{#if typeof result === "string"}
  <div class="spacer" />
  <div class="error">{result}</div>
{:else if result.length === 0}
  <div class="spacer" />
  <div>no result</div>
{:else}
  {#each result as resultSet ({})}
    <div class="spacer" />
    <SqlResult result={resultSet} />
  {/each}
{/if}
