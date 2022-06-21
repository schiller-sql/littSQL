<script lang="ts">
  import {
    Add20,
    Add24,
    CaretDown16,
    CaretUp16,
    Delete16,
  } from "carbon-icons-svelte";

  import type Task from "../types/Task";

  export let task: Task;
  export let taskNumber: number;
  export let onMove: (up: boolean) => void;
  export let onDelete: () => void;
  export let onNewQuestion: () => void;
</script>

<!-- give a line next to each question on the level of the task -->

<li class="bx--tree-node">
  <div class="task-content">
    <span class="task-number">{taskNumber}. </span>
    {task.description}
    <div />
    <CaretUp16 on:click={() => onMove(true)} />
    <CaretDown16 on:click={() => onMove(false)} />
    <Delete16 on:click={onDelete} />
  </div>
  <slot />
  <a> <Add24 on:click={onNewQuestion} style="margin-left: 0.8rem;" /></a>
</li>

<style>
  .task-number {
    font-weight: 700;
  }
  li {
    padding: 8px;
    padding-top: 1rem;
  }

  .task-content {
    display: grid;
    grid-template-columns: auto auto 1fr auto auto auto 1rem;
    margin-bottom: 20px;
  }
</style>
