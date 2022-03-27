<script lang="ts">
  import { authStore, UserType } from './auth'
  import AuthRouter from './routes/auth/Router.svelte'
  import TeacherRouter from './routes/teacher/Router.svelte'
  import ParticipantRouter from './routes/participant/Router.svelte'
  import { onDestroy } from 'svelte'

  let firstVal = true
  const unsubscribe = authStore.subscribe((_) => {
    if (firstVal) {
      firstVal = false
    } else {
      window.location.hash = ''
    }
  })

  onDestroy(unsubscribe)
</script>

<main>
  {#if $authStore == null}
    <AuthRouter />
  {:else if $authStore.type === UserType.teacher}
    <TeacherRouter />
  {:else}
    <ParticipantRouter />
  {/if}
</main>
