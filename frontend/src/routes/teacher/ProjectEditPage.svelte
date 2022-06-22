<script lang="ts">
  import { Button, InlineNotification } from "carbon-components-svelte";

  import {
    Save20,
    Delete20,
    Add24,
    TaskSettings16,
    Add20,
  } from "carbon-icons-svelte";

  import { onMount } from "svelte";
  import TaskComponent from "../../components/Task.svelte";
  import QuestionComponent from "../../components/Question.svelte";
  import { authStore, fetchWithToken, requestWithToken } from "../../auth";
  import type Project from "../../types/Project";
  import type Question from "../../types/Question";
  import type Task from "../../types/Task";

  export let params: {
    projectId: number;
  };

  window.onkeyup = function (e) {};

  onMount(async () => {
    try {
      project = await fetchWithToken(
        `projects/${params.projectId}`,
        "get",
        $authStore.token
      );
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
    )
      return;
    console.log(taskNumber, up);
    // move task
    const task = project.tasks[taskNumber];
    project.tasks.splice(taskNumber, 1);
    project.tasks.splice(taskNumber + (up ? -1 : -1), 0, task);
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
  }

  function newTask() {
    const newTask: Task = {
      description: "description",
      isVoluntary: false,
      questions: [],
    };
    project.tasks.push(newTask);
    project = project;
  }
</script>

<!-- TODO: edit history -->

{#if error !== undefined}
  <p style="color: red">{error.toString()}</p>
{:else}
  {#if project !== undefined}
    <!-- svelte-ignore missing-declaration -->
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
        >
          <ul class:bx--tree={true} class:bx--tree--default={true}>
            {#each task.questions as question, questionNumber (question)}
              <QuestionComponent
                {question}
                {questionNumber}
                onDelete={() => deleteQuestion(taskNumber, questionNumber)}
                onMove={(up) => moveQuestion(taskNumber, questionNumber, up)}
              />
            {/each}
          </ul>
        </TaskComponent>
      {/each}
      <li class="bx--tree-node" style="padding-left: 0">
        <Button
          size="small"
          kind="ghost"
          on:click={newTask}
          style="align-items:center; display:flex; padding-left: 10px"
        >
          <Add20 />
          <span style="font-size:smaller"> add new task</span>
        </Button>
      </li>
    </ul>
  {/if}
  <div style="height: 1em" />
  {#if projectIsPrivate}
    <Button skeleton={loading} disabled={!edited} on:click={save} icon={Save20}
      >Save</Button
    >
    <Button kind="danger" skeleton={loading} icon={Delete20}>Delete</Button>
  {:else}
    <InlineNotification
      lowContrast
      kind="info-square"
      title="Not editable"
      subtitle="this project is not editable or deletable as it not owned by you, clone it to edit the project"
    />
  {/if}
{/if}