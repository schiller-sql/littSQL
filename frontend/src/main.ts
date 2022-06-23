// @ts-nocheck
import LoggedInRouter from "./LoggedInRouter.svelte";
import "carbon-components-svelte/css/g100.css";
import "./styles/reset.scss";
import "./styles/font.scss";
import "./styles/colors.scss";
import "./styles/containers.scss";
import "./database";
import { initSqlite } from "./database";

initSqlite();

const app = new LoggedInRouter({
  target: document.body,
});

export default app;
