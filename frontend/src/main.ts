// @ts-nocheck
import LoggedInRouter from "./LoggedInRouter.svelte";
import "carbon-components-svelte/css/g100.css";
import "./styles/reset.scss";
import "./styles/font.scss";
import "./styles/colors.scss";
import "./styles/containers.scss";
import "./stores/database";
import "./stores/performance";
import "./stores/global_stores";

const app = new LoggedInRouter({
  target: document.body,
});

export default app;
