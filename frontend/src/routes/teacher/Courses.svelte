<script lang="ts">
    //import { onMount } from "svelte";
    import { authStore, fetchWithToken, requestWithToken } from "../../auth";
    import { Button, Loading, Modal, TextInput } from "carbon-components-svelte";
    import ProjectTile from "../../components/ProjectTile.svelte";
    import Add from "carbon-icons-svelte/lib/Add16/Add16.svelte";
    import { push } from "svelte-spa-router";
    import type CourseListing from "../../types/CourseListing";
    import CourseTile from "../../components/CourseTile.svelte";
    /*
    onMount(async () => {
        try {
            courses = await fetchWithToken("projects", "GET", $authStore.token);
            loading = false;
        } catch (e) {
            error = e.toString();
            console.log(e);
        }
    });
    */
    let open = false;
    let newCourseName: string;

    let loading = false;
    let error: string;
    let courses: CourseListing[] = [{id: 0, name: "test1"},{id: 1, name: "test2"},{id: 2, name: "test3"}];

    async function deleteCourse(event) {
        const id = event.detail.id;
        try {
            await requestWithToken(`courses/${id}`, "DELETE", $authStore.token);
            courses = courses.filter((project) => project.id != id);
        } catch (e) {
            console.log(e);
            error = "could not delete project";
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
            console.log(newCourse);
            // sort new course into courses
            courses = [...courses, newCourse].sort((a, b) => {
                if (a.name < b.name) return -1;
                if (a.name > b.name) return 1;
                return 0;
            });
        } catch (e) {
            console.log(e);
            error = "could not add project";
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
        <CourseTile {course} on:open={openCourse} on:delete={deleteCourse} />
    {/each}
    <Button on:click={() => (open = true)} icon={Add}>Create project</Button>
    <Modal
            bind:open
            modalHeading="Create new project"
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
{/if}
</body>
