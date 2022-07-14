<script lang="ts">
  import { Modal } from "carbon-components-svelte";

  export let open: boolean;
  export let onLeave: () => void;
</script>

<Modal
  danger
  modalHeading="Are you sure you want to leave this page?"
  primaryButtonText="Leave"
  secondaryButtonText="Cancel"
  bind:open
  on:click:button--primary={onLeave}
  on:click:button--secondary={() => (open = false)}
>
  <p>You still have unsaved changes</p>
</Modal>
