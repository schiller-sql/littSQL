<script lang="ts">
  import type { SqlStatus } from "../util/db_util";
  import { databaseIsReadyStore } from "../database";
  import { InlineLoading } from "carbon-components-svelte";

  export let sqlStatus: SqlStatus;
</script>

{#if !$databaseIsReadyStore}
  <InlineLoading status="inactive" description="loading sql engine..." />
{:else if sqlStatus.status === "loading"}
  <InlineLoading status="active" description="checking" />
{:else if sqlStatus.status === "error"}
  <InlineLoading status="error" description={sqlStatus.error} />
{:else}
  <InlineLoading status="finished" description="no errors" />
{/if}
