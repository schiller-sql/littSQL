<script lang="ts">
  import { Modal } from "carbon-components-svelte";

  export let open: boolean = false;
  export let courseName: string;
</script>

<Modal
  on:submit
  danger
  bind:open
  modalHeading={`are you sure that you want to delete the course'${courseName}'?`}
  primaryButtonText="Delete"
  secondaryButtonText="Cancel"
  on:click:button--secondary={({ detail: { text } }) => {
    if (text === "Cancel") open = false;
  }}>Deleting a project is non-reversable.</Modal
>
