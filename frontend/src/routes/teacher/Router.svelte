<script lang="ts">
  import Router from "svelte-spa-router";
  import AlreadyLoggedIn from "../../components/AlreadyLoggedIn.svelte";
  import NotFound from "../../components/NotFound.svelte";
  import HomePage from "./HomePage.svelte";
  import ProjectEditPage from "./ProjectEditPage.svelte";
  import CourseEditPage from "./CourseEditPage.svelte";
  import ParticipantsPage from "./ParticipantsPage.svelte";

  const routes = {
    "/": HomePage,
    "/courses/:courseId": CourseEditPage,
    "/projects/:projectId": ProjectEditPage,
    "/courses/:courseId/participants": ParticipantsPage,
    // in case somebody presses back after logging in
    "/student-login": AlreadyLoggedIn,
    "/teacher-login": AlreadyLoggedIn,
    "/teacher-signup": AlreadyLoggedIn,
    "*": NotFound,
  };
</script>

<main>
  <Router {routes} />
</main>
