<!--TODO: complete css layout; see: https://github.com/carbon-design-system/carbon-components-svelte/issues/503-->
<script lang="ts">
    import {Link, OverflowMenu, OverflowMenuItem, Tile} from "carbon-components-svelte";
    import Edit from "carbon-icons-svelte/lib/Edit16";
    import View from "carbon-icons-svelte/lib/View16";
    import type Project from "../types/Project";
    import { createEventDispatcher } from 'svelte';

    export let project: Project;

    const dispatch = createEventDispatcher();

    function openProject() {
        dispatch('open', {id: project.id})
    }
    function deleteProject() {
        dispatch('delete', {id: project.id})
    }
</script>

<Tile>
    {project.name}
    {#if !project.is_public}
        <Link size="lg" on:click={openProject} icon={Edit}/>
    {:else}
        <Link size="lg" on:click={openProject} icon={View}/>
    {/if}
    <OverflowMenu open flipped>
        {#if project.is_public}
            <OverflowMenuItem on:click={openProject} text="View project"/>
        {:else}
            <OverflowMenuItem on:click={openProject} text="Edit project"/>
            <OverflowMenuItem on:click={deleteProject} danger text="Delete project"/>
        {/if}
    </OverflowMenu>
</Tile>