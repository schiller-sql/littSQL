<script lang="ts">
  import { Loading } from "carbon-components-svelte";

  import { createMarkdownRenderStore } from "../util/md_util";

  export let rawMarkdown: string;
  export let light: boolean = false;
  const renderStore = createMarkdownRenderStore(350);

  $: renderStore.rawMarkdownUpdate(rawMarkdown);
</script>

<div
  class="markdown-render-container"
  class:light
  class:center-content={$renderStore.status === "loading"}
>
  {#if $renderStore.status === "rendered"}
    {@html $renderStore.renderedMarkdown}
  {:else}
    <Loading withOverlay={false} />
  {/if}
</div>

<style>
  div.markdown-render-container {
    padding: 12px;
    height: 100%;
    width: 100%;
    overflow-y: scroll;
  }

  div.center-content {
    display: flex;
    align-items: center;
    justify-content: center;
  }
  div.light {
    background-color: #393939;
  }
</style>
