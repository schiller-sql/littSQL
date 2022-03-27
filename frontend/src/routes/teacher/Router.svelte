<script lang="ts">
  import Router from 'svelte-spa-router'
  import AlreadyLoggedIn from '../../components/AlreadyLoggedIn.svelte'
  import NotFound from '../../components/NotFound.svelte'
  import HomePage from './HomePage.svelte'

  const routes = {
    '/': HomePage,
    // in case somebody presses back after logging in
    '/student-login': AlreadyLoggedIn,
    '/teacher-login': AlreadyLoggedIn,
    '/teacher-signup': AlreadyLoggedIn,
    '*': NotFound,
  }
</script>

<main>
  <Router {routes} />
</main>
