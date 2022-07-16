<script lang="ts">
  import {
    Button,
    ButtonSkeleton,
    InlineNotification,
    SkeletonText,
    Tab,
    TabContent,
    Tabs,
    TabsSkeleton,
    TextArea,
    TextAreaSkeleton,
    TextInput,
    TextInputSkeleton,
    Toggle,
  } from "carbon-components-svelte";

  import {
    Save20,
    Delete20,
    Add24,
    Close24,
    Undo20,
  } from "carbon-icons-svelte";

  import { afterUpdate, onDestroy, onMount } from "svelte";
  import TaskComponent from "../../components/Task.svelte";
  import QuestionComponent from "../../components/Question.svelte";
  import { authStore, fetchWithToken, requestWithToken } from "../../auth";
  import type Project from "../../types/Project";
  import type Question from "../../types/Question";
  import type Task from "../../types/Task";
  import QuestionEditor from "./QuestionEditor.svelte";
  import { pop } from "svelte-spa-router";
  import DeleteProjectModal from "../../components/DeleteProjectModal.svelte";
  import SqlTextArea from "../../components/SqlTextArea.svelte";
  import DatabaseTemplatesModal from "./DatabaseTemplatePickerModal.svelte";
  import DatabaseTemplatePickerModal from "./DatabaseTemplatePickerModal.svelte";
  import UnsavedNavigateBackModal from "../../components/UnsavedNavigateBackModal.svelte";

  // -- initially fetch project using id provided by params --
  export let params: {
    projectId: number;
  };

  onMount(initialProjectFetch);

  // -- scrolling down if new task was added --
  let shouldScrollOnNextTickDown = false;
  let tasksScrollNode; // the DOM-node to scroll down when adding new tasks

  function shouldScrollDown() {
    shouldScrollOnNextTickDown = true;
  }

  afterUpdate(() => {
    if (shouldScrollOnNextTickDown) {
      tasksScrollNode.scrollTop = tasksScrollNode.scrollHeight;
      shouldScrollOnNextTickDown = false;
    }
  });

  // -- fetching and saving project,
  // initialProjectFetch is called on mount
  // and save when pressed on save button --
  let loading = true;
  let error: string | undefined;
  let project: Project | undefined;
  let edited = false; // if the project has been edited,
  // the save button is only showed if this is true
  $: projectIsPrivate = !project?.is_public; // if this is false, everything should not be editable

  /// called on mount
  async function initialProjectFetch() {
    try {
      project = await fetchWithToken(
        `projects/${params.projectId}`,
        "get",
        $authStore.token
      );
      addCurrentProjectToHistory();
    } catch (e) {
      error = e.toString();
    } finally {
      loading = false;
    }
  }

  async function save() {
    try {
      await requestWithToken(
        `projects/${params.projectId}`,
        "put",
        $authStore.token,
        project
      );
    } catch (e) {
      console.error(e);
      error = e;
    }
    edited = false;
  }

  // -- deleting project by first calling pendingDeletingProject,
  // which opens a modal, which calls deleteProject on confirmation --
  let openDeleteProjectModal = false;

  function pendingDeletingProject() {
    openDeleteProjectModal = true;
  }

  async function deleteProject() {
    try {
      await requestWithToken(
        `projects/${project.id}`,
        "DELETE",
        $authStore.token
      );
      pop();
    } catch (e) {
      console.error(e);
      error = "could not delete project";
    }
  }

  // -- editing history and updating on project change --
  let history: string[] = []; // the 'undo' history (as stringified jsons)
  $: canGoBackInHistory = history.length > 1; // if this is false the undo button is disabled

  function addCurrentProjectToHistory() {
    history = [...history, JSON.stringify(project)];
  }

  function goBackInHistory() {
    history.splice(history.length - 1, 1);
    history = history;
    project = JSON.parse(history[history.length - 1]);
    edited = true;
  }

  // make sure to not be able to leave page when project is edited
  let unsavedNavigateBackModalOn = false;

  let popConfirm = false;
  let onPopStateHasAlreadyBeenCalled = false;
  function onpopstateStay(e) {
    if (popConfirm) return;
    e.preventDefault();
    window.history.go(1);
    if (onPopStateHasAlreadyBeenCalled) {
      onPopStateHasAlreadyBeenCalled = false;
    } else {
      unsavedNavigateBackModalOn = true;
      onPopStateHasAlreadyBeenCalled = true;
    }
  }
  function leaveConfirm() {
    popConfirm = true;
    window.history.go(-1);
  }

  $: {
    if (edited) {
      window.onbeforeunload = () => true;
      window.onpopstate = onpopstateStay;
    } else {
      window.onbeforeunload = null;
      window.onpopstate = null;
    }
  }

  onDestroy(() => {
    window.onbeforeunload = null;
    window.onpopstate = null;
  });

  /// called if something has edited the project,
  /// to add to history and mark project as edited
  function hasEditedProject() {
    project = project;
    edited = true;
    addCurrentProjectToHistory();
  }

  // -- generell editing --
  function unselectQuestion() {
    selectedQuestion = undefined;
  }

  function editName(event) {
    project.name = event.detail;
    hasEditedProject();
  }

  function editDocumentation(event) {
    project.documentation_md = event.srcElement.value;
    hasEditedProject();
  }

  // -- task editing --
  function newTask() {
    const newTask: Task = {
      description: "description",
      is_voluntary: false,
      questions: [],
    };
    project.tasks.push(newTask);
    hasEditedProject();
    shouldScrollDown();
  }

  function onTaskDescriptionChange() {
    hasEditedProject();
  }

  function moveTask(taskNumber: number, up: boolean) {
    if (
      (up && taskNumber === 0) ||
      (!up && taskNumber === project!.tasks.length - 1)
    ) {
      return;
    }
    // move task
    const task = project.tasks[taskNumber];
    project.tasks.splice(taskNumber, 1);
    project.tasks.splice(taskNumber + (up ? -1 : 1), 0, task);
    project = project;
    edited = true;
  }

  function deleteTask(taskNumber: number) {
    project.tasks.splice(taskNumber, 1);
    project = project;
    edited = true;
  }

  // -- question editing --
  function newQuestion(taskNumber: number) {
    const newQuestion: Question = {
      question: "new question",
      type: "sql",
      solution: null,
    };
    project.tasks[taskNumber].questions.push(newQuestion);
    hasEditedProject();
    selectQuestion(taskNumber, project.tasks[taskNumber].questions.length - 1);
  }

  function deleteQuestion(taskNumber: number, questionNumber: number) {
    _deleteQuestion(taskNumber, questionNumber);
    project = project;
    edited = true;
    if (
      selectedQuestion?.taskNumber === taskNumber &&
      selectedQuestion?.questionNumber === questionNumber
    ) {
      selectedQuestion = undefined;
    }
  }

  function _deleteQuestion(
    taskNumber: number,
    questionNumber: number
  ): Question {
    return project.tasks[taskNumber].questions.splice(questionNumber, 1)[0];
  }

  function moveQuestion(
    taskNumber: number,
    questionNumber: number,
    up: boolean
  ) {
    if (up && questionNumber === 0) {
      if (taskNumber === 0) return;
      const movedQuestion = _deleteQuestion(taskNumber, questionNumber);
      project.tasks[taskNumber - 1].questions.push(movedQuestion);
    } else if (
      !up &&
      questionNumber === project.tasks[taskNumber].questions.length - 1
    ) {
      if (taskNumber === project.tasks.length - 1) return;
      const movedQuestion = _deleteQuestion(taskNumber, questionNumber);
      project.tasks[taskNumber + 1].questions.unshift(movedQuestion);
    } else {
      const movedQuestion = _deleteQuestion(taskNumber, questionNumber);
      const direction = up ? -1 : 1;
      project.tasks[taskNumber].questions.splice(
        questionNumber + direction,
        0,
        movedQuestion
      );
    }
    project = project;
    edited = true;
  }

  // -- selecting a question to be edited --
  let selectedQuestion:
    | {
        taskNumber: number;
        questionNumber: number;
        question: Question;
      }
    | undefined;

  function selectQuestion(taskNumber: number, questionNumber: number) {
    selectedQuestion = {
      taskNumber,
      questionNumber,
      question: project.tasks[taskNumber].questions[questionNumber],
    };
  }

  // -- tab index saved in localstorage
  let tabIndex: number;
  let tabIndexKey: string;

  $: if (project !== undefined) {
    initTabIndex();
  }

  // called after project is loaded
  function initTabIndex() {
    tabIndexKey = `project:${project.id}:tab_index`;
    const rawLocalStorageTabIndex = localStorage.getItem(tabIndexKey);
    let localStorageTabIndex: number;
    if (rawLocalStorageTabIndex !== null) {
      localStorageTabIndex = Number.parseInt(rawLocalStorageTabIndex);
    }
    tabIndex = localStorageTabIndex ?? 0;
  }

  $: if (tabIndex !== undefined) {
    localStorage.setItem(tabIndexKey, tabIndex.toString());
  }

  // -- database templates --
  let databaseTemplatePickerModalOpen = false;

  function databaseTemplateSelected(databaseTemplateSql: string) {
    project.sql = databaseTemplateSql;
    hasEditedProject();
  }
</script>

{#if error !== undefined}
  <p class="error">{error.toString()}</p>
{:else if loading}
  <!-- show first tab as loading -->
  <SkeletonText heading />
  <div class="spacer double" />
  <TabsSkeleton count={3} type="container" />
  <div class="spacer" />
  <TextAreaSkeleton rows={12} />
  <div class="spacer" />
  <TextInputSkeleton />
  <div class="spacer" />
  <ButtonSkeleton />
  <ButtonSkeleton />
{:else}
  <!-- undo button -->
  <Button
    style="float: right"
    icon={Undo20}
    size="small"
    kind="secondary"
    iconDescription="undo"
    disabled={!canGoBackInHistory}
    on:click={goBackInHistory}
  />
  <h2>
    {project.name}
  </h2>
  <div class="spacer" />
  <Tabs type="container" bind:selected={tabIndex}>
    <Tab>details</Tab>
    <Tab>database (optional)</Tab>
    <Tab>tasks</Tab>
    <svelte:fragment slot="content">
      <!-- first tab: editing name and documentation -->
      <TabContent style="padding: 0">
        <div class="page-overflow-scroll, dark-padded-tab-content">
          <TextInput
            value={project.name}
            disabled={!projectIsPrivate}
            light
            labelText="project title (max 50 characters)"
            placeholder="project title..."
            invalid={project.name.length === 0}
            invalidText="do not let the project name empty"
            maxlength={50}
            on:input={editName}
          />
          <div class="spacer" />
          <TextArea
            disabled={!projectIsPrivate}
            light
            value={project.documentation_md}
            labelText="project documentation"
            placeholder="project documentation..."
            rows={16}
            maxCount={10000}
            on:input={editDocumentation}
          />
        </div>
      </TabContent>
      <!-- TODO: implement -->
      <!-- second tab: editing and testing database/or no database (it is optional) -->
      <TabContent>
        {#if projectIsPrivate}
          <Toggle
            size="sm"
            toggled={project.sql !== null}
            on:toggle={() => {
              console.log("asdf");
              if (project.sql === null) {
                project.sql = "";
              } else {
                project.sql = null;
              }
              project = project;
              edited = true;
              addCurrentProjectToHistory();
            }}
            labelText="project should have a database"
            labelA="no database"
            labelB="database"
          />
          <div class="spacer" />
        {/if}
        {#if !projectIsPrivate && project.sql === null}
          <p>project has no database</p>
        {/if}
        {#if project.sql !== null}
          <SqlTextArea
            code={project.sql}
            disabled={!projectIsPrivate}
            on:change={(event) => {
              project.sql = event["detail"];
              edited = true;
              addCurrentProjectToHistory();
            }}
          />
          {#if projectIsPrivate}
            <Button
              size="small"
              on:click={() => (databaseTemplatePickerModalOpen = true)}
            >
              choose a database template
            </Button>
            <DatabaseTemplatePickerModal
              bind:open={databaseTemplatePickerModalOpen}
              onSelectTemplateSql={databaseTemplateSelected}
            />
          {/if}
        {/if}
      </TabContent>
      <!-- third tab: editing tasks -->
      <TabContent>
        <div id="question-edit-seperator">
          <div bind:this={tasksScrollNode} class="page-overflow-scroll">
            <!-- show all tasks in an ul -->
            <ul class:bx--tree={true} class:bx--tree--default={true}>
              {#each project.tasks as task, taskNumber (task)}
                <TaskComponent
                  {task}
                  {taskNumber}
                  onMove={(up) => {
                    moveTask(taskNumber, up);
                  }}
                  onDelete={() => {
                    deleteTask(taskNumber);
                  }}
                  onNewQuestion={() => {
                    newQuestion(taskNumber);
                  }}
                  onDescriptionChange={onTaskDescriptionChange}
                  editable={projectIsPrivate}
                >
                  <!-- show all questions in each task, also in an ul -->
                  <ul class:bx--tree={true} class:bx--tree--default={true}>
                    {#each task.questions as question, questionNumber (question)}
                      <QuestionComponent
                        {question}
                        {questionNumber}
                        onDelete={() =>
                          deleteQuestion(taskNumber, questionNumber)}
                        onMove={(up) =>
                          moveQuestion(taskNumber, questionNumber, up)}
                        on:click={() =>
                          selectQuestion(taskNumber, questionNumber)}
                        selected={selectedQuestion?.taskNumber === taskNumber &&
                          selectedQuestion?.questionNumber === questionNumber}
                        editable={projectIsPrivate}
                      />
                    {/each}
                  </ul>
                </TaskComponent>
              {/each}
              <!-- button to add new tasks located under all tasks -->
              <li class="bx--tree-node" style="padding-left: 0">
                <Button
                  disabled={!projectIsPrivate}
                  size="small"
                  kind="ghost"
                  on:click={newTask}
                  style="width: 100%; max-width: none; line-height:0; padding-left: 0.66rem; display: grid; grid-template-columns: auto 2px auto 1fr auto;"
                >
                  <Add24 style="display:block" />
                  <div />
                  <span id="add-task-button-label">add new task</span>
                  <div />
                  <div style="height: 38px" />
                </Button>
              </li>
            </ul>
          </div>
          <div />
          <!-- the right view to see the question currently being edited -->
          {#if selectedQuestion}
            <div class="page-overflow-scroll selected-question-box">
              <!-- close button to not view the current question (to unselect it) -->
              <Button
                tooltipPosition="left"
                tooltipAlignment="end"
                iconDescription="close"
                kind="ghost"
                size="small"
                icon={Close24}
                style="float: right"
                on:click={unselectQuestion}
              />
              <QuestionEditor
                editable={projectIsPrivate}
                taskNumber={selectedQuestion.taskNumber}
                questionNumber={selectedQuestion.questionNumber}
                question={selectedQuestion.question}
                onQuestionEdit={hasEditedProject}
              />
            </div>
          {:else}
            <!-- show if no question is currently being edited -->
            <div
              class="page-overflow-scroll, selected-question-box placeholder"
            >
              <center>
                no question selected,<br />
                click on a question to {projectIsPrivate ? "edit" : "view"} it
              </center>
            </div>
          {/if}
        </div>
      </TabContent>
    </svelte:fragment>
  </Tabs>
  <!-- modal to confirm deleting project -->
  <DeleteProjectModal
    bind:open={openDeleteProjectModal}
    projectName={project.name}
    on:submit={deleteProject}
  />
  <!-- modal to confirm leaving page, if the project is edited -->
  <UnsavedNavigateBackModal
    bind:open={unsavedNavigateBackModalOn}
    onLeave={leaveConfirm}
  />
  <!-- show a buttons to save and delete or info, depending on if the project is private (and therefore editable) -->
  {#if projectIsPrivate}
    <Button disabled={!edited} on:click={save} icon={Save20}>Save</Button>
    <Button kind="danger" icon={Delete20} on:click={pendingDeletingProject}
      >Delete</Button
    >
  {:else}
    <!--TODO: Clone projects-->
    <InlineNotification
      lowContrast
      kind="info-square"
      title="Not editable:"
      subtitle="Project not editable or deletable as not owned by you. Clone to
      edit project."
    />
  {/if}
{/if}

<style>
  p.error {
    color: red;
  }

  .dark-padded-tab-content {
    background-color: #262626;
    padding: 16px;
  }

  div.spacer {
    height: 16px;
  }

  div.spacer.double {
    height: 32px;
  }

  #add-task-button-label {
    float: left;
    text-align: left;
    font-size: 15px;
  }

  .selected-question-box {
    padding: 16px;
    background-color: #262626;
  }

  .selected-question-box.placeholder {
    height: 250px;
    padding: calc(125px - 16px);
    padding-left: 0;
    padding-right: 0;
    font-size: 16px;
    color: #4c4c4c;
  }

  #question-edit-seperator {
    display: grid;
    grid-template-columns: 1fr 3px 1fr;
  }

  .page-overflow-scroll {
    overflow-y: auto;
    max-height: calc(100vh - 296px);
  }
</style>
