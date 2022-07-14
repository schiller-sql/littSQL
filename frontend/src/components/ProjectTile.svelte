<script lang="ts">
  import { ClickableTile, TooltipIcon } from "carbon-components-svelte";
  import Edit16 from "carbon-icons-svelte/lib/Edit16/Edit16.svelte";
  import View16 from "carbon-icons-svelte/lib/View16/View16.svelte";
  import type ProjectListing from "../types/ProjectListing";
  import { createEventDispatcher } from "svelte";
  import { TrashCan16 } from "carbon-icons-svelte";

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
  style="display: grid; grid-template-columns: 1fr auto auto; padding: 0; padding-right: 16px"
>
  <div on:click={openProject} style="padding: 16px">
    {project.name}
  </div>
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
  {#if !project.is_public}
    <TooltipIcon
      tooltipText="delete project"
      icon={TrashCan16}
      on:click={deleteProject}
      style="margin-left: 0.66rem"
    />
  {/if}
</ClickableTile>
