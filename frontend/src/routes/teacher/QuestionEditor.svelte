<script lang="ts">
  import {
    RadioButton,
    RadioButtonGroup,
    TextArea,
    Toggle,
  } from "carbon-components-svelte";
  import { onMount } from "svelte";
  import SqlTextArea from "../../components/SqlTextArea.svelte";
  import MonoTextArea from "../../components/SqlTextArea.svelte";

  import type Question from "../../types/Question";
  import { letterFromNumber } from "../../util/utli";

  export let taskNumber: number;
  export let questionNumber: number;
  export let editable: boolean;
  export let question: Question;
  export let onQuestionEdit: () => void;

  onMount(() => {
    solution = question.solution;
    console.log(question.solution);
    toggled = question.solution !== null;
  });

  let solution: string | undefined;
  let toggled: boolean;

  function editedQuestion() {
    onQuestionEdit();
    question = question;
  }

  function editQuestion(event) {
    console.log(event);
    question.question = event.srcElement.value;
    editedQuestion();
  }

  function selectRadio(event) {
    question.type = event.detail;
    editedQuestion();
  }

  function toggleHasSolution() {
    console.log(solution);
    if (solution !== null) {
      solution = null;
      toggled = false;
    } else {
      solution = "";
      toggled = true;
    }
    question.solution = solution;
    console.log(solution);
    editedQuestion();
  }

  function editSolution(event) {
    question.solution = event.srcElement.value;
    editedQuestion();
  }

  function editSqlSolution(event) {
    question.solution = event.detail;
    editedQuestion();
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
    value={solution !== null ? "on" : "off"}
    labelText="with solution"
    disabled={!editable}
    labelA="no solution"
    labelB="solution"
    on:change={toggleHasSolution}
    bind:toggled
  />
{/if}

<div class="spacer" />

{#if solution !== null}
  {#if question.type === "sql"}
    {"solution: " + solution}
    <SqlTextArea value={solution} on:change={editSqlSolution} />
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
  .spacer {
    height: 16px;
  }

  span#task-number {
    font-size: 28px;
    font-weight: 700;
  }

  span#question-number {
    font-size: 24px;
    font-weight: 600;
  }
</style>
