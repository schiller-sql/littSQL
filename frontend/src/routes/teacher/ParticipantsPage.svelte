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
  import { Add16 } from "carbon-icons-svelte";

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

  let isEditingParticipant = false;
  // if the participant that is currently being edited, does not exist, and has created
  let editingParticipantIsNew = false;
  let editingParticipant: Participant | undefined;
  let editingParticipantHasName: boolean;
  let editingParticipantName: string;

  function newParticpant() {
    isEditingParticipant = true;
    editingParticipantIsNew = true;
    editingParticipantHasName = true;
    editingParticipantName = "";
  }

  function editParticipant(participant: Participant) {
    isEditingParticipant = true;
    editingParticipantIsNew = false;
    editingParticipant = participant;
    editingParticipantHasName = participant.name !== null;
    editingParticipantName = participant.name;
  }

  function compareParticipants(p1: Participant, p2: Participant): number {
    const name1 = p1.name?.toUpperCase();
    const name2 = p2.name?.toUpperCase();
    if (name1 === name2) return p1.id - p2.id;
    if (name1 === null) return -1;
    if (name2 === null) return 1;
    if (name1 < name2) return -1;
    return 1;
  }

  async function editParticipantConfirm() {
    try {
      const newName = editingParticipantHasName ? editingParticipantName : null;
      if (!editingParticipantIsNew) {
        await requestWithAuthorization(
          `courses/${params.courseId}/participants/${editingParticipant.id}`,
          "PUT",
          { name: newName }
        );
        editingParticipant.name = newName;
        participants = participants;
      } else {
        const newParticipant = await fetchWithAuthorization(
          `courses/${params.courseId}/participants`,
          "POST",
          { name: newName }
        );
        participants = [...participants, newParticipant].sort(
          compareParticipants
        );
      }
    } catch (e) {
      error = e.toString();
    } finally {
      isEditingParticipant = false;
    }
  }

  let modalHeading: string;
  $: {
    if (editingParticipantIsNew) {
      modalHeading = "Create new particpant";
    } else {
      let participantName = "";
      if (editingParticipant?.name) {
        participantName = " '" + editingParticipant?.name + "'";
      }
      modalHeading = "Change name of participant" + participantName;
    }
  }
</script>

<!-- participants page should be embedded in the CourseEditPage, as it needs the title of the course (title could be: participants of course 1 or so) -->

<!-- export and import to csv (https://stackoverflow.com/questions/14964035/how-to-export-javascript-array-info-to-csv-on-client-side) -->
{#if loading}
  <Loading />
{:else if error !== undefined}
  <p class="error">{error}</p>
{:else}
  {#if participants.length === 0}
    <div class="missing-information">
      no students, press on 'Create new student' to add one
    </div>
  {/if}
  {#each participants as participant (participant.id)}
    <ParticipantTile
      {participant}
      onDelete={deleteParticpant}
      onEditName={editParticipant}
      onRefreshAccessCode={refreshParticipantAccessCode}
    />
  {/each}
  <Button on:click={newParticpant} icon={Add16}>Create new student</Button>
{/if}

<DeleteEntityModal
  open={openDeleteParticipantModal}
  entityName={pendingDeletionParticipant?.name ?? " "}
  entityType="student"
  on:submit={confirmDeleteParticpant}
/>

<Modal
  bind:open={isEditingParticipant}
  selectorPrimaryFocus="#edit-participant-name"
  {modalHeading}
  primaryButtonText="Confirm"
  primaryButtonDisabled={editingParticipantHasName && !editingParticipantName}
  secondaryButtonText="Cancel"
  on:click:button--secondary={({ detail: { text } }) => {
    if (text === "Cancel") isEditingParticipant = false;
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
  <div class="spacer double" />
  <TextInput
    id="edit-participant-name"
    bind:value={editingParticipantName}
    spellcheck="false"
    disabled={!editingParticipantHasName}
    labelText="Stundent name"
    helperText="The student name has to consist of at least one letter"
    placeholder="Enter student name..."
  />
</Modal>
