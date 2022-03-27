<script lang="ts">
  import {
    Form,
    TextInput,
    PasswordInput,
    Button,
    ButtonSet,
    InlineNotification,
  } from 'carbon-components-svelte'
  import { replace } from 'svelte-spa-router'
  import { authStore, UserType } from '../../auth'
  export let isLogin
  let email = ''
  let password = ''
  let confirmPassword = ''

  let confirmedPasswordInvalid = false
  let requestError: string | undefined

  $: disabled =
    email.length == 0 ||
    password.length == 0 ||
    (isLogin ? false : confirmPassword.length == 0)

  function changePath() {
    replace(isLogin ? '/teacher-signup' : '/teacher-login')
  }
  async function submit() {
    requestError = undefined
    confirmedPasswordInvalid = false

    if (!isLogin) {
      if (password !== confirmPassword) {
        confirmedPasswordInvalid = true
        return
      }
    }
    const url = `http://localhost:8080/auth/${isLogin ? 'login' : 'signup'}`
    const data = { email, password }
    const res = await fetch(url, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(data),
    })

    if (!res.ok) {
      const text = await res.text()
      if (text.length === 0) {
        requestError = res.statusText
      } else {
        requestError = JSON.parse(text).error
      }
    } else {
      if (isLogin) {
        const json = await res.json()
        authStore.set({ jwt: json['token'], type: UserType.teacher })
      } else {
        replace('/teacher-login')
      }
    }
  }
</script>

<Form on:submit={submit}>
  <TextInput bind:value={email} placeholder="Enter email address..." required />
  <PasswordInput
    bind:value={password}
    placeholder="Enter password..."
    required />
  {#if !isLogin}
    <PasswordInput
      bind:value={confirmPassword}
      invalid={confirmedPasswordInvalid}
      invalidText="Please repeat your password correctly"
      placeholder="Confirm password..."
      required />
  {/if}
  <ButtonSet>
    <Button type="submit" {disabled}>
      {#if isLogin}Login{:else}Sign up{/if}
    </Button>
    <Button kind="secondary" on:click={changePath}>
      {#if !isLogin}Go to login{:else}Go to sign up{/if}
    </Button>
  </ButtonSet>
</Form>

{#if requestError !== undefined}
  <InlineNotification
    lowContrast
    kind="error"
    title="Error:"
    subtitle={requestError} />
{/if}

<Button kind="tertiary" href="#/student-login">Login as student?</Button>
