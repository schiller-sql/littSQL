<script lang="ts">
  import { Button, CodeSnippet, Tile } from "carbon-components-svelte";
  import { Copy16, Delete16, Edit16, Renew16 } from "carbon-icons-svelte";
  import type Participant from "../types/Participant";

  export let participant: Participant;

  export let onDelete: (participant: Participant) => void;
  export let onEditName: (participant: Participant) => void;
  export let onRefreshAccessCode: (participant: Participant) => void;

  function _onDelete(): void {
    onDelete(participant);
  }

  function _onEditName(): void {
    onEditName(participant);
  }

  function _refreshAccessCode(): void {
    onRefreshAccessCode(participant);
  }

  function copyAccessCode() {
    if (navigator.clipboard && window.isSecureContext) {
      navigator.clipboard.writeText(participant.access_code);
    } else {
      let textArea = document.createElement("textarea");
      textArea.value = participant.access_code;
      textArea.style.position = "fixed";
      textArea.style.left = "-999999px";
      textArea.style.top = "-999999px";
      document.body.appendChild(textArea);
      textArea.focus();
      textArea.select();
      const ok = document.execCommand("copy");
      if (!ok) {
        alert("could not copy " + participant.access_code);
      }
      textArea.remove();
    }
  }

  let showAccessCode = false;
</script>

<Tile>
  <div class="tile-layout">
    <div class="edit-name-layout">
      <Button
        tooltipPosition="right"
        kind="ghost"
        size="small"
        iconDescription="edit name"
        icon={Edit16}
        on:click={_onEditName}
      >
        <div class:no-name={participant.name === null}>
          {participant.name ?? "no name"}
        </div>
      </Button>
      <div />
    </div>
    <div>
      {#if showAccessCode}
        <Tile
          style="padding:9px; min-height: 0; display: inline-block; min-width: 0; font-family: IBM Plex Mono"
          light>{participant.access_code}</Tile
        >
        <Button
          kind="secondary"
          size="small"
          iconDescription="copy access code"
          icon={Copy16}
          on:click={copyAccessCode}
        />
        <Button
          kind="secondary"
          size="small"
          iconDescription="refresh access code"
          icon={Renew16}
          on:click={_refreshAccessCode}
        />
      {:else}
        <Button
          style="width: 140.31px; padding-right: 0"
          kind="secondary"
          size="small"
          on:click={() => (showAccessCode = true)}>show access code</Button
        >
      {/if}
    </div>
    <div class="delete-participant-layout">
      <div />
      <Button
        tooltipPosition="left"
        kind="ghost"
        size="small"
        iconDescription="delete {participant.name ?? 'participant'}"
        icon={Delete16}
        on:click={_onDelete}
      />
    </div>
  </div>
</Tile>

<style>
  .tile-layout {
    display: grid;
    grid-template-columns: 1fr auto 1fr;
  }

  .edit-name-layout {
    display: grid;
    grid-template-columns: auto auto 1fr;
  }

  .delete-participant-layout {
    display: grid;
    grid-template-columns: 1fr auto;
  }

  .no-name {
    color: grey;
    background-color: #353535;
    font-size: 13px;
    border-radius: 3px;
    padding-left: 6px;
    padding-right: 6px;
  }
</style>
