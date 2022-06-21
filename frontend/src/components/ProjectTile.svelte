<script lang="ts">
  import {
    ClickableTile,
    OverflowMenu,
    OverflowMenuItem,
    TooltipIcon,
  } from "carbon-components-svelte";
  import Edit16 from "carbon-icons-svelte/lib/Edit16/Edit16.svelte";
  import View16 from "carbon-icons-svelte/lib/View16/View16.svelte";
  import type ProjectListing from "../types/ProjectListing";
  import { createEventDispatcher } from "svelte";

  export let project: ProjectListing;

  const dispatch = createEventDispatcher();

  function openProject() {
    dispatch("open", { id: project.id });
  }
  function deleteProject() {
    dispatch("delete", { id: project.id });
  }
</script>

<ClickableTile
  style="display: grid; grid-template-columns: 1fr auto auto"
  on:click={openProject}
>
  {project.name}
  {#if !project.is_public}
    <TooltipIcon
      tooltipText="private project"
      on:click={openProject}
      icon={Edit16}
    />
  {:else}
    <TooltipIcon
      tooltipText="public project"
      on:click={openProject}
      icon={View16}
    />
  {/if}
  <OverflowMenu flipped>
    {#if project.is_public}
      <OverflowMenuItem on:click={openProject} text="View project" />
    {:else}
      <OverflowMenuItem on:click={openProject} text="Edit project" />
      <OverflowMenuItem on:click={deleteProject} danger text="Delete project" />
    {/if}
  </OverflowMenu>
</ClickableTile>
