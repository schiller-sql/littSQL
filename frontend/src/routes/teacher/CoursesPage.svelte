<script lang="ts">
  import { Button, Loading, Modal, TextInput } from "carbon-components-svelte";
  import { Add16 } from "carbon-icons-svelte";
  import { push } from "svelte-spa-router";
  import type CourseListing from "../../types/CourseListing";
  import CourseTile from "../../components/CourseTile.svelte";
  import { onMount } from "svelte";
  import DeleteEntityModal from "../../components/DeleteEntityModal.svelte";
  import {
    fetchWithAuthorization,
    requestWithAuthorization,
  } from "../../util/auth_http_util";

  onMount(async () => {
    try {
      courses = await fetchWithAuthorization("courses", "GET");
      loading = false;
    } catch (e) {
      error = e.toString();
      console.error(e);
    }
  });

  let open = false;
  let newCourseName: string;

  let loading = true;
  let error: string;
  let courses: CourseListing[];

  let openDeleteCourseModal = false;
  let pendingDeletionCourse: CourseListing | undefined;

  async function deleteCourse() {
    try {
      await requestWithAuthorization(
        `courses/${pendingDeletionCourse.id}`,
        "DELETE"
      );
      courses = courses.filter(
        (course) => course.id != pendingDeletionCourse.id
      );
    } catch (e) {
      console.error(e);
      error = "could not delete course";
    } finally {
      openDeleteCourseModal = false;
    }
  }

  //TODO: navigation
  function openCourse(event) {
    const id = event.detail.id;
    push(`#/courses/${id}`);
  }

  async function addCourse() {
    try {
      const newCourse = await fetchWithAuthorization(`courses`, "POST", {
        name: newCourseName,
      });
      // sort new course into courses
      courses = [...courses, newCourse].sort((a, b) => {
        if (a.name < b.name) return -1;
        if (a.name > b.name) return 1;
        return 0;
      });
    } catch (e) {
      console.error(e);
      error = e;
    }
  }
</script>

{#if loading}
  <Loading description="Active loading indicator" withOverlay={false} />
{:else if error}
  <p style="color: red">{error.toString()}</p>
{:else}
  {#each courses as course (course.id)}
    <CourseTile
      {course}
      on:open={openCourse}
      on:delete={() => {
        openDeleteCourseModal = true;
        pendingDeletionCourse = course;
      }}
    />
  {/each}
  <Button on:click={() => (open = true)} icon={Add16}>Create new course</Button>
  <Modal
    bind:open
    modalHeading="Create new course"
    primaryButtonText="Confirm"
    primaryButtonDisabled={!newCourseName}
    spellcheck="false"
    secondaryButtonText="Cancel"
    on:click:button--secondary={({ detail: { text } }) => {
      if (text === "Cancel") open = false;
    }}
    on:open={() => (newCourseName = "")}
    on:submit={() => {
      open = false;
      addCourse();
    }}
  >
    <TextInput
      bind:value={newCourseName}
      labelText="Course name"
      helperText="The course name required"
      placeholder="Enter course name..."
    />
  </Modal>
  <DeleteEntityModal
    entityName={pendingDeletionCourse?.name ?? ""}
    entityType="course"
    bind:open={openDeleteCourseModal}
    on:submit={deleteCourse}
  />
{/if}
