import {wrap} from 'svelte-spa-router/wrap';

import NotFound from "./routes/NotFound.svelte";

import Home from "./routes/Home.svelte";
import StudentLogin from "./routes/StudentLogin.svelte";
import TeacherAuth from "./routes/TeacherAuth.svelte"

export default {
  "/": Home,
  "/student-login": StudentLogin,
  "/teacher-login": wrap({component: TeacherAuth, props: {isLogin:true}}),
  "/teacher-signup": wrap({component: TeacherAuth, props: {isLogin:false}}),
  "*": NotFound
};
