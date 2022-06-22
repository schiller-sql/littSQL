<script lang="ts">
  import type Question from "../types/Question";
  import { CaretDown16, CaretUp16, Delete16 } from "carbon-icons-svelte";

  export let question: Question;
  export let questionNumber: number;
  export let onMove: (up: boolean) => void;
  export let onDelete: () => void;
</script>

<li class="bx--tree-node" style="border-left: 2px solid grey;">
  <span class="question-number">{questionNumber}.&nbsp</span>
  {question.question}
  {question.solution}
  <div />
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
    grid-template-columns: auto auto 1fr auto auto auto;
  }
</style>
