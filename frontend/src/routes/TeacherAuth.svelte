<script lang="ts">
  import {
    Form,
    TextInput,
    PasswordInput,
    Button,
    ButtonSet,
  } from "carbon-components-svelte";
  import { replace } from "svelte-spa-router";
  export let isLogin;
  let email = "";
  let password = "";
  let confirmPassword = "";

  $: disabled =
    email.length == 0 ||
    password.length == 0 ||
    (isLogin ? false : confirmPassword.length == 0);

  function changePath() {
    replace(isLogin ? "/teacher-signup" : "/teacher-login");
  }
  function submit() {}
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
      placeholder="Repeat password..."
      required
    />
  {/if}
  <ButtonSet>
    <Button type="submit" {disabled}
      >{#if isLogin} Login {:else} Sign up {/if}</Button
    >
    <Button kind="secondary" on:click={changePath}
      >{#if !isLogin} Go to login {:else} Sign up {/if}</Button
    >
  </ButtonSet>
</Form>
