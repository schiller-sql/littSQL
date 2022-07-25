<script lang="ts">
  import { authStore, fetchWithToken, requestWithToken } from "../../auth";
  import { Button, Loading, Modal, TextInput } from "carbon-components-svelte";
  import Add from "carbon-icons-svelte/lib/Add16/Add16.svelte";
  import { push } from "svelte-spa-router";
  import type CourseListing from "../../types/CourseListing";
  import CourseTile from "../../components/CourseTile.svelte";
  import { onMount } from "svelte";
  import DeleteCourseModal from "../../components/DeleteCourseModal.svelte";

  onMount(async () => {
    try {
      courses = await fetchWithToken("courses", "GET", $authStore.token);
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
      await requestWithToken(
        `courses/${pendingDeletionCourse.id}`,
        "DELETE",
        $authStore.token
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
      const newCourse = await fetchWithToken(
        `courses`,
        "POST",
        $authStore.token,
        { name: newCourseName }
      );
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

<body>
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
    <Button on:click={() => (open = true)} icon={Add}>Create course</Button>
    <Modal
      bind:open
      modalHeading="Create new course"
      primaryButtonText="Confirm"
      primaryButtonDisabled={!newCourseName}
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
    <DeleteCourseModal
      courseName={pendingDeletionCourse?.name ?? ""}
      bind:open={openDeleteCourseModal}
      on:submit={deleteCourse}
    />
  {/if}
</body>
