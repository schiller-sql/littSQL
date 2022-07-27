<script lang="ts">
  import { performanceModeStore } from "./performance";
  import { UserType, userTypeToString } from "./auth";
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

  let firstVal = true;
  const unsubscribe = authStore.subscribe((_) => {
    if (firstVal) {
      firstVal = false;
    } else {
      window.location.hash = "";
    }
  });

  onDestroy(unsubscribe);

  function logOut(): void {
    authStore.logOut();
  }
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
      <HeaderAction icon={UserAvatarFilledAlt20}>
        <HeaderPanelLinks>
          <HeaderPanelDivider
            >logged in as a {userTypeToString(
              authStore.getUserType()
            )}</HeaderPanelDivider
          >
          <HeaderPanelLink on:click={logOut}>Log out</HeaderPanelLink>
        </HeaderPanelLinks>
      </HeaderAction>
    </HeaderUtilities>
  {/if}
</Header>

<Content>
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
