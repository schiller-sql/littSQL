<script lang="ts">
    import { Form, TextInput, PasswordInput, Button, ButtonSet } from "carbon-components-svelte";
    import { replace } from "svelte-spa-router";
    export let isLogin;

    function changePath() {
        replace(isLogin ? "/teacher-signup" : "/teacher-login");
    }
    function submit() {

    }
</script>

<Form on:submit={submit}>
    <TextInput placeholder="Enter email address..." required></TextInput>
    <PasswordInput placeholder="Enter password..." required></PasswordInput>
    {#if !isLogin}
    <PasswordInput placeholder="Repeat password..." required></PasswordInput>
    {/if}
    <ButtonSet>
        <Button type="submit">{#if isLogin} Login {:else} Sign up {/if}</Button>
        <Button kind="secondary" on:click={changePath}>{#if !isLogin} Go to login {:else} Sign up {/if}</Button>
    </ButtonSet>
</Form>
