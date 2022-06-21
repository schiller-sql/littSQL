<script lang="ts">
  import { Add16, CaretDown16, CaretUp16, Delete16 } from "carbon-icons-svelte";

  import type Task from "../types/Task";

  export let task: Task;
  export let taskNumber: number;
  export let onMove: (up: boolean) => void;
  export let onDelete: () => void;
  export let onNewQuestion: () => void;
</script>

<div class="outer-box">
  <div class="outer-box">
    <span class="task-number">{taskNumber}.</span>
    {task.description}
    <div class="up-down-box">
      <CaretUp16 style="height: 20px" on:click={() => onMove(true)} />
      <CaretDown16 style="height: 20px" on:click={() => onMove(false)} />
      <Delete16 on:click={onDelete} />
    </div>
  </div>
  <slot />
  <Add16 on:click={onNewQuestion} />
</div>

<style>
  .up-down-box {
    float: right;
    width: 10;
    flex-direction: column;
  }

  .task-number {
    font-weight: 700;
  }

  .outer-box {
    margin: 10em 0 1em 0;
  }
</style>
