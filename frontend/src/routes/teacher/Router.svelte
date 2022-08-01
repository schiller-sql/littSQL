<script lang="ts">
  import Router from "svelte-spa-router";
  import AlreadyLoggedIn from "../../components/AlreadyLoggedIn.svelte";
  import NotFound from "../../components/NotFound.svelte";
  import HomePage from "./HomePage.svelte";
  import ProjectEditPage from "./ProjectEditPage.svelte";
  import CourseEditPage from "./CourseEditPage.svelte";
  import CourseAssignmentsEditPage from "./CourseAssignmentsEditPage.svelte";
  import ParticipantsEditPage from "./ParticipantsEditPage.svelte";

  const routes = {
    "/": HomePage,
    // in case somebody presses back after logging in
    "/student-login": AlreadyLoggedIn,
    "/teacher-login": AlreadyLoggedIn,
    "/teacher-signup": AlreadyLoggedIn,
    "/courses/:courseId": CourseEditPage,
    "/projects/:projectId": ProjectEditPage,
    "/courses/:courseId/participants": ParticipantsEditPage,
    "*": NotFound,
  };
</script>

<main>
  <Router {routes} />
</main>
