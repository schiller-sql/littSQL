<script lang="ts">
  import { ImageLoader, Form, TextInput, Button } from "carbon-components-svelte";

  let accesstoken = "";
  $: disabled = accesstoken.length != 6;
  $: invalid = accesstoken.length > 6;

  function submit() {

  }
</script>

<ImageLoader src=""></ImageLoader>
<Form on:submit={submit}>
  <TextInput placeholder="Enter access token..." required bind:value={accesstoken} invalid={invalid}></TextInput>
  <Button type="submit" disabled={disabled}>Login</Button>
</Form>
<Button kind="tertiary" href="#/teacher-login">Login as student?</Button>
