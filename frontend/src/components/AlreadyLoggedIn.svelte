<script lang="ts">
  import { ButtonSet, Button } from 'carbon-components-svelte'
  import { authStore } from '../auth'
</script>

<p>You are already logged in</p>
<ButtonSet>
  <Button href="#/">Go back to home</Button>
  <Button kind="secondary" on:click={() => authStore.set(null)}>
    Log out instead
  </Button>
</ButtonSet>
