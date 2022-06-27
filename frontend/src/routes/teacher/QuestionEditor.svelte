<script lang="ts">
  import {
    RadioButton,
    RadioButtonGroup,
    TextArea,
    Toggle,
  } from "carbon-components-svelte";
  import MonoTextArea from "../../components/MonoTextArea.svelte";

  import type Question from "../../types/Question";
  import { letterFromNumber } from "../../util/utli";

  export let taskNumber: number;
  export let questionNumber: number;
  export let editable: boolean;
  export let question: Question;
  export let onQuestionEdit: () => void;

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
    }
    editedQuestion();
  }

  function editSolution(event) {
    question.solution = event.srcElement.value;
    editedQuestion();
  }
</script>

<h3>
  <span id="task-number">{letterFromNumber(taskNumber) + ") "}</span>
  <span id="question-number">{questionNumber + 1 + ". "}</span>
  {question.question}
</h3>

<div class="spacer" />
<div class="spacer" />

<TextArea
  invalid={question.question.length === 0}
  invalidText="do not let the question empty"
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
  <MonoTextArea
    invalid={question.solution.length === 0}
    invalidText="do not let the solution empty"
    placeholder="enter a solution..."
    labelText="solution"
    light
    mono={question.type === "sql"}
    value={question.solution}
    disabled={!editable}
    on:input={editSolution}
  />
{/if}

<style>
  .spacer {
    height: 16px;
  }

  span#task-number {
    font-size: 32px;
    font-weight: 700;
  }

  span#question-number {
    font-weight: 600;
  }
</style>
