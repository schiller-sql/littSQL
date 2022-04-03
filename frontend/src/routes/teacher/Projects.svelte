<script lang="ts">
  import { authStore, DEFAULT_URL } from "../../auth";
  import { Loading } from "carbon-components-svelte";

  const url = DEFAULT_URL + "projects";
  const request = fetch(url, {
    method: "GET",
    headers: { Authorization: `Bearer ${$authStore.token}` },
  }).then((res) => res.json());
</script>

<body>
  {#await request}
    <Loading description="Active loading indicator" withOverlay={false} />
  {:then projects}
    {#each projects as project}
      <p>
        {`name des projectes: "${project.name}", ` +
          `id des projectes: ${project.id}, ` +
          `ist das project public?: ${project.is_public ? "ja" : "nein"}`}
      </p>
      <br /><br />
    {/each}
  {:catch error}
    <p style="color: red">{error.toString()}</p>
  {/await}
</body>
