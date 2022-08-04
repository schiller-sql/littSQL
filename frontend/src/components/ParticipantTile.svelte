<script lang="ts">
  import { Button, Tile } from "carbon-components-svelte";
  import { Copy16, Delete16, Edit16 } from "carbon-icons-svelte";

  export let id: number;
  export let name: string | null;
  export let access_code: string;

  export let onDelete: (id: number) => void;
  export let onEditName: (id: number, name: string) => void;

  function _onDelete(): void {
    onDelete(id);
  }

  function _onEditName(): void {
    onEditName(id, name);
  }

  let showAccessCode = false;
</script>

<Tile>
  <div class="tile-layout">
    <div class:no-name={name === null}>
      {name ?? "no name"}
    </div>
    <Button
      kind="ghost"
      size="small"
      iconDescription="edit name"
      icon={Edit16}
      on:click={_onEditName}
    />
    <div />
    <div>
      {#if showAccessCode}
        <Tile
          style="padding:9px; min-height: 0; display: inline-block; min-width: 0; font-family: IBM Plex Mono"
          light>{access_code}</Tile
        >
        <Button
          kind="secondary"
          size="small"
          iconDescription="copy access code"
          icon={Copy16}
        />
      {:else}
        <Button
          style="width: 104.84px; padding-right: 0"
          kind="secondary"
          size="small"
          on:click={() => (showAccessCode = true)}>show code</Button
        >
      {/if}
    </div>
    <div />
    <Button
      kind="ghost"
      size="small"
      iconDescription="delete {name ?? 'participant'}"
      icon={Delete16}
      on:click={_onDelete}
    />
  </div>
</Tile>

<style>
  .tile-layout {
    display: grid;
    grid-template-columns: auto auto 1fr auto 1fr auto;
  }

  .no-name {
    color: grey;
  }
</style>
