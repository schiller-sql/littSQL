<script lang="ts">
  import type Question from "../types/Question";
  import { CaretDown16, CaretUp16, Delete16 } from "carbon-icons-svelte";

  export let question: Question;
  export let questionNumber: number;
  export let onMove: (up: boolean) => void;
  export let onDelete: () => void;
</script>

<li class="bx--tree-node">
  <span class="question-number">{questionNumber}. </span>
  {question.question}
  {question.solution}
  <div />
  <CaretUp16 on:click={() => onMove(true)} />
  <CaretDown16 on:click={() => onMove(false)} />
  <Delete16 on:click={() => onDelete()} />
</li>

<style>
  .question-number {
    font-weight: 700;
  }
  li {
    padding: 8px;
    display: grid;
    grid-template-columns: auto auto 1fr auto auto auto;
  }
</style>
