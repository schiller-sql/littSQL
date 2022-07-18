<script lang="ts">
  import { Button } from "carbon-components-svelte";

  import { getContext } from "svelte";
  import SqlResults from "../../components/SqlResults.svelte";
  import SqlTextArea from "../../components/SqlTextArea.svelte";
  import type { QueryExecResult } from "../../sql-js/sql-wasm";
  import { execStatementOnDatabase } from "../../util/db_util";

  export let projectSqlHasError: boolean;

  const projectData: { id: number; sql: string } = getContext("project-data");

  const localStorageTestQueryKey = `project:${projectData.id}:test_query`;
  let testQuery = localStorage.getItem(localStorageTestQueryKey) ?? "";
  $: localStorage.setItem(localStorageTestQueryKey, testQuery);

  let result: QueryExecResult[] | string | undefined;

  function executeTestSql() {
    result = execStatementOnDatabase(projectData.sql, testQuery);
  }
</script>

<label
  id="project-database-tester-label"
  for="project-database-tester"
  class:bx--toggle-input__label={true}
>
  test query on project database (will only be saved in local browser)
</label>
<div class="spacer smaller" />
<SqlTextArea
  id="project-database"
  maxHeight
  code={testQuery}
  on:change={(e) => (testQuery = e["detail"])}
/>
<Button disabled={projectSqlHasError} size="small" on:click={executeTestSql}
  >Test</Button
>

{#if result !== undefined}
  <SqlResults {result} />
{/if}

<style>
  #project-database-tester-label {
    display: block;
  }
</style>
