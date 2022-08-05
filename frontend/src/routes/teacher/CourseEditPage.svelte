<script lang="ts">
  import {
    SkeletonText,
    Button,
    Modal,
    TextInput,
    Loading,
  } from "carbon-components-svelte";
  import { Edit20, TrashCan20, UserMultiple20 } from "carbon-icons-svelte";
  import { pop, push } from "svelte-spa-router";
  import DeleteEntityModal from "../../components/DeleteEntityModal.svelte";
  import type Course from "../../types/Course";
  import {
    fetchWithAuthorization,
    requestWithAuthorization,
  } from "../../util/auth_http_util";
  import AssignmentsPage from "./AssignmentsPage.svelte";

  export let params: { courseId: number };

  let error: string | undefined;
  let loading = true;
  let course: Course;
  fetchWithAuthorization(`courses/${params.courseId}`)
    .then((c) => {
      course = c;
      loading = false;
    })
    .catch((e) => (error = e.toString()));

  let nameTextInputValue: string;

  let openEditNameModal: boolean;

  let savingName: boolean = false;

  function editName() {
    openEditNameModal = true;
    nameTextInputValue = course.name;
  }

  function openParticipantsEditPage() {
    push(`#/courses/${params.courseId}/participants`);
  }

  async function submitName() {
    savingName = true;
    try {
      await requestWithAuthorization(`courses/${params.courseId}`, "PUT", {
        name: nameTextInputValue,
      });
      course.name = nameTextInputValue;
      course = course;
    } catch (e) {
      error = e;
    } finally {
      openEditNameModal = false;
    }
  }

  let openDeleteModal = false;

  function deleteThisCourse() {
    openDeleteModal = true;
  }

  async function confirmDeleteThisCourse() {
    try {
      await requestWithAuthorization(`courses/${params.courseId}`, "DELETE");
      pop();
    } catch (e) {
      error = e;
      openDeleteModal = false;
    }
  }
</script>

{#if error !== undefined}
  <p class="error">{error}</p>
{:else if loading}
  <SkeletonText />
{:else}
  <div class="top-buttons">
    <Button
      icon={Edit20}
      size="small"
      kind="ghost"
      iconDescription="edit name"
      on:click={editName}
    />
    <div />
    <Button
      icon={UserMultiple20}
      size="small"
      kind="ghost"
      iconDescription="go to participants"
      on:click={openParticipantsEditPage}
    />
    <div />
    <Button
      icon={TrashCan20}
      size="small"
      kind="ghost"
      iconDescription="delete"
      on:click={deleteThisCourse}
    />
  </div>

  <h2>
    {course.name}
  </h2>
{/if}

<Modal
  modalHeading="Change project name"
  selectorPrimaryFocus="#name-text-input"
  primaryButtonText="Confirm"
  secondaryButtonText="Cancel"
  primaryButtonDisabled={!nameTextInputValue}
  bind:open={openEditNameModal}
  on:click:button--secondary={({ detail: { text } }) => {
    if (text === "Cancel") openEditNameModal = false;
  }}
  on:submit={submitName}
>
  <TextInput
    id="name-text-input"
    labelText="Course name"
    helperText="The course name is required"
    placeholder="Enter new course name..."
    bind:value={nameTextInputValue}
  />
  {#if savingName}
    <Loading />
  {/if}
</Modal>

{#if error === undefined}
  <div class="spacer double" />
  <AssignmentsPage courseId={params.courseId} />
{/if}

<DeleteEntityModal
  bind:open={openDeleteModal}
  on:submit={confirmDeleteThisCourse}
  entityName={course?.name}
  entityType="course"
/>

<style>
  div.top-buttons {
    float: right;
    display: grid;
    grid-template-columns: auto 4px auto 4px auto;
  }
</style>
