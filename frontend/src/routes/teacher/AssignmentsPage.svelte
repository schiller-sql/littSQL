<script lang="ts">
  import {
    Tag,
    Accordion,
    AccordionItem,
    Button,
    InlineLoading,
  } from "carbon-components-svelte";
  import { dndzone } from "svelte-dnd-action";
  import { flip } from "svelte/animate";
  import { Add16, Draggable24 } from "carbon-icons-svelte";
  import {
    fetchWithAuthorization,
    requestWithAuthorization,
  } from "../../util/auth_http_util";
  import type AssignmentListing from "../../types/AssignmentListing";

  export let courseId: number;

  let loading = true;
  let assignments: (AssignmentListing & { order: number })[] | undefined;
  let error: string | undefined;

  let saving = false;

  fetchWithAuthorization(`courses/${courseId}/assignments`)
    .then((a) => {
      assignments = a;
      loading = false;
      a.map((a, index) => (a.order = index));
    })
    .catch((e) => (error = e.toString()));

  function assignmentStatusToColor(status) {
    if (status === "locked") {
      return "outline";
    } else if (status === "finished") {
      return "gray";
    }
    return "green";
  }

  function reorderChange(e) {
    assignments = e.detail.items;
  }

  async function reorderPlace(e) {
    reorderChange(e);
    assignments.forEach((a, index) => (a.order = index));
    const changedAssignmentId = e.detail.info.id;
    const { order: changedAssignmentOrder } = assignments.find(
      (assignment) => assignment.id === changedAssignmentId
    );
    saving = true;
    try {
      await requestWithAuthorization(
        `courses/${courseId}/assignments/${changedAssignmentId}/order`,
        "POST",
        changedAssignmentOrder
      );
    } catch (e) {
      error = e.toString();
    }
    saving = false;
  }
</script>

{#if error}
  <p class="error">{error}</p>
{:else if loading}
  <Accordion skeleton align="start" />
{:else}
  <div class="line" />
  <Accordion>
    <section
      on:consider={reorderChange}
      on:finalize={reorderPlace}
      use:dndzone={{
        items: assignments,
        flipDurationMs: 250,
        dropTargetStyle: {},
      }}
    >
      {#each assignments as assignment (assignment.id)}
        <div animate:flip={{ duration: 250 }}>
          <AccordionItem>
            <div slot="title" class="assignment-tile">
              <!-- fix css, make draggable icon in vertical center-->
              <Draggable24 />
              <div />
              <h4>
                {assignment.name}
              </h4>
              <Tag
                style="float: right; margin-right: 12px"
                type={assignmentStatusToColor(assignment.status)}
              >
                {assignment.status}
              </Tag>
            </div>
            {#if assignment.comment !== null}
              <p>{assignment.comment}</p>
            {/if}
          </AccordionItem>
        </div>
      {/each}
    </section>
  </Accordion>
  <div class="line" />
  <div class="spacer smaller" />
  {#if assignments.length === 0}
    <div class="missing-information">
      no assignments, press on 'Add new assignment' to add one
    </div>
  {/if}
  <Button size="small" icon={Add16}>Add new assignment</Button>
  {#if saving}
    <InlineLoading description="saving..." />
  {:else}
    <InlineLoading status="finished" description="saved" />
  {/if}
{/if}

<style>
  div.line {
    background-color: rgb(57, 57, 57);
    height: 1px;
  }

  div.assignment-tile {
    display: grid;
    grid-template-columns: auto 12px 1fr auto;
  }
</style>
