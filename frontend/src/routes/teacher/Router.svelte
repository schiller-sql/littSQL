<script lang="ts">
  import Router from "svelte-spa-router";
  import wrap from "svelte-spa-router/wrap";
  import AlreadyLoggedIn from "../../components/AlreadyLoggedIn.svelte";
  import NotFound from "../../components/NotFound.svelte";
  import HomePage from "./HomePage.svelte";
  import ProjectEditPage from "./ProjectEditPage.svelte";
  import CoursesEditPage from "./CoursesEditPage.svelte";

  const routes = {
    "/": HomePage,
    // in case somebody presses back after logging in
    "/student-login": AlreadyLoggedIn,
    "/teacher-login": AlreadyLoggedIn,
    "/teacher-signup": AlreadyLoggedIn,
    "/courses/:coursesId": CoursesEditPage,
    "/projects/:projectId": ProjectEditPage,
    "*": NotFound,
  };
</script>

<main>
  <Router {routes} />
</main>
