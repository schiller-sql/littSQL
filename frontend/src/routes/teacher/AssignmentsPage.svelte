<script lang="ts">
  import {
    Tag,
    Accordion,
    AccordionItem,
    Button,
    InlineLoading,
    Modal,
    Toggle,
    TextInput,
  } from "carbon-components-svelte";
  import { dndzone } from "svelte-dnd-action";
  import { flip } from "svelte/animate";
  import { Add16 } from "carbon-icons-svelte";
  import {
    fetchWithAuthorization,
    requestWithAuthorization,
  } from "../../util/auth_http_util";
  import type Assignment from "../../types/Assignment";
  import AssignmentTile from "../../components/AssignmentTile.svelte";
  import DeleteEntityModal from "../../components/DeleteEntityModal.svelte";

  export let courseId: number;

  let loading = true;
  let assignments: (Assignment & { order: number })[] | undefined;
  let error: string | undefined;

  let saving = false;

  fetchWithAuthorization(`courses/${courseId}/assignments`)
    .then((a) => {
      assignments = a;
      loading = false;
      a.map((a, index) => (a.order = index));
    })
    .catch((e) => (error = e.toString()));

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

  let pendingDeleteAssignment: Assignment | undefined;
  let openDeleteAssignmentModal: boolean;

  function deleteAssignment(assignment: Assignment) {
    openDeleteAssignmentModal = true;
    pendingDeleteAssignment = assignment;
  }

  async function confirmDeleteAssignment() {
    const pendingDeleteAssignmentId = pendingDeleteAssignment.id;
    assignments = assignments.filter(
      (assignment) => assignment !== pendingDeleteAssignment
    );
    pendingDeleteAssignment = undefined;
    saving = true;
    try {
      await requestWithAuthorization(
        `courses/${courseId}/assignments/${pendingDeleteAssignmentId}`
      );
    } catch (e) {
      error = e.toString();
    }
    saving = false;
  }

  let isEditingAssignment = false;
  // if the participant that is currently being edited, does not exist, and has created
  let editingAssignmentIsNew = false;
  let editingAssignment: Assignment | undefined;
  let editingAssignmentName: string;
  let editingAssignmentHasComment: boolean;
  let editingAssignmentComment: string;

  function newAssignment() {
    isEditingAssignment = true;
    editingAssignmentIsNew = true;
    editingAssignmentName = "";
    editingAssignmentHasComment = false;
    editingAssignmentComment = "";
  }

  function editAssignment(assignment: Assignment) {
    isEditingAssignment = true;
    editingAssignmentIsNew = false;
    editingAssignment = assignment;
    editingAssignmentName = assignment.name;
    editingAssignmentHasComment = assignment.comment !== null;
    editingAssignmentComment = assignment.comment;
  }

  async function editAssignmentConfirm() {
    const newName = editingAssignmentName;
    const newComment = editingAssignmentHasComment
      ? editingAssignmentComment
      : null;
    if (!editingAssignmentIsNew) {
      editingAssignment.name = newName;
      editingAssignment.comment = newComment;
      assignments = assignments;
      saveEditedAssignment(editingAssignment);
    } else {
      try {
        const newAssignment = await fetchWithAuthorization(
          `courses/${courseId}/assignments`,
          "POSt",
          {
            name: newName,
            comment: newComment,
          }
        );
        assignments = [...assignments, newAssignment];
      } catch (e) {
        error = e.toString();
      }
    }
    isEditingAssignment = false;
  }

  async function saveEditedAssignment(assignment: Assignment) {
    saving = true;
    await requestWithAuthorization(
      `courses/${courseId}/assignments/${assignment.id}`,
      "PUT",
      assignment
    );
    saving = false;
  }

  let modalHeading: string;
  $: {
    if (editingAssignmentIsNew) {
      modalHeading = "Create new assignment";
    } else {
      let assignmentName = "";
      if (editAssignment?.name) {
        assignmentName = " '" + editAssignment?.name + "'";
      }
      modalHeading = "Edit assignment" + assignmentName;
    }
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
          <AssignmentTile
            {assignment}
            onSave={saveEditedAssignment}
            onEditNameAndComment={editAssignment}
            onDelete={deleteAssignment}
          />
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
  <Button size="field" icon={Add16} on:click={newAssignment}
    >Add new assignment</Button
  >
  {#if saving}
    <InlineLoading description="saving..." />
  {:else}
    <InlineLoading status="finished" description="saved" />
  {/if}
{/if}

<DeleteEntityModal
  bind:open={openDeleteAssignmentModal}
  entityName={pendingDeleteAssignment?.name}
  entityType="assignment"
  on:submit={confirmDeleteAssignment}
/>

<Modal
  hasForm={true}
  bind:open={isEditingAssignment}
  selectorPrimaryFocus="#edit-assignment-name"
  {modalHeading}
  primaryButtonText="Confirm"
  primaryButtonDisabled={!editingAssignmentName ||
    (editingAssignmentHasComment &&
      (editingAssignmentComment?.length ?? 0) < 12)}
  secondaryButtonText="Cancel"
  on:click:button--secondary={({ detail: { text } }) => {
    if (text === "Cancel") isEditingAssignment = false;
  }}
  on:submit={editAssignmentConfirm}
>
  <TextInput
    id="edit-assignment-name"
    bind:value={editingAssignmentName}
    spellcheck="false"
    labelText="Assignment name"
    helperText="The assignment name has to consist of at least one letter"
    placeholder="Enter assignment name..."
  />
  <div class="spacer double" />
  <Toggle
    labelA="without comment"
    labelB="with comment"
    labelText="Add a comment"
    size="sm"
    bind:toggled={editingAssignmentHasComment}
  />
  <div class="spacer double" />
  <TextInput
    bind:value={editingAssignmentComment}
    spellcheck="false"
    disabled={!editingAssignmentHasComment}
    invalid={editingAssignmentHasComment &&
      (editingAssignmentComment?.length ?? 0) < 12}
    invalidText="The comment has to consist of at least 12 letters"
    labelText="Comment (for students to read)"
    helperText="The comment has to consist of at least 12 letters"
    placeholder="Enter comment..."
    minlength={12}
  />
</Modal>

<style>
  div.line {
    background-color: rgb(57, 57, 57);
    height: 1px;
  }
</style>
