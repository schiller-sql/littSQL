<script lang="ts">
  import { DataTable } from "carbon-components-svelte";
  import type { DataTableRow } from "carbon-components-svelte/types/DataTable/DataTable.svelte";

  import type { QueryExecResult } from "../sql-js/sql-wasm";

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
<DataTable
  size="medium"
  stickyHeader
  headers={result.columns.map((column) => {
    return {
      key: column === "id" ? replacementKeyForId : column,
      value: column,
    };
  })}
  rows={rows()}
/>
