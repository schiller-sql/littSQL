<script lang="ts">
  import { DataTable } from "carbon-components-svelte";
  import type { DataTableRow } from "carbon-components-svelte/types/DataTable/DataTable.svelte";

  import type { QueryExecResult } from "../sql-js/sql-wasm";
  import SqlCell from "./SqlCell.svelte";

  export let title: string = "";
  export let result: QueryExecResult;

  const replacementKeyForId = "%"; // "id" cannot be used as a key in the DataTable

  function rows(): DataTableRow[] {
    return (result as QueryExecResult).values.map((rawRow: any, i: number) => {
      const row = { id: i };
      for (let j = 0; j < (result as QueryExecResult).columns.length; j++) {
        let columnName = (result as QueryExecResult).columns[j];
        if (columnName === "id") {
          columnName = replacementKeyForId;
        }
        row[columnName] = rawRow[j];
      }
      return row;
    });
  }
</script>

<!-- zebra -->
<div style="overflow-x: scroll; max-height:250px">
  <DataTable
    useStaticWidth
    title={title || undefined}
    size="short"
    headers={result.columns.map((column) => {
      return {
        key: column === "id" ? replacementKeyForId : column,
        value: column,
      };
    })}
    rows={rows()}
  >
    <!-- TODO: performant mode, without this behaviour -->
    <svelte:fragment slot="cell" let:cell>
      <span
        class:number={typeof cell.value === "number"}
        class:null={typeof cell.value === "object"}
        class:string={typeof cell.value === "string"}
      >
        {#if typeof cell.value === "string"}
          "{cell.value}"
        {:else}
          {cell.value}
        {/if}
      </span>
    </svelte:fragment>
  </DataTable>
</div>

<style>
  span.number {
    color: #a7f0ba;
  }

  span.null {
    color: #bb8eff;
  }

  span.string {
    color: #fa75a6;
  }
</style>
