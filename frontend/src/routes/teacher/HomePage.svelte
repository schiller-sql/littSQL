<script lang="ts">
  import { Tabs, Tab, TabContent } from 'carbon-components-svelte'
  import Courses from './Courses.svelte'
  import Projects from './Projects.svelte'
</script>

<p>Logged in as a teacher!</p>

<!--TODO: Remember tab for going back -->
<Tabs>
  <Tab label="Courses" />
  <Tab label="Projects" />
  <svelte:fragment slot="content">
    <TabContent> <Courses /> </TabContent>
    <TabContent> <Projects /> </TabContent>
  </svelte:fragment>
</Tabs>
