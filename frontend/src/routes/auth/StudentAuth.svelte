<script lang="ts">
  import {
    Form,
    TextInput,
    Button,
    InlineNotification,
  } from 'carbon-components-svelte'
  import { authStore, UserType } from '../../auth'

  let accessCode = ''
  $: disabled = accessCode.length != 6
  $: invalid = accessCode.length > 6

  let requestError: string | undefined

  async function submit() {
    requestError = undefined

    const url = 'http://localhost:8080/auth/login'
    const data = { access_code: accessCode }
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
      const json = await res.json()
      authStore.set({ jwt: json['token'], type: UserType.student })
    }
  }
</script>

<Form on:submit={submit}>
  <TextInput
    placeholder="Enter access token..."
    required
    bind:value={accessCode}
    {invalid} />
  <Button type="submit" {disabled}>Login</Button>
</Form>

{#if requestError !== undefined}
  <InlineNotification
    lowContrast
    kind="error"
    title="Error:"
    subtitle={requestError} />
{/if}

<Button kind="tertiary" href="#/teacher-login">Login as teacher?</Button>
