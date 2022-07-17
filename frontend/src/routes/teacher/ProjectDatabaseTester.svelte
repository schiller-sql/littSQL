<script lang="ts">
  import { Button } from "carbon-components-svelte";

  import { getContext } from "svelte";
  import SqlTextArea from "../../components/SqlTextArea.svelte";
  import type { QueryExecResult } from "../../sql-js/sql-wasm";

  const { projectId, databaseSql } = getContext("project-data");

  const localStorageTestQueryKey = `project:${projectId}:test_query`;
  let testQuery = localStorage.getItem(localStorageTestQueryKey);
  $: localStorage.setItem(localStorageTestQueryKey, testQuery);

  let result: QueryExecResult[] | undefined;
</script>

<label
  id="project-database-label"
  for="project-database"
  class:bx--toggle-input__label={true}
>
  test query on project database
</label>
<SqlTextArea
  id="project-database"
  code={testQuery}
  on:change={(e) => (testQuery = e["detail"])}
/>
<Button>Test</Button>
<Button kind="secondary">Show all tables</Button>

<style>
  #project-database-label {
    display: block;
  }
</style>
