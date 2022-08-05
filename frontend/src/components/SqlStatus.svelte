<script lang="ts">
  import type { SqlStatus } from "../stores/sql_status";
  import { InlineLoading } from "carbon-components-svelte";
  import { databaseStore } from "../stores/global_stores";
  import { DatabaseState } from "../stores/database";

  export let sqlStatus: SqlStatus;
</script>

{#if $databaseStore === DatabaseState.unready}
  <InlineLoading status="inactive" description="loading sql engine..." />
{:else if sqlStatus.status === "loading"}
  <InlineLoading status="active" description="checking" />
{:else if sqlStatus.status === "error"}
  <InlineLoading status="error" description={sqlStatus.error} />
{:else}
  <InlineLoading status="finished" description="no errors" />
{/if}
