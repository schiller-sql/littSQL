import {wrap} from 'svelte-spa-router/wrap';

import NotFound from "./routes/NotFound.svelte";

import Home from "./routes/Home.svelte";
import StudentLogin from "./routes/StudentLogin.svelte";
import StudentHomepage from './routes/StudentHomepage.svelte';
import AssignmentPage from './routes/AssignmentPage.svelte';
import TeacherAuth from "./routes/TeacherAuth.svelte"
import TeacherHomepage from './routes/TeacherHomepage.svelte';

export default {
  "/": Home,
  "/student-login": StudentLogin,
  "/teacher-login": wrap({component: TeacherAuth, props: {isLogin: true}}),
  "/teacher-signup": wrap({ component: TeacherAuth, props: { isLogin: false } }),
  "/student-home": StudentHomepage,
  "/assignments": AssignmentPage,
  "/teacher-home": TeacherHomepage,
  "*": NotFound
};
