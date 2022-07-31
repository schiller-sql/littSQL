<script>
  import {
    SkeletonText,
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

  export let params;

  const course = fetchWithAuthorization(`courses/${params.courseId}`);

  let assignmentsLoading = true;
  let assignmentsSaving = false;
  let savingError;

  let assignmentsError;

  let assignments;

  fetchWithAuthorization(`courses/${params.courseId}/assignments`)
    .then((a) => {
      assignments = a;
      assignmentsLoading = false;
      a.map((a, index) => (a.order = index));
    })
    .catch((e) => (assignmentsError = e));

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
    assignmentsSaving = true;
    try {
      await requestWithAuthorization(
        `courses/${params.courseId}/assignments/${changedAssignmentId}/order`,
        "POST",
        changedAssignmentOrder
      );
    } catch (e) {
      savingError = e;
    }
    assignmentsSaving = false;
  }
</script>

{#await course}
  <SkeletonText />
{:then course}
  <h2>
    {course.name}
  </h2>
{:catch e}
  <p class="error">{e.error}</p>
{/await}

<div class="spacer double" />

{#if assignmentsError}
  <p class="error">{assignmentsError}</p>
{:else if assignmentsLoading}
  <Accordion skeleton align="start" />
{:else}
  <div style="background-color: rgb(57, 57, 57); height: 1px" />
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
          <AccordionItem align="start">
            <div
              slot="title"
              style="display: grid; grid-template-columns: auto 12px 1fr auto;"
            >
              <!-- fix css, make draggable in veritical middle -->
              <Draggable24 />
              <div />
              <div style="display: inline-block">
                <h5>
                  {assignment.name}
                </h5>
              </div>
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
  <div style="background-color: rgb(57, 57, 57); height: 1px" />
  <div class="spacer smaller" />
  <Button size="small" icon={Add16}>add new assignment</Button>
  {#if savingError}
    <InlineLoading status="error" description="saving error: {savingError}" />
  {:else if assignmentsSaving}
    <InlineLoading description="saving..." />
  {:else}
    <InlineLoading status="finished" description="saved" />
  {/if}
{/if}
