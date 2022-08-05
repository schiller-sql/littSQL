<script lang="ts">
  import { ButtonSet, Button } from "carbon-components-svelte";
  import { authStore } from "../stores/global_stores";
</script>

<h4>You are already logged in</h4>
<div class="spacer" />
<ButtonSet>
  <Button href="#/">Go back to home</Button>
  <Button kind="secondary" on:click={() => authStore.logOut()}>
    Log out instead
  </Button>
</ButtonSet>
