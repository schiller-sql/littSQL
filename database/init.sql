-- UTILS --
DROP SCHEMA IF EXISTS utils CASCADE;
CREATE SCHEMA utils;

CREATE OR REPLACE FUNCTION utils.random_string(length INTEGER) RETURNS CHAR AS
$$
SELECT ARRAY_TO_STRING(
               ARRAY(SELECT SUBSTR('0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ', ((RANDOM() * (36 - 1) + 1)::INTEGER), 1)
                     FROM GENERATE_SERIES(1, length)), '')
$$ LANGUAGE sql;


DROP SCHEMA IF EXISTS public CASCADE;
CREATE SCHEMA public;
-- TEACHER SIDE --
-- represents a teacher
CREATE TABLE teachers
(
    id       INTEGER PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    email    VARCHAR UNIQUE NOT NULL,
    password VARCHAR        NOT NULL
);

-- represents a sample database that can be used in a project, an existing database_template should not be changed
CREATE TABLE database_templates
(
    id          INTEGER PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    name        VARCHAR NOT NULL,
    description TEXT    NOT NULL,
    sql         TEXT    NOT NULL
);

-- represents a template for a project a teacher can use as an assignment
CREATE TABLE projects
(
    id               INTEGER PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    name             VARCHAR NOT NULL CHECK (LENGTH(name) > 0),
    documentation_md TEXT    NOT NULL DEFAULT '',
    db_sql           TEXT,
    owner_id         INTEGER REFERENCES teachers
);

CREATE TABLE cached_projects_sql_data
(
    project_id INTEGER PRIMARY KEY REFERENCES projects,
    data       bytea NOT NULL
);

-- a group of questions, can be voluntary or not
CREATE TABLE tasks
(
    project_id   INTEGER  NOT NULL REFERENCES projects ON DELETE CASCADE,
    number       SMALLINT NOT NULL CHECK ( number >= 0 ), -- position of the task (#1, #2, etc.)
    description  TEXT     NOT NULL,
    is_voluntary bool     NOT NULL DEFAULT FALSE,
    PRIMARY KEY (project_id, number)
);

-- type of question
CREATE TYPE question_type AS ENUM (
    --    'multiple_choice',
--    'true/false',
--    'sql-without-question', -- a sql question but as a question you get a query output that you have to come to
    'sql',
    'text'
    );

-- TODO: maybe add (optional) hints for questions
-- a question asked to the participant, is part of a task and contains the solution
CREATE TABLE questions
(
    project_id  INTEGER       NOT NULL REFERENCES projects ON DELETE CASCADE,
    task_number INTEGER       NOT NULL,
    question    VARCHAR       NOT NULL,
    number      SMALLINT      NOT NULL CHECK ( number >= 0 ),
    type        question_type NOT NULL,
    solution    VARCHAR,
    PRIMARY KEY (project_id, task_number, number),
    FOREIGN KEY (project_id, task_number) REFERENCES tasks (project_id, number) ON DELETE CASCADE
);

-- STUDENT SIDE --
-- TODO: more than one teacher for the same course
-- a course with participants, belongs to a teacher
CREATE TABLE courses
(
    id         INTEGER PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    teacher_id INTEGER NOT NULL REFERENCES teachers ON DELETE CASCADE,
    name       VARCHAR NOT NULL CHECK ( LENGTH(name) > 0 )
);

-- TODO: allow participants to join multiple courses
-- TODO: link/QR-code for joinig
-- a participant of a course, a participant can't be in multiple courses (by now)
CREATE TABLE participants
(
    id          INTEGER PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    course_id   INTEGER        NOT NULL REFERENCES courses ON DELETE CASCADE,
    name        VARCHAR,
    access_code CHAR(6) UNIQUE NOT NULL DEFAULT (utils.random_string(6))
);

CREATE TYPE correction_behaviour AS ENUM (
    'show_no_correction',
    'show_correction',
    'show_correction_and_solution'
    );

-- an assignment given to the participant of a course, contains a project
CREATE TABLE assignments
(
    id                                      INTEGER PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    number                                  SMALLINT             NOT NULL CHECK ( number >= 0),
    name                                    VARCHAR              NOT NULL CHECK (LENGTH(name) > 1),
    comment                                 VARCHAR,
    course_id                               INTEGER              NOT NULL REFERENCES courses ON DELETE CASCADE,
    project_id                              INTEGER              REFERENCES projects ON DELETE SET NULL,
    -- after the timestamp, if it is not null, students can no longer submit,
    -- however the project is not locked,
    -- the questions can still be viewed,
    -- and depending on show_solution_after_finished, also the solution.
    -- to manually make the project this way, simply set submission_date to true
    finished_date                           TIMESTAMP,
    -- sql questions, with a valid answer, can be auto corrected,
    -- by comparing output of the student and solution query
    enable_auto_correction_on_sql_questions BOOLEAN              NOT NULL DEFAULT TRUE,
    -- always (however before finished) show the result of the solution query as a table, provided there is one,
    -- only works for sql questions with (valid) solutions
    show_query_solution                     BOOLEAN              NOT NULL DEFAULT TRUE,
    -- after a student has submitted, the student cannot do it again,
    -- the student can also only read the answer after submitting, if this is actually possible then,
    -- is determined by
    submit_only_once                        BOOLEAN              NOT NULL DEFAULT TRUE,
    -- before the project is finished, if the student should see correction and solution, only correction or nothing,
    -- seeing the correction also means seeing the correction comment, if one exists.
    -- if submit_only_once is active solution and/or correction is always shown after the first and only submission
    active_correction_behaviour             correction_behaviour NOT NULL DEFAULT 'show_correction',
    -- after the project is finished, if the student should see correction and solution, only correction or nothing,
    -- seeing the correction also means seeing the correction comment, if one exists.
    finished_correction_behavior            correction_behaviour NOT NULL DEFAULT 'show_correction_and_solution',
    -- after the project is finished, if the student should not be able to see their own answer
    finished_hide_answers                   BOOLEAN              NOT NULL DEFAULT FALSE,
    -- if the solution should be shown after the finished_date has arrived (provided there is a finished_date)
    -- if locked, nothing can be viewed exception the assignment name by the student, submission is not possible
    locked                                  BOOLEAN              NOT NULL DEFAULT TRUE
);

CREATE VIEW assignments_listing AS
SELECT id,
       number,
       name,
       course_id,
       (CASE
            WHEN locked THEN 'locked'
            ELSE (CASE
                      WHEN finished_date IS NULL OR finished_date > NOW() THEN 'open'
                      ELSE 'finished'
                END)
           END) AS status
FROM assignments;

-- if a question has been answered correctly
CREATE TYPE correct AS ENUM (
    'unknown',
    'correct',
    'false'
    );

-- TODO: possibly use two different tables for answer and correction
-- time and content the participant has answered to a question, including corrects
CREATE TABLE answers
(
    id                         INTEGER PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    course_id                  INTEGER   NOT NULL REFERENCES courses ON DELETE CASCADE,
    participant_id             INTEGER   NOT NULL REFERENCES participants ON DELETE CASCADE,
    project_id                 INTEGER   NOT NULL REFERENCES projects ON DELETE CASCADE,
    task_number                SMALLINT  NOT NULL,
    question_number            SMALLINT  NOT NULL,
    last_answer_update         TIMESTAMP NOT NULL DEFAULT NOW(),
    answer                     VARCHAR   NOT NULL,
    is_correct_automatic       correct   NOT NULL DEFAULT 'unknown',
    is_correct_manual_approval correct   NOT NULL DEFAULT 'unknown',
    correction_comment         VARCHAR,
    FOREIGN KEY (project_id, task_number, question_number) REFERENCES questions (project_id, task_number, number) ON DELETE CASCADE
);
