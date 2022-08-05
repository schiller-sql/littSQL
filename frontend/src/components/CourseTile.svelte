<script lang="ts">
  import { ClickableTile, TooltipIcon } from "carbon-components-svelte";
  import { createEventDispatcher } from "svelte";
  import { TrashCan16 } from "carbon-icons-svelte";
  import type CourseListing from "../types/CourseListing";

  export let course: CourseListing;

  const dispatch = createEventDispatcher();

  function openCourse() {
    dispatch("open", { id: course.id });
  }

  function deleteCourse() {
    dispatch("delete", { id: course.id });
  }
</script>

<ClickableTile
  style="display: grid; grid-template-columns: 1fr auto; padding: 0; padding-right: 16px"
>
  <div on:click={openCourse} style="padding: 16px">
    {course.name}
  </div>
  <TooltipIcon
    tooltipText="delete course"
    icon={TrashCan16}
    on:click={deleteCourse}
    style="margin-left: 0.66rem"
  />
</ClickableTile>
