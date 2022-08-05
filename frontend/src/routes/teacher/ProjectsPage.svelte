<script lang="ts">
  import { onMount } from "svelte";
  import {
    Button,
    Modal,
    SkeletonPlaceholder,
    TextInput,
  } from "carbon-components-svelte";
  import ProjectTile from "../../components/ProjectTile.svelte";
  import type ProjectListing from "../../types/ProjectListing";
  import Add from "carbon-icons-svelte/lib/Add16/Add16.svelte";
  import { push } from "svelte-spa-router";
  import DeleteEntityModal from "../../components/DeleteEntityModal.svelte";
  import {
    fetchWithAuthorization,
    requestWithAuthorization,
  } from "../../util/auth_http_util";

  onMount(async () => {
    try {
      projects = await fetchWithAuthorization("projects", "GET");
      loading = false;
    } catch (e) {
      error = e.toString();
      console.error(e);
    }
  });
  let openCreateNewProjectModal = false;
  let pendingDeletionProject: ProjectListing | undefined;
  let openDeleteProjectModal = false;
  let newProjectName: string;

  let loading = true;
  let error: string;
  let projects: ProjectListing[] = [];

  function pendingDeletingProject(project: ProjectListing) {
    openDeleteProjectModal = true;
    pendingDeletionProject = project;
  }

  async function deleteProject() {
    try {
      await requestWithAuthorization(
        `projects/${pendingDeletionProject.id}`,
        "DELETE"
      );
      projects = projects.filter(
        (project) => project.id != pendingDeletionProject.id
      );
    } catch (e) {
      console.error(e);
      error = "could not delete project";
    } finally {
      openDeleteProjectModal = false;
    }
  }

  //TODO: navigation
  function openProject(event) {
    const id = event.detail.id;
    push(`#/projects/${id}`);
  }

  async function addProject() {
    try {
      const newProject = await fetchWithAuthorization(`projects`, "POST", {
        name: newProjectName,
      });
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
      console.error(e);
      error = "could not add project";
    }
  }
</script>

<body>
  {#if loading}
    <SkeletonPlaceholder style="" />
    <SkeletonPlaceholder style="" />
    <SkeletonPlaceholder style="" />
  {:else if error}
    <p style="color: red">{error.toString()}</p>
  {:else}
    {#each projects as project (project.id)}
      <ProjectTile
        {project}
        on:open={openProject}
        on:delete={() => pendingDeletingProject(project)}
      />
    {/each}
    <Button on:click={() => (openCreateNewProjectModal = true)} icon={Add}
      >Create project</Button
    >
    <Modal
      bind:open={openCreateNewProjectModal}
      selectorPrimaryFocus="#new-project-name"
      modalHeading="Create new project"
      primaryButtonText="Confirm"
      primaryButtonDisabled={!newProjectName}
      secondaryButtonText="Cancel"
      on:click:button--secondary={({ detail: { text } }) => {
        if (text === "Cancel") openCreateNewProjectModal = false;
      }}
      on:open={() => (newProjectName = "")}
      on:submit={() => {
        openCreateNewProjectModal = false;
        addProject();
      }}
    >
      <TextInput
        id="new-project-name"
        bind:value={newProjectName}
        labelText="Project name"
        helperText="The project name is required"
        placeholder="Enter project name..."
      />
    </Modal>
    <DeleteEntityModal
      entityName={pendingDeletionProject?.name ?? ""}
      entityType="project"
      bind:open={openDeleteProjectModal}
      on:submit={deleteProject}
    />
  {/if}
</body>
