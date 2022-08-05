<script lang="ts">
  import { Tabs, Tab, TabContent } from "carbon-components-svelte";
  import { onMount } from "svelte";
  import CoursesPage from "./CoursesPage.svelte";
  import ProjectsPage from "./ProjectsPage.svelte";

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
    <TabContent><CoursesPage /></TabContent>
    <TabContent><ProjectsPage /></TabContent>
  </svelte:fragment>
</Tabs>
