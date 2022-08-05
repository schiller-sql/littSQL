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
  import { UserType } from "../../stores/auth";
  import { apiUrl } from "../../config";
  import { authStore } from "../../stores/global_stores";
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

    if (!isLogin && password !== confirmPassword) {
      confirmedPasswordInvalid = true;
      return;
    }
    const url = `${apiUrl}/auth/${isLogin ? "login" : "signup"}`;
    const data = { email, password };
    const res = await fetch(url, {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(data),
    });

    if (!res.ok) {
      const text = await res.text();
      if (text.length === 0) {
        requestError = res.statusText;
      } else {
        requestError = JSON.parse(text).error;
      }
    } else {
      if (isLogin) {
        const json = await res.json();
        authStore.logIn(json["token"], UserType.teacher);
      } else {
        replace("/teacher-login");
      }
    }
  }
</script>

<h2>Sign in as teacher</h2>

<div class="spacer double" />

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
      {#if isLogin}Login{:else}Sign up{/if}
    </Button>
    <Button kind="secondary" on:click={changePath}>
      {#if !isLogin}Sign up instead{:else}Sign up instead{/if}
    </Button>
  </ButtonSet>
</Form>

{#if requestError !== undefined}
  <InlineNotification
    lowContrast
    kind="error"
    title="Error:"
    subtitle={requestError}
  />
{:else}
  <div class="spacer" />
{/if}

<Button size="field" kind="tertiary" href="#/student-login"
  >Login as student instead?</Button
>
