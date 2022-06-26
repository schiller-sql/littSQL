<script lang="ts">
  import { Button, Tag, TextInput } from "carbon-components-svelte";

  import { Add20, CaretDown20, CaretUp20, Delete20 } from "carbon-icons-svelte";
  import { onMount } from "svelte";

  import type Task from "../types/Task";
  import { letterFromNumber } from "../util/utli";

  onMount(() => {
    taskDescription = task.description;
  });

  export let task: Task;
  export let taskNumber: number;
  let taskDescription: string;

  $: taskNumberDisplay = letterFromNumber(taskNumber);

  export let onDescriptionChange: () => void;
  export let onMove: (up: boolean) => void;
  export let onDelete: () => void;
  export let onNewQuestion: () => void;
  export let editable: boolean;

  function onInput(event) {
    const value = event.srcElement.value;
    task.description = value;
    onDescriptionChange();
  }
</script>

<!-- give a line next to each question on the level of the task -->

<li
  class="bx--tree-node"
  style="padding-left: 0; padding-bottom: 12px; padding-top: 4; "
>
  <div class="task-content">
    <div
      style="padding-top: 12px; padding-bottom: 12px; background-color: #222222"
    >
      <div style="font-size: large; display:inline-block">
        <span class="task-number">{taskNumberDisplay}.&nbsp</span>
        <input
          spellcheck={false}
          class="description"
          value={taskDescription}
          on:input={onInput}
        />
      </div>
    </div>
    <div style=" background-color: #222222">
      {#if task.is_voluntary}
        <Tag
          type="purple"
          size="sm"
          style="float: right; margin-right: 16px; margin-top: 24px; margin-bottom: 24px"
          >voluntary</Tag
        >
      {/if}
    </div>

    <div class="icons">
      {#if editable}
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
      {/if}
    </div>
  </div>
  <slot />
  <div style="margin-left: 28px; display: grid; grid-template-colums: 1fr">
    <Button
      disabled={!editable}
      size="small"
      kind="ghost"
      on:click={onNewQuestion}
      style="padding-left: 0.66rem; border-left: 2px solid #393939; line-height: 0; display: grid; grid-template-columns: auto 2px auto 1fr; width: 100%; max-width:none; height: 36px"
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
    background-color: #222222;
    grid-template-columns: auto 1fr auto 8px;
    padding-left: 16px;
  }
  input.description {
    text-decoration: underline double green;
    text-underline-position: under;
    background-color: transparent;
    color: white;
    font-size: 18px;
    font-weight: 500; /*TODO: better font weight*/
    background: linear-gradient(white, white) center bottom 5px /
      calc(100% - 15px) 1px no-repeat;
    padding: 7.5px;
    appearance: none;
    border: none;
  }
  input.description:focus {
    border: none;
    outline: none;
    appearance: none;
  }
</style>
