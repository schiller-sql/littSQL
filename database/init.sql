-- UTILS --
DROP SCHEMA IF EXISTS utils CASCADE;
CREATE SCHEMA utils;

-- TODO: Better method for random strings
CREATE OR REPLACE FUNCTION utils.random_string(length INTEGER) RETURNS CHAR AS
$$
DECLARE
    chars  TEXT[]  := '{0,1,2,3,4,5,6,7,8,9,A,B,C,D,E,F,G,H,I,J,K,L,M,N,O,P,Q,R,S,T,U,V,W,X,Y,Z,a,b,c,d,e,f,g,h,i,j,k,l,m,n,o,p,q,r,s,t,u,v,w,x,y,z}';
    result TEXT    := '';
    i      INTEGER := 0;
BEGIN
    IF length < 0 THEN
        RAISE EXCEPTION 'Given length cannot be less than 0';
    END IF;
    FOR i IN 1..length
        LOOP
            result := result || chars[1 + RANDOM() * (ARRAY_LENGTH(chars, 1) - 1)];
        END LOOP;
    RETURN result;
END;
$$ LANGUAGE plpgsql;


DROP SCHEMA IF EXISTS public CASCADE;
CREATE SCHEMA public;
-- TEACHER SIDE --
CREATE TABLE teachers -- represents a teacher
(
    id         INTEGER PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    first_name VARCHAR NOT NULL,
    last_name  VARCHAR NOT NULL,
    username   VARCHAR GENERATED ALWAYS AS ( first_name || '.'::VARCHAR || last_name ) STORED,
    password   VARCHAR CHECK (LENGTH(password) > 6)
);

CREATE TABLE databases -- represents a sample database that can be used in a project
(
    id           INTEGER PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    name         VARCHAR NOT NULL,
    data         bytea   NOT NULL,
    picture_path VARCHAR NOT NULL
);

-- TODO: Project groups
CREATE TABLE projects -- represents a template for a project a teacher can use as a assignment
(
    id               INTEGER PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    database_id      INTEGER REFERENCES databases,
    name             VARCHAR NOT NULL,
    documentation_md TEXT,
    owner            INTEGER REFERENCES teachers
);

-- TODO: question removen, task kann beliebig kinder haben (task von task)
-- TOOD: Trees?
CREATE TABLE task -- a group of questions, can be voluntary or not
(
    project_id   INTEGER  NOT NULL REFERENCES projects ON DELETE CASCADE,
    number       SMALLINT NOT NULL CHECK ( number > 0 ), -- position of the task (nr. 1, nr. 2, etc.)
    description  VARCHAR  NOT NULL,
    is_voluntary bool     NOT NULL DEFAULT FALSE,
    PRIMARY KEY (project_id, number)
);

CREATE TYPE question_type AS ENUM ( -- which type of question a question is
--    'multiple_choice',
--    'true/false',
--    'sql-without-question', -- a sql question but as a question you get a query output that you have to come to
    'sql',
    'text'
    );

CREATE TABLE questions -- a question asked to the participant, is part of a task and contains the solution
(
    project_id  INTEGER       NOT NULL REFERENCES projects ON DELETE CASCADE,
    task_number INTEGER       NOT NULL,
    number      SMALLINT      NOT NULL CHECK ( number > 0 ),
    type        question_type NOT NULL,
    solution    VARCHAR       NOT NULL,
    PRIMARY KEY (project_id, task_number, number),
    FOREIGN KEY (project_id, task_number) REFERENCES task (project_id, number) ON DELETE CASCADE
);

-- STUDENT SIDE --
-- TODO: mehrere Lehrer für einen Kurs (vlt)
CREATE TABLE courses -- a course with participants, belongs to a teacher
(
    id         INTEGER PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    teacher_id INTEGER NOT NULL REFERENCES teachers ON DELETE CASCADE,
    name       VARCHAR NOT NULL CHECK ( LENGTH(name) > 2 )
);

-- TODO: ein schüler mehrere Kurse (vlt)
-- TODO: Link/QR-Code zum Beitrittund
CREATE TABLE participants -- a participant of a course, a participant can't be in multiple courses
(
    id          INTEGER PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    course_id   INTEGER        NOT NULL REFERENCES courses ON DELETE CASCADE,
    name        VARCHAR,
    access_code CHAR(6) UNIQUE NOT NULL DEFAULT (utils.random_string(6))
);

CREATE TYPE assignment_status AS ENUM ( -- the status of one assignment of a course
    'finished',
    'open',
    'locked'
    );

CREATE TYPE assignment_solution_mode AS ENUM ( -- how the solutions should be shown to the participant
    'exam', -- no solutions, no query output, no submitting
    'tryout', -- on sql questions can see if the query wrong/the supposed query output is shown,
    -- after submitting, the solutions are shown, however the query can't be resubmitted
    'no-solutions-tryout', -- same as tryout, but after submitting the solutions are still not shown
    'voluntary' -- can always see if requested, no submitting
    );

-- TODO: TIMER FOR ASSIGNMENT
CREATE TABLE assignments -- a assignment given to the participant of a course, contains a project
(
    project_id    INTEGER                  NOT NULL REFERENCES projects ON DELETE CASCADE,
    course_id     INTEGER                  NOT NULL REFERENCES courses ON DELETE CASCADE,
    status        assignment_status        NOT NULL,
    solution_mode assignment_solution_mode NOT NULL,
    sequence      SMALLINT                 NOT NULL CHECK ( sequence > 0),
    PRIMARY KEY (project_id, course_id)
);

CREATE TYPE correct AS ENUM ( -- if a question has been answered correctly
    'unknown',
    'correct',
    'false'
    );

CREATE TABLE answers -- when and what a participant has answered to a question, including corrects
(
    id                         INTEGER PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    course_id                  INTEGER   NOT NULL REFERENCES courses ON DELETE CASCADE,
    participant_id             INTEGER   NOT NULL REFERENCES participants ON DELETE CASCADE, -- TODO: Ohne constraints?
    project_id                 INTEGER   NOT NULL REFERENCES projects ON DELETE CASCADE,
    task_number                SMALLINT  NOT NULL,
    question_number            SMALLINT  NOT NULL,
    created_at                 TIMESTAMP NOT NULL DEFAULT NOW(),
    answer                     VARCHAR   NOT NULL,
    is_correct_automatic       correct   NOT NULL DEFAULT 'unknown',
    is_correct_manual_approval correct   NOT NULL DEFAULT 'unknown',
    FOREIGN KEY (project_id, task_number, question_number) REFERENCES questions (project_id, task_number, number)
);