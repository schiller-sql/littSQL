<script lang="ts">
  import type Question from "../types/Question";
  import { CaretDown16, CaretUp16, Delete16 } from "carbon-icons-svelte";

  export let question: Question;
  export let questionNumber: number;
  export let onMove: (up: boolean) => void;
  export let onDelete: () => void;
</script>

<div class="outer-box">
  <span class="question-number">{questionNumber}.</span>
  {question.question}
  {question.solution}
  <div class="up-down-box">
    <CaretUp16 on:click={() => onMove(true)} />
    <CaretDown16 on:click={() => onMove(false)} />
    <Delete16 on:click={() => onDelete()} />
  </div>
</div>

<style>
  .up-down-box {
    float: right;
    width: 10;
    flex-direction: column;
  }

  .question-number {
    font-weight: 700;
  }

  .outer-box {
    padding: 1em;
    margin-left: 10em;
  }
</style>
