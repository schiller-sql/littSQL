import NotFound from "./routes/404.svelte";

import Home from "./routes/Home.svelte";

export default {
  "/": Home,
  "*": NotFound,
};
