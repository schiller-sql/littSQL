<script lang="ts">
  import { formatDistanceToNowStrict, isPast } from "date-fns";
  import {
    AccordionItem,
    Button,
    ButtonSet,
    Tag,
  } from "carbon-components-svelte";
  import {
    Delete16,
    Draggable24,
    Edit16,
    Locked16,
    Unlocked16,
  } from "carbon-icons-svelte";
  import type Assignment from "../types/Assignment";

  export let assignment: Assignment;
  export let onEditNameAndComment: (assignment: Assignment) => void;
  export let onSave: (assignment: Assignment) => void;
  export let onDelete: (assignment: Assignment) => void;

  function _onDelete() {
    onDelete(assignment);
  }

  function _onEditNameAndComment() {
    onEditNameAndComment(assignment);
  }

  function toggleLockAssignment() {
    assignment.locked = !assignment.locked;
    assignment = assignment;
    onSave(assignment);
  }

  $: finishedDate =
    assignment.finished_date !== null
      ? new Date(assignment.finished_date)
      : null;

  function assignmentToStatusColor(
    assignment: Assignment
  ): "outline" | "gray" | "green" {
    if (assignment.locked) {
      return "outline";
    } else if (finishedDate && isPast(finishedDate)) {
      return "gray";
    }
    return "green";
  }

  function assignmentToStatus(assignment: Assignment): string {
    if (assignment.locked) {
      return "locked";
    } else if (finishedDate) {
      if (isPast(finishedDate)) {
        return "finished";
      }
      return `finished in ${formatDistanceToNowStrict(finishedDate)}`;
    }
    return "open";
  }
</script>

<AccordionItem>
  <div slot="title" class="assignment-tile">
    <!-- fix css, make draggable icon in vertical center-->
    <Draggable24 style="margin-top: 2px" />
    <div />
    <div>
      <h4>
        {assignment.name}
      </h4>
      {#if assignment.comment !== null}
        <p class="comment">
          {assignment.comment}
        </p>
      {/if}
    </div>
    <Tag
      style="float: right; margin-right: 12px; height: 24px"
      type={assignmentToStatusColor(assignment)}
    >
      {assignmentToStatus(assignment)}
    </Tag>
  </div>
  <Button
    icon={Edit16}
    size="small"
    kind="secondary"
    on:click={_onEditNameAndComment}>Edit name & comment</Button
  >
  <Button
    icon={assignment.locked ? Unlocked16 : Locked16}
    size="small"
    kind="secondary"
    on:click={toggleLockAssignment}
    >{assignment.locked ? "Unlock" : "Lock"} assignment</Button
  >
  <Button icon={Delete16} kind="secondary" size="small" on:click={_onDelete}
    >Delete assignment</Button
  >
</AccordionItem>

<style>
  div.assignment-tile {
    display: grid;
    grid-template-columns: auto 12px 1fr auto;
  }

  p.comment {
    max-width: 75%;
    font-size: 16px;
    color: rgb(168, 168, 168);
  }
</style>
