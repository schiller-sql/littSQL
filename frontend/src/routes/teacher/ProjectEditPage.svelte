<script lang="ts">
  import { Button, InlineNotification } from "carbon-components-svelte";

  import { Save20, Delete20, Add24, Close24 } from "carbon-icons-svelte";

  import { afterUpdate, onMount } from "svelte";
  import TaskComponent from "../../components/Task.svelte";
  import QuestionComponent from "../../components/Question.svelte";
  import { authStore, fetchWithToken, requestWithToken } from "../../auth";
  import type Project from "../../types/Project";
  import type Question from "../../types/Question";
  import type Task from "../../types/Task";
  import QuestionEditor from "./QuestionEditor.svelte";
  import { letterFromNumber } from "../../util/utli";

  export let params: {
    projectId: number;
  };

  window.onkeyup = function (e) {};

  let shouldScrollOnNextTickDown = false;

  afterUpdate(() => {
    if (shouldScrollOnNextTickDown) {
      questionsScrollNode.scrollTop = questionsScrollNode.scrollHeight;
      shouldScrollOnNextTickDown = false;
    }
  });

  onMount(async () => {
    try {
      project = await fetchWithToken(
        `projects/${params.projectId}`,
        "get",
        $authStore.token
      );
      console.log(project);
    } catch (e) {
      error = e.toString();
    } finally {
      loading = false;
    }
  });

  let loading = true;
  let error: string | undefined;
  let project: Project | undefined;
  let edited = false;
  $: projectIsPrivate = !project?.is_public;

  async function save() {
    try {
      await requestWithToken(
        `projects/${params.projectId}`,
        "put",
        $authStore.token,
        project
      );
    } catch (e) {
      console.log(e);
      error = e;
    }
    edited = false;
  }

  function deleteTask(taskNumber: number) {
    project.tasks.splice(taskNumber, 1);
    project = project;
    edited = true;
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

  function _deleteQuestion(
    taskNumber: number,
    questionNumber: number
  ): Question {
    return project.tasks[taskNumber].questions.splice(questionNumber, 1)[0];
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

  function newQuestion(taskNumber: number) {
    const newQuestion: Question = {
      question: "new question",
      type: "sql",
      solution: null,
    };
    project.tasks[taskNumber].questions.push(newQuestion);
    project = project;
    edited = true;
    selectQuestion(taskNumber, project.tasks[taskNumber].questions.length - 1);
  }

  function hasEditedProject() {
    project = project;
    edited = true;
  }

  function newTask() {
    const newTask: Task = {
      description: "description",
      is_voluntary: false,
      questions: [],
    };
    project.tasks.push(newTask);
    project = project;
    edited = true;
    shouldScrollOnNextTickDown = true;
  }

  function onTaskDescriptionChange() {
    edited = true;
  }

  function selectQuestion(taskNumber: number, questionNumber: number) {
    selectedQuestion = {
      taskNumber,
      questionNumber,
      question: project.tasks[taskNumber].questions[questionNumber],
    };
  }

  function unselectQuestion() {
    selectedQuestion = undefined;
  }

  let selectedQuestion:
    | {
        taskNumber: number;
        questionNumber: number;
        question: Question;
      }
    | undefined;

  let questionsScrollNode;
</script>

<!-- TODO: edit history -->

{#if error !== undefined}
  <p style="color: red">{error.toString()}</p>
{:else}
  {#if project !== undefined}
    <!-- svelte-ignore missing-declaration -->
    <div class="separator">
      <div bind:this={questionsScrollNode} class="page-overflow-scroll">
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
              <ul class:bx--tree={true} class:bx--tree--default={true}>
                {#each task.questions as question, questionNumber (question)}
                  <QuestionComponent
                    {question}
                    {questionNumber}
                    onDelete={() => deleteQuestion(taskNumber, questionNumber)}
                    onMove={(up) =>
                      moveQuestion(taskNumber, questionNumber, up)}
                    on:click={() => selectQuestion(taskNumber, questionNumber)}
                    selected={selectedQuestion?.taskNumber === taskNumber &&
                      selectedQuestion?.questionNumber === questionNumber}
                    editable={projectIsPrivate}
                  />
                {/each}
              </ul>
            </TaskComponent>
          {/each}
          <li class="bx--tree-node" style="padding-left: 0">
            <Button
              disabled={!projectIsPrivate}
              size="small"
              kind="ghost"
              on:click={newTask}
              style="padding-left: 0.66rem; line-height: 0; display: grid; grid-template-columns: auto 2px auto 1fr auto; width: 100%; max-width:none; "
            >
              <Add24 style="display:block" />
              <div />
              <div style="display:grid grid-template-columns: auto 1fr">
                <span style="float:left; text-align:left; font-size:15px"
                  >add new task</span
                >
                <div />
              </div>
              <div />
              <div style="height: 38px" />
            </Button>
          </li>
        </ul>
      </div>
      <div />
      {#if selectedQuestion}
        <div
          style="background-color:#262626; display: absolute"
          class="page-overflow-scroll info-text edit-question-box"
        >
          <Button
            tooltipPosition="left"
            tooltipAlignment="end"
            iconDescription="close"
            kind="ghost"
            size="small"
            icon={Close24}
            style="float:right"
            on:click={unselectQuestion}
          />
          <div style="width: 100%; height: 100%">
            <QuestionEditor
              editable={projectIsPrivate}
              taskNumber={selectedQuestion.taskNumber}
              questionNumber={selectedQuestion.questionNumber}
              question={selectedQuestion.question}
              onQuestionEdit={hasEditedProject}
            />
          </div>
        </div>
      {:else}
        <div
          style="background-color:#262626; height: 500px"
          class="page-overflow-scroll"
        >
          <div
            class="info-text"
            style="margin-top: calc(250px - 16px); margin-bottom: calc(250px - 16px);"
          >
            <center>
              no question selected,<br />
              click on a question to {projectIsPrivate ? "edit" : "view"} it
            </center>
          </div>
        </div>
      {/if}
    </div>
  {/if}
  <div style="height: 1em" />
  {#if projectIsPrivate}
    <Button skeleton={loading} disabled={!edited} on:click={save} icon={Save20}>
      Save
    </Button>
    <Button kind="danger" skeleton={loading} icon={Delete20}>Delete</Button>
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
  .edit-question-box {
    padding: 16px;
  }

  .separator {
    display: grid;
    grid-template-columns: 1fr 3px 1fr;
  }

  .page-overflow-scroll {
    overflow-y: auto;
    max-height: calc(100vh - 174px);
  }

  .info-text {
    font-size: 16px;
    color: #4c4c4c;
  }
</style>
