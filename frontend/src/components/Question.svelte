<script lang="ts">
  import type Question from "../types/Question";
  import { CaretDown16, CaretUp16, Delete16 } from "carbon-icons-svelte";
  import { Tag } from "carbon-components-svelte";

  export let question: Question;
  export let questionNumber: number;
  $: questionNumberDisplay = (questionNumber + 1).toString();
  export let onMove: (up: boolean) => void;
  export let onDelete: () => void;
  export let selected: boolean;
</script>

<li
  class="bx--tree-node {selected ? 'selected-li' : ''}"
  style="border-left: 2px solid #393939; margin-left: 28px"
>
  <div style="padding-top: 10px; padding-bottom: 10px" on:click>
    <div style="display:inline-block">
      <span class="question-number">{questionNumberDisplay}.&nbsp</span>
      {question.question}
    </div>
  </div>
  <div on:click>
    <Tag
      type={question.type === "sql" ? "green" : "cool-gray"}
      size="sm"
      style="float: right; margin-right: 16px; margin-top: 9px; margin-bottom: 9px"
      >{question.type}</Tag
    >
    {#if question.solution !== null}
      <Tag
        type="blue"
        size="sm"
        style="float: right; margin-right: 8px; margin-top: 9px; margin-bottom: 9px"
        >solution</Tag
      >
    {/if}
  </div>
  <div class="icons">
    <div class="up-down-box">
      <CaretUp16
        on:click={() => onMove(true)}
        style="position: relative; top: 6px; "
      />
      <CaretDown16
        on:click={() => onMove(false)}
        style="position: relative; top: -6px; "
      />
    </div>
    <Delete16 on:click={onDelete} />
  </div>
</li>

<style>
  .icons {
    grid-template-columns: repeat(2, 1fr);
    align-items: center;
    column-gap: 0.5rem;
    display: flex;
    margin-right: 12px;
  }

  .up-down-box {
    display: grid;
    grid-template-rows: repeat(2, 1fr);
    row-gap: 4px;
  }

  .question-number {
    font-weight: 700;
  }

  li {
    display: grid;
    grid-template-columns: auto 1fr auto;
  }

  li:hover {
    background-color: #353535;
  }

  li:active {
    background-color: #525252;
  }

  .selected-li {
    background-color: #525252;
  }

  .selected-li:hover {
    background-color: #525252;
  }
</style>
