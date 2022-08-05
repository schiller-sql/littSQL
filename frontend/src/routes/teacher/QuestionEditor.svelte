<script lang="ts">
  import {
    Button,
    RadioButton,
    RadioButtonGroup,
    TextArea,
    ToastNotification,
    Toggle,
  } from "carbon-components-svelte";
  import { getContext, onMount } from "svelte";
  import SqlResults from "../../components/SqlResults.svelte";
  import SqlStatus from "../../components/SqlStatus.svelte";
  import SqlTextArea from "../../components/SqlTextArea.svelte";
  import type { QueryExecResult } from "../../sql-js/sql-wasm";
  import { performanceStore } from "../../stores/global_stores";

  import type Question from "../../types/Question";
  import { createSqlStatusStore } from "../../stores/sql_status";
  import { letterFromNumber } from "../../util/utli";
  import { execStatementOnDatabase } from "../../util/db_util";

  export let taskNumber: number;
  export let questionNumber: number;
  export let editable: boolean;
  export let question: Question;
  export let onQuestionEdit: () => void;
  export let databaseError: string | undefined;

  let projectData: { id: number; sql: string } = getContext("project-data");

  let sqlStatusStore = createSqlStatusStore(
    performanceStore.getCurrentMode() === "high" ? 800 : 500
  );
  onMount(() => {
    if (question.type === "sql" && question.solution !== null) {
      sqlStatusStore.sqlUpdate(question.solution, projectData.sql);
    }
  });

  $: hasSolution = question.solution !== null;

  function editedQuestion() {
    onQuestionEdit();
    question = question;
  }

  function editQuestion(event) {
    question.question = event.srcElement.value;
    editedQuestion();
  }

  function selectRadio(event) {
    question.type = event.detail;
    editedQuestion();
  }

  function toggleHasSolution() {
    const hasSolution = question.solution !== null;
    if (hasSolution) {
      question.solution = null;
    } else {
      question.solution = "";
      sqlStatusStore.sqlUpdate("");
    }
    editedQuestion();
  }

  function editSolution(event) {
    question.solution = event.srcElement.value;
    editedQuestion();
  }

  function editSqlSolution(event) {
    let q = question;
    q.solution = event.detail;
    onQuestionEdit();
    sqlStatusStore.sqlUpdate(event.detail, projectData.sql);
    solutionResult = undefined;
  }

  let solutionResult: QueryExecResult[] | string | undefined;

  function testSqlSoution() {
    solutionResult = execStatementOnDatabase(
      projectData.sql,
      question.solution
    );
    // TODO: Scroll down
  }
</script>

<h4>
  <span id="task-number">{letterFromNumber(taskNumber) + ") "}</span>
  <span id="question-number">{questionNumber + 1 + ". "}</span>
  {question.question}
</h4>

<div class="spacer" />
<div class="spacer" />

<TextArea
  invalid={question.question.length === 0}
  invalidText="do not let the question empty"
  rows={2}
  light
  value={question.question}
  placeholder="ask a question..."
  labelText="question"
  disabled={!editable}
  on:input={editQuestion}
/>

<div class="spacer" />

<RadioButtonGroup
  legendText="type of question"
  selected={question.type}
  disabled={!editable}
  on:change={selectRadio}
>
  <RadioButton labelText="sql query" value="sql" />
  <RadioButton labelText="text question" value="text" />
</RadioButtonGroup>

{#if editable}
  <div class="spacer" />

  <Toggle
    size="sm"
    value={hasSolution ? "on" : "off"}
    labelText="with solution"
    disabled={!editable}
    labelA="no solution"
    labelB="solution"
    on:change={toggleHasSolution}
    bind:toggled={hasSolution}
  />
{/if}

<div class="spacer" />

{#if hasSolution}
  {#if question.type === "sql"}
    <SqlTextArea code={question.solution} on:change={editSqlSolution} />
    <Button
      size="small"
      on:click={testSqlSoution}
      disabled={databaseError !== undefined || $sqlStatusStore.status !== "ok"}
      >test solution</Button
    >
    {#if databaseError !== undefined}
      <div class="spacer" />
      <ToastNotification
        lowContrast
        kind="error"
        title="Error in database, solution not testable"
        subtitle={databaseError}
        hideCloseButton
      />
    {:else}
      <SqlStatus sqlStatus={$sqlStatusStore} />
      {#if solutionResult !== undefined}
        <SqlResults result={solutionResult} />
      {/if}
    {/if}
  {:else}
    <TextArea
      invalid={question.solution.length === 0}
      invalidText={"do not let the solution empty"}
      placeholder={"enter a solution..."}
      labelText="solution"
      light
      value={question.solution}
      disabled={!editable}
      on:input={editSolution}
    />
  {/if}
{/if}

<style>
  span#task-number {
    font-size: 28px;
    font-weight: 700;
  }

  span#question-number {
    font-size: 24px;
    font-weight: 600;
  }
</style>
