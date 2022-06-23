<script lang="ts">
  import { onMount } from "svelte";
  import { authStore, fetchWithToken, requestWithToken } from "../../auth";
  import { Button, Loading, Modal, TextInput } from "carbon-components-svelte";
  import ProjectTile from "../../components/ProjectTile.svelte";
  import type ProjectListing from "../../types/ProjectListing";
  import Add from "carbon-icons-svelte/lib/Add16/Add16.svelte";
  import { push } from "svelte-spa-router";

  onMount(async () => {
    try {
      projects = await fetchWithToken("projects", "GET", $authStore.token);
      loading = false;
    } catch (e) {
      error = e.toString();
      console.log(e);
    }
  });
  let open = false;
  let newProjectName: string;

  let loading = true;
  let error: string;
  let projects: ProjectListing[] = [];

  async function deleteProject(event) {
    const id = event.detail.id;
    try {
      await requestWithToken(`projects/${id}`, "DELETE", $authStore.token);
      projects = projects.filter((project) => project.id != id);
    } catch (e) {
      console.log(e);
      error = "could not delete project";
    }
  }

  //TODO: navigation
  function openProject(event) {
    const id = event.detail.id;
    push(`#/projects/${id}`);
  }

  async function addProject() {
    try {
      const newProject = await fetchWithToken(
        `projects`,
        "POST",
        $authStore.token,
        { name: newProjectName }
      );
      console.log(newProject);
      // sort new project into projects
      projects = [...projects, newProject].sort((a, b) => {
        if (a.is_public != b.is_public) {
          return a.is_public - b.is_public;
        }
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
    {#each projects as project (project.id)}
      <ProjectTile {project} on:open={openProject} on:delete={deleteProject} />
    {/each}
    <Button on:click={() => (open = true)} icon={Add}>Create project</Button>
    <Modal
      bind:open
      modalHeading="Create new project"
      primaryButtonText="Confirm"
      primaryButtonDisabled={!newProjectName}
      secondaryButtonText="Cancel"
      on:click:button--secondary={({ detail: { text } }) => {
        if (text === "Cancel") open = false;
      }}
      on:open={() => (newProjectName = "")}
      on:submit={() => {
        open = false;
        addProject();
      }}
    >
      <TextInput
        bind:value={newProjectName}
        labelText="Project name"
        helperText="The project name required"
        placeholder="Enter project name..."
      />
    </Modal>
  {/if}
</body>
