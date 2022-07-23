<script lang="ts">
  export let cellData: string | number | null;

  let dataType = typeof cellData;
</script>

<span
  class:number={dataType === "number"}
  class:null={dataType === "object"}
  class:string={dataType === "string"}
>
  {#if dataType === "string"}
    "{cellData}"
  {:else}
    {cellData}
  {/if}
</span>

<style>
  span.number {
    color: #a7f0ba;
  }

  span.null {
    color: #bb8eff;
  }

  span.string {
    color: #fa75a6;
  }
</style>
