<script lang="ts">
  import Router from "svelte-spa-router";
  // @ts-ignore
  import routes from "routes.ts";
</script>

<main>
  <Router {routes} />
</main>
