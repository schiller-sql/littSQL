<script lang="ts">
  import { authStore, UserType, userTypeToString } from "./auth";
  import AuthRouter from "./routes/auth/Router.svelte";
  import TeacherRouter from "./routes/teacher/Router.svelte";
  import ParticipantRouter from "./routes/participant/Router.svelte";
  import { UserAvatarFilledAlt16 } from "carbon-icons-svelte";
  import { onDestroy } from "svelte";
  import {
    Header,
    HeaderAction,
    HeaderUtilities,
    HeaderPanelLinks,
    HeaderPanelLink,
    Content,
    HeaderPanelDivider,
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
    authStore.set(null);
  }
</script>

<Header company="littSQL" href="#/">
  {#if $authStore != null}
    <HeaderUtilities>
      <HeaderAction
        icon={UserAvatarFilledAlt16}
        closeIcon={UserAvatarFilledAlt16}
      >
        <HeaderPanelLinks>
          <HeaderPanelDivider
            >logged in as a {userTypeToString(
              $authStore.type
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
    {#if $authStore == null}
      <AuthRouter />
    {:else if $authStore.type === UserType.teacher}
      <TeacherRouter />
    {:else}
      <ParticipantRouter />
    {/if}
  </main>
</Content>
