<script lang="ts">
  import { Delete16, Edit16 } from "carbon-icons-svelte";
  import { Button, Loading, Tile } from "carbon-components-svelte";

  import type Participant from "../../types/Participant";
  import { fetchWithAuthorization } from "../../util/auth_http_util";
  import ParticipantTile from "../../components/ParticipantTile.svelte";

  export let params: {
    courseId: number;
  };

  let loading = true;
  let participants: Participant[] | undefined;
  let error: string | undefined;

  fetchWithAuthorization(`courses/${params.courseId}/participants`)
    .then((p) => {
      participants = p;
      loading = false;
    })
    .catch((e) => (error = e.toString()));
</script>

<!-- export and import to csv (https://stackoverflow.com/questions/14964035/how-to-export-javascript-array-info-to-csv-on-client-side) -->
{#if loading}
  <Loading />
{:else if error !== undefined}
  <p class="error">{error}</p>
{:else}
  {#each participants as participant}
    <ParticipantTile
      id={participant.id}
      name={participant.name}
      access_code={participant.access_code}
    />
  {/each}
{/if}

<style>
</style>
