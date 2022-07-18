<script lang="ts">
  import { Modal } from "carbon-components-svelte";
  import type { QueryExecResult } from "../sql-js/sql-wasm";
  import { getAllTables } from "../util/db_util";
  import SqlResult from "./SqlResult.svelte";

  export let title: string;
  export let sql: string;
  export let open: boolean;

  let lastRenderedSql: string | undefined;
  let lastRenderedTables: Map<string, QueryExecResult> | undefined;

  $: if (open && lastRenderedSql !== sql) {
    lastRenderedTables = getAllTables(sql);
    lastRenderedSql = sql;
  }
</script>

<Modal modalHeading={title} hasScrollingContent passiveModal bind:open>
  {#if lastRenderedTables !== undefined}
    {#key lastRenderedTables}
      {#each [...lastRenderedTables] as [table, tableContent]}
        <div class="spacer" />
        <h4 style="padding: 16px; padding-bottom: 24;">
          {table}
        </h4>
        {#if tableContent !== undefined}
          <SqlResult result={tableContent} />
        {:else}
          <p style="color: grey">table is empty</p>
        {/if}
      {/each}
    {/key}
  {/if}
</Modal>
