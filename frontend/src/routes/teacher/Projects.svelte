<script lang="ts">
    import {authStore, DEFAULT_URL} from "../../auth";
    import {Loading} from "carbon-components-svelte";

const url = DEFAULT_URL+'projects';
const request = fetch(url, {
    method: 'GET',
    headers: {'Authorization': `Bearer ${$authStore.token}` },
})
</script>

<body>
{#await request}
    <Loading description="Active loading indicator" withOverlay={false}/>
{:then projects}
    <p>dummy</p>
{:catch error}
    <p style="color: red">{JSON.stringify(error)}</p>
{/await}
</body>