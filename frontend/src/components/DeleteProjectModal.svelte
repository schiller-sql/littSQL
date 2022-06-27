<script lang="ts">
  import { Modal } from "carbon-components-svelte";

  export let open: boolean = false;
  export let projectName: string;
</script>

<Modal
  on:submit
  danger
  bind:open
  modalHeading={`are you sure that you want to delete the project '${projectName}'?`}
  primaryButtonText="Delete"
  secondaryButtonText="Cancel"
  on:click:button--secondary={({ detail: { text } }) => {
    if (text === "Cancel") open = false;
  }}>Deleting a project is non-reversable.</Modal
>
