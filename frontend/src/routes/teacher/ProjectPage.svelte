<script lang="ts">
  import {
    Button,
    Loading,
    TextArea,
    TextInput,
  } from "carbon-components-svelte";

  import { Save20 } from "carbon-icons-svelte";

  import { onMount } from "svelte";
  import { authStore, fetchWithToken } from "../../auth";
  import type Project from "../../types/ProjectListing";

  export let params: {
    projectId: number;
  };

  onMount(async () => {
    try {
      project = await fetchWithToken(
        `projects/${params.projectId}`,
        "get",
        $authStore.token
      );
    } catch (e) {
      error = e.toString();
    } finally {
      loading = false;
    }
  });

  let loading = true;
  let error: string;
  let project: Project | undefined;

  function save() {}
</script>

{#if loading}
  <Loading description="Active loading indicator" withOverlay={false} />
{:else if error}
  <p style="color: red">{error.toString()}</p>
{:else}
  {JSON.stringify(project)}
  <Button on:click={save} icon={Save20}>Save</Button>
{/if}
