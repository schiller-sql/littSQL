<script lang="ts">
  import {
    Button,
    Loading,
    TextArea,
    TextInput,
  } from "carbon-components-svelte";

  import { Save20 } from "carbon-icons-svelte";

  import { onMount } from "svelte";
  import { authStore, fetchWithToken, requestWithToken } from "../../auth";
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
  let error: string | undefined;
  let project: Project | undefined;
  $: hasData = project !== undefined;

  async function save() {
    loading = true;
    try {
      await requestWithToken(
        `projects/${params.projectId}`,
        "put",
        $authStore.token,
        project
      );
    } catch (e) {
      console.log(e);
      error = e;
    }
    loading = false;
  }
</script>

{#if error !== undefined}
  <p style="color: red">{error.toString()}</p>
{:else}
  {#if hasData}
    {JSON.stringify(project)}
  {/if}
  <Button disabled={loading} on:click={save} icon={Save20}>Save</Button>
{/if}
