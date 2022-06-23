<script lang="ts">
    import {
        ClickableTile,
        TooltipIcon,
    } from "carbon-components-svelte";
    import Edit16 from "carbon-icons-svelte/lib/Edit16/Edit16.svelte";
    import type ProjectListing from "../types/ProjectListing";
    import {createEventDispatcher} from "svelte";
    import {TrashCan16} from "carbon-icons-svelte";

    export let course: ProjectListing;

    const dispatch = createEventDispatcher();

    function openCourse() {
        dispatch("open", {id: course.id});
    }

    function deleteCourse() {
        dispatch("delete", {id: course.id});
    }
</script>

<ClickableTile on:click={openCourse} style="display: grid; grid-template-columns: 1fr auto auto">
        {course.name}
    <TooltipIcon tooltipTest="edit course" icon={Edit16} />
    <TooltipIcon tooltipTest="delete course" icon={TrashCan16} on:click={deleteCourse} style="margin-left: 0.66rem"/>
</ClickableTile>