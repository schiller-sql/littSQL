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

  let editingParticipant: Participant | undefined;
  $: openParticipantEditingModal = editingParticipant !== undefined;
  let editingParticipantHasName: boolean;
  let editingParticipantName: string;

  function editParticipantName(participant: Participant) {
    editingParticipant = participant;
    editingParticipantHasName = participant.name !== null;
    editingParticipantName = participant.name;
  }

  async function editParticipantConfirm() {
    try {
      const newName = editingParticipantHasName ? editingParticipantName : null;
      await requestWithAuthorization(
        `courses/${params.courseId}/participants/${editingParticipant.id}`,
        "PUT",
        { name: newName }
      );
      editingParticipant.name = newName;
      editingParticipant = undefined;
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
  bind:open={openParticipantEditingModal}
  selectorPrimaryFocus="#edit-participant-name"
  modalHeading="Change name of particpant{editingParticipant?.name
    ? " '" + editingParticipant?.name + "'"
    : ''}"
  primaryButtonText="Confirm"
  primaryButtonDisabled={editingParticipantHasName && !editingParticipantName}
  secondaryButtonText="Cancel"
  on:click:button--secondary={({ detail: { text } }) => {
    if (text === "Cancel") editingParticipant = undefined;
  }}
  on:submit={editParticipantConfirm}
>
  <Toggle
    labelA="without name"
    labelB="with name"
    labelText="With/without particpant name"
    size="sm"
    bind:toggled={editingParticipantHasName}
  />
  {#if editingParticipantHasName}
    <div class="spacer double" />
    <TextInput
      id="edit-participant-name"
      bind:value={editingParticipantName}
      labelText="Participant name"
      helperText="The participant name has to consist of at least one letter"
      placeholder="Enter participant name..."
    />
  {/if}
</Modal>
