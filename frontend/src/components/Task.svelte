<script lang="ts">
  import { Button } from "carbon-components-svelte";

  import { Add20, CaretDown20, CaretUp20, Delete20 } from "carbon-icons-svelte";

  import type Task from "../types/Task";
  const alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ";

  export let task: Task;
  export let taskNumber: number;
  $: taskNumberDisplay = alphabet[taskNumber];
  export let onMove: (up: boolean) => void;
  export let onDelete: () => void;
  export let onNewQuestion: () => void;
</script>

<!-- give a line next to each question on the level of the task -->

<li class="bx--tree-node" style="padding-bottom: 12px; padding-top: 4">
  <div class="task-content">
    <div style="padding-top: 12px; padding-bottom: 12px;">
      <div style="font-size: large">
        <span class="task-number">{taskNumberDisplay}.&nbsp</span>
        {task.description}
      </div>
    </div>
    <div />
    <div class="icons">
      <div class="up-down-box">
        <CaretUp20
          on:click={() => onMove(true)}
          style="position: relative; top: 7px;"
        />
        <CaretDown20
          on:click={() => onMove(false)}
          style="position: relative; top: -7px;"
        />
      </div>
      <Delete20 on:click={onDelete} />
    </div>
  </div>
  <slot />
  <div style="margin-left: 28px; display: grid; grid-template-colums: 1fr">
    <Button
      size="small"
      kind="ghost"
      on:click={onNewQuestion}
      style="padding-left: 0.66rem; border-left: 2px solid grey; line-height: 0; display: grid; grid-template-columns: auto 2px auto 1fr; width: 100%; max-width:none; height: 36px"
    >
      <Add20 style="display:block" />
      <div />
      <div style="display:grid grid-template-columns: auto 1fr">
        <span style="font-size: smaller; float:left; text-align:left"
          >add new question</span
        >
        <div />
      </div>
      <div />
    </Button>
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

  .task-number {
    font-weight: 700;
    font-size: 18px;
  }

  .task-content {
    display: grid;
    grid-template-columns: auto 1fr auto 8px;
  }
</style>
