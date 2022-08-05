<script lang="ts">
  import {
    Button,
    Loading,
    Modal,
    TextInput,
    Toggle,
  } from "carbon-components-svelte";

  import type Participant from "../../types/Participant";
  import {
    fetchWithAuthorization,
    requestWithAuthorization,
  } from "../../util/auth_http_util";
  import ParticipantTile from "../../components/ParticipantTile.svelte";
  import DeleteEntityModal from "../../components/DeleteEntityModal.svelte";

  export let params: {
    courseId: number;
  };

  let loading = true;
  let participants: Participant[] | undefined;
  let error: string | undefined;

  fetchWithAuthorization(`courses/${params.courseId}/participants`)
    .then((p) => {
      participants = p;
      loading = false;
    })
    .catch((e) => (error = e.toString()));

  let pendingDeletionParticipant: Participant | undefined;
  $: openDeleteParticipantModal = pendingDeletionParticipant !== undefined;

  function deleteParticpant(participant: Participant) {
    pendingDeletionParticipant = participant;
  }

  async function confirmDeleteParticpant() {
    try {
      await requestWithAuthorization(
        `courses/${params.courseId}/participants/${pendingDeletionParticipant.id}`,
        "DELETE"
      );
      participants = participants.filter(
        (participant) => participant.id !== pendingDeletionParticipant.id
      );
      pendingDeletionParticipant = undefined;
    } catch (e) {
      error = e.toString();
    }
  }

  async function refreshParticipantAccessCode(participant: Participant) {
    try {
      const { access_code } = await fetchWithAuthorization(
        `courses/${params.courseId}/participants/${participant.id}/refresh-access-code`,
        "PUT"
      );
      participant.access_code = access_code;
      participants = participants;
    } catch (e) {
      error = e.toString();
    }
  }

  let nameEditingParticipant: Participant | undefined;
  $: openNameEditingModal = nameEditingParticipant !== undefined;
  let nameEditingModalHasName: boolean;
  let nameEditingModalNewName: string;

  function editParticipantName(participant: Participant) {
    nameEditingParticipant = participant;
    nameEditingModalHasName = participant.name !== null;
    nameEditingModalNewName = participant.name;
  }

  async function editParticipantConfirm() {
    try {
      const newName = nameEditingModalHasName ? nameEditingModalNewName : null;
      await requestWithAuthorization(
        `courses/${params.courseId}/participants/${nameEditingParticipant.id}`,
        "PUT",
        { name: newName }
      );
      nameEditingParticipant.name = newName;
      nameEditingParticipant = undefined;
      participants = participants;
    } catch (e) {
      error = e.toString();
    }
  }
</script>

<!-- export and import to csv (https://stackoverflow.com/questions/14964035/how-to-export-javascript-array-info-to-csv-on-client-side) -->
{#if loading}
  <Loading />
{:else if error !== undefined}
  <p class="error">{error}</p>
{:else}
  {#each participants as participant}
    <ParticipantTile
      {participant}
      onDelete={deleteParticpant}
      onEditName={editParticipantName}
      onRefreshAccessCode={refreshParticipantAccessCode}
    />
  {/each}
{/if}

<DeleteEntityModal
  open={openDeleteParticipantModal}
  entityName={pendingDeletionParticipant?.name ?? " "}
  entityType="participant"
  on:submit={confirmDeleteParticpant}
/>

<Modal
  bind:open={openNameEditingModal}
  selectorPrimaryFocus="#edit-participant-name"
  modalHeading="Change name of particpant{nameEditingParticipant?.name
    ? " '" + nameEditingParticipant?.name + "'"
    : ''}"
  primaryButtonText="Confirm"
  primaryButtonDisabled={nameEditingModalHasName && !nameEditingModalNewName}
  secondaryButtonText="Cancel"
  on:click:button--secondary={({ detail: { text } }) => {
    if (text === "Cancel") nameEditingParticipant = undefined;
  }}
  on:submit={editParticipantConfirm}
>
  <Toggle
    labelA="without name"
    labelB="with name"
    labelText="With/without particpant name"
    size="sm"
    bind:toggled={nameEditingModalHasName}
  />
  {#if nameEditingModalHasName}
    <div class="spacer double" />
    <TextInput
      id="edit-participant-name"
      bind:value={nameEditingModalNewName}
      labelText="Participant name"
      helperText="The participant name has to consist of at least one letter"
      placeholder="Enter participant name..."
    />
  {/if}
</Modal>
