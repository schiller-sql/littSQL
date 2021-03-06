<script lang="ts">
  import { performanceModeStore } from "./performance";
  import { UserType, userTypeToString, type AuthState } from "./auth";
  import { authStore } from "./stores";
  import AuthRouter from "./routes/auth/Router.svelte";
  import TeacherRouter from "./routes/teacher/Router.svelte";
  import ParticipantRouter from "./routes/participant/Router.svelte";
  import { UserAvatarFilledAlt20, SettingsAdjust20 } from "carbon-icons-svelte";
  import { onDestroy } from "svelte";
  import {
    Header,
    HeaderAction,
    HeaderUtilities,
    HeaderPanelLinks,
    HeaderPanelLink,
    Content,
    HeaderPanelDivider,
    Toggle,
    Loading,
  } from "carbon-components-svelte";

  let lastAuthState: AuthState | undefined;
  const unsubscribe = authStore.subscribe((currentAuthState) => {
    if (lastAuthState && lastAuthState.status !== "autologin_loading") {
      window.location.hash = "";
    }
    lastAuthState = currentAuthState;
  });

  onDestroy(unsubscribe);

  function logOut(): void {
    authStore.logOut();
    openUserPanel = false;
  }

  let openUserPanel = false;
</script>

<!--TODO: feedback css on header logo when clicked-->
<Header company="littSQL" href="#/">
  {#if $authStore != null}
    <HeaderUtilities>
      <HeaderAction icon={SettingsAdjust20}>
        <HeaderPanelLinks>
          <HeaderPanelDivider>
            High performance mode, turn on to save battery or if device cpu is
            too slow, comes with user experience penalty. Will be saved in
            browser.</HeaderPanelDivider
          >
          <Toggle
            toggled={$performanceModeStore === "high"}
            on:toggle={(e) =>
              ($performanceModeStore = e.detail.toggled ? "high" : "low")}
            style="margin-left: 24px"
            labelText="performance mode (refresh to fully take effect)"
            size="sm"
            labelA="default performance"
            labelB="high performance mode"
          />
        </HeaderPanelLinks>
      </HeaderAction>
      <HeaderAction icon={UserAvatarFilledAlt20} bind:open={openUserPanel}>
        <HeaderPanelLinks>
          <HeaderPanelDivider>
            {#if $authStore.status === "logged_in"}
              logged in as a {userTypeToString($authStore.type)}
            {:else}
              not currently logged in
            {/if}
          </HeaderPanelDivider>
          {#if $authStore.status === "logged_in"}
            <HeaderPanelLink on:click={logOut}>Log out</HeaderPanelLink>
          {/if}
        </HeaderPanelLinks>
      </HeaderAction>
    </HeaderUtilities>
  {/if}
</Header>

<Content style="background-color: #1e1e1e">
  <main>
    {#if $authStore.status === "logged_in"}
      {#if $authStore.type === UserType.teacher}
        <TeacherRouter />
      {:else}
        <ParticipantRouter />
      {/if}
    {:else if $authStore.status === "logged_out"}
      <AuthRouter />
    {:else}
      <Loading />
    {/if}
  </main>
</Content>
