<script lang="ts">
  import { Button } from "carbon-components-svelte";

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
    <div style="padding-top: 35%; padding-bottom: 35%;">
      <span class="task-number">{taskNumber}.&nbsp</span>
      {task.description}
    </div>
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
  </div>
  <slot />
  <Button
    size="small"
    kind="ghost"
    on:click={onNewQuestion}
    style="align-items:center; display:flex; padding-left: 0.66rem; border-left: 2px solid grey; line-height: 0"
  >
    <Add20 />
    <span style="font-size:smaller"> add new question </span>
  </Button>
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

  .task-number {
    font-weight: 700;
    font-size: 18px;
  }
  /* li {
    padding: 8px;
    padding-right: 8px;
    padding-top: 0;
  } */

  .task-content {
    display: grid;
    grid-template-columns: auto 1fr auto 8px;
  }
</style>
