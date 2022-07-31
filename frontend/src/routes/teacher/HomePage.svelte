<script lang="ts">
  import { Tabs, Tab, TabContent } from "carbon-components-svelte";
  import { onMount } from "svelte";
  import Courses from "./Courses.svelte";
  import Projects from "./Projects.svelte";

  const lastSelectedTabKey = "home_page_last_selected_tab";

  onMount(() => {
    const rawLastSelectedTab = localStorage.getItem(lastSelectedTabKey);
    if (rawLastSelectedTab !== null) {
      selected = Number.parseInt(rawLastSelectedTab);
    } else {
      selected = 0;
    }
  });

  let selected: number;
  $: {
    if (selected !== undefined) {
      localStorage.setItem(lastSelectedTabKey, selected.toString());
    }
  }
</script>

<Tabs bind:selected>
  <Tab label="Courses" />
  <Tab label="Projects" />
  <svelte:fragment slot="content">
    <TabContent><Courses /></TabContent>
    <TabContent><Projects /></TabContent>
  </svelte:fragment>
</Tabs>
