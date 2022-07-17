<script lang="ts">
  import { ClickableTile, Loading, Modal } from "carbon-components-svelte";
  import { authStore, fetchWithToken } from "../../auth";
  import type DatabaseTemplateListing from "../../types/DatabaseTemplateListing";
  import DatabaseTemplateConfirmModal from "./DatabaseTemplateConfirmModal.svelte";

  export let open: boolean;
  export let onSelectTemplateSql: (sql: string) => void;

  function _onSelectTemplateSql(sql: string) {
    open = false;
    onSelectTemplateSql(sql);
  }

  let error: string | undefined;
  let loading: boolean = true;
  let databaseTemplates: DatabaseTemplateListing[] | undefined;

  let selectedId: number | null = null;

  $: if (open && databaseTemplates === undefined) {
    fetchWithToken("database-templates", "get", $authStore.token)
      .then((data) => (databaseTemplates = data))
      .catch((e) => {
        console.error(e);
        error = e;
      })
      .finally(() => (loading = false));
  }
</script>

<Modal
  modalHeading="view a database template by clicking on it"
  hasScrollingContent
  passiveModal
  bind:open
>
  {#if loading}
    <Loading />
  {:else if error !== undefined}
    <p class="error">
      {error}
    </p>
  {:else}
    <ul>
      {#each databaseTemplates as databaseTemplate}
        <ClickableTile on:click={() => (selectedId = databaseTemplate.id)}>
          <h5>{databaseTemplate.name}</h5>
          <div>{databaseTemplate.description}</div>
        </ClickableTile>
      {/each}
    </ul>
    <div class="spacer double" />
  {/if}
</Modal>
<DatabaseTemplateConfirmModal
  bind:selectedId
  onSelectTemplateSql={_onSelectTemplateSql}
/>
