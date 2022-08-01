<script lang="ts">
  import {
    SkeletonText,
    Button,
    Modal,
    TextInput,
    Loading,
  } from "carbon-components-svelte";
  import { Edit20, UserMultiple20 } from "carbon-icons-svelte";
  import { push } from "svelte-spa-router";
  import type Course from "../../types/Course";
  import {
    fetchWithAuthorization,
    requestWithAuthorization,
  } from "../../util/auth_http_util";
  import CourseAssignmentsEditPage from "./CourseAssignmentsEditPage.svelte";

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
</script>

{#if loading}
  <SkeletonText />
{:else if course !== undefined}
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
  </div>

  <h2>
    {course.name}
  </h2>
{:else}
  <p class="error">{error}</p>
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

<div class="spacer double" />

<CourseAssignmentsEditPage courseId={params.courseId} />

<style>
  div.top-buttons {
    float: right;
    display: grid;
    grid-template-columns: auto 4px auto;
  }
</style>
