<script lang="ts">
  import { onMount } from 'svelte'
  import { authStore, fetchWithToken } from '../../auth'
  import { Button, Loading, Modal, TextInput } from 'carbon-components-svelte'
  import ProjectTile from '../../components/ProjectTile.svelte'
  import type Project from '../../types/Project'
  import Add from 'carbon-icons-svelte/lib/Add.svelte'

  onMount(async () => {
    try {
      projects = await fetchWithToken('projects', 'GET', $authStore.token)
      loading = false
    } catch (e) {
      error = error.toString()
    }
  })
  let open = false
  let newProjectName: string

  let loading = true
  let error: string
  let projects: Project[] = []

  //TODO: send request to the server
  async function deleteProject(event) {
    const id = event.detail.id
    try {
    } catch (e) {}
    projects = projects.filter((project) => project.id != id)
  }

  //TODO: navigation
  function openProject(event) {
    const id = event.detail.id
  }

  //TODO: send request to the server
  function addProject() {}
</script>

<body>
  {#if loading}
    <Loading description="Active loading indicator" withOverlay={false} />
  {:else if error}
    <p style="color: red">{error.toString()}</p>
  {:else}
    {#each projects as project}
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
        if (text === 'Cancel') open = false
      }}
      on:open={() => (newProjectName = '')}
      on:submit={() => {
        open = false
        addProject()
      }}>
      <TextInput
        bind:value={newProjectName}
        labelText="Project name"
        helperText="The project name required"
        placeholder="Enter project name..." />
    </Modal>
  {/if}
</body>
