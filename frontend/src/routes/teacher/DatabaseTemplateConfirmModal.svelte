<script lang="ts">
  import { Loading, Modal } from "carbon-components-svelte";
  import SqlTextArea from "../../components/SqlTextArea.svelte";
  import type DatabaseTemplate from "../../types/DatabaseTemplate";
  import { fetchWithAuthorization } from "../../util/auth_http_util";

  export let selectedId: number | null;
  export let onSelectTemplateSql: (sql: string) => void;

  function _onSelectTemplateSql() {
    if (databaseTemplate === undefined) return;
    selectedId = null;
    onSelectTemplateSql(databaseTemplate.sql);
  }

  const cachedDatabaseTemplates = new Map<number, DatabaseTemplate>();
  let loading: boolean = true;
  let error: string | undefined;
  let databaseTemplate: DatabaseTemplate;

  $: if (selectedId !== null) {
    loadDatabaseTemplate();
  }

  async function loadDatabaseTemplate() {
    if (cachedDatabaseTemplates.has(selectedId)) {
      databaseTemplate = cachedDatabaseTemplates.get(selectedId);
    } else {
      loading = true;
      try {
        databaseTemplate = await fetchWithAuthorization(
          `database-templates/${selectedId}`,
          "get"
        );
        cachedDatabaseTemplates.set(selectedId, databaseTemplate);
      } catch (e) {
        console.error(e);
        error = e.toString();
      }
    }
    loading = false;
  }
</script>

<Modal
  open={selectedId !== null}
  modalHeading={databaseTemplate?.description}
  modalLabel={databaseTemplate?.name}
  hasScrollingContent
  primaryButtonText="Use as database sql"
  secondaryButtonText="Cancel"
  on:click:button--secondary={() => (selectedId = null)}
  on:click:button--primary={_onSelectTemplateSql}
>
  {#if loading}
    <Loading />
  {:else if error !== undefined}
    <p class="error">
      {error}
    </p>
  {:else}
    <SqlTextArea popover code={databaseTemplate.sql} />
  {/if}
</Modal>
