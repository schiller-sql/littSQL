<script lang="ts">
  import wrap from "svelte-spa-router/wrap";

  import HomePage from "./HomePage.svelte";
  import StudentAuth from "./StudentAuth.svelte";
  import TeacherAuth from "./TeacherAuth.svelte";

  import NotFound from "../../components/NotFound.svelte";
  import Router from "svelte-spa-router/Router.svelte";

  const routes = {
    "/": HomePage,
    "/student-login": StudentAuth,
    "/teacher-login": wrap({
      component: TeacherAuth,
      props: { isLogin: true },
    }),
    "/teacher-signup": wrap({
      component: TeacherAuth,
      props: { isLogin: false },
    }),
    "*": NotFound,
  };
</script>

<Router {routes} />
