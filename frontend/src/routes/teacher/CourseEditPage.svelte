<script lang="ts">
  import { SkeletonText } from "carbon-components-svelte";
  import type Course from "../../types/Course";
  import { fetchWithAuthorization } from "../../util/auth_http_util";

  export let params: {
    courseId: number;
  };

  const course: Promise<Course> = fetchWithAuthorization(
    `courses/${params.courseId}`
  );
</script>

{#await course}
  <SkeletonText />
{:then course}
  <h2>
    {course.name}
  </h2>
{:catch e}
  <p class="error">{e.error}</p>
{/await}
