<script lang="ts">
  import { onMount } from "svelte";

  export let params: {
    projectId: number;
  };

  onMount(() => {
    console.log(params);
  });
</script>
