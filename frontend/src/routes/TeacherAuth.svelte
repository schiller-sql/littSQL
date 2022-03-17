<script lang="ts">
  import {
    Form,
    TextInput,
    PasswordInput,
    Button,
    ButtonSet,
    InlineNotification,
  } from "carbon-components-svelte";
  import { replace } from "svelte-spa-router";
  export let isLogin;
  let email = "";
  let password = "";
  let confirmPassword = "";

  let confirmedPasswordInvalid = false;
  let requestError: string | undefined;

  $: disabled =
    email.length == 0 ||
    password.length == 0 ||
    (isLogin ? false : confirmPassword.length == 0);

  function changePath() {
    replace(isLogin ? "/teacher-signup" : "/teacher-login");
  }
  async function submit() {
    requestError = undefined;
    confirmedPasswordInvalid = false;

    if (!isLogin) {
      if (password !== confirmPassword) {
        confirmedPasswordInvalid = true;
        return;
      }
    }
    const url = `http://localhost:8080/auth/${isLogin ? "login" : "signup"}`;
    const data = { email, password };
    const res = await fetch(url, {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(data),
    });
    console.log(res);
    if (!res.ok) {
      const json = await res.json();
      requestError = json.error.toLowerCase();
    }
  }
</script>

<Form on:submit={submit}>
  <TextInput bind:value={email} placeholder="Enter email address..." required />
  <PasswordInput
    bind:value={password}
    placeholder="Enter password..."
    required
  />
  {#if !isLogin}
    <PasswordInput
      bind:value={confirmPassword}
      invalid={confirmedPasswordInvalid}
      invalidText="Please repeat your password correctly"
      placeholder="Confirm password..."
      required
    />
  {/if}
  <ButtonSet>
    <Button type="submit" {disabled}>
      {#if isLogin} Login {:else} Sign up {/if}
    </Button>
    <Button kind="secondary" on:click={changePath}>
      {#if !isLogin} Go to login {:else} Sign up {/if}
    </Button>
  </ButtonSet>
</Form>

{#if requestError !== undefined}
  <InlineNotification
    lowContrast
    kind="error"
    title="Error:"
    subtitle={requestError.length == 0 ? "lamo" : requestError}
  />
{/if}
