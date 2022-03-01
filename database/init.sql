-- UTILS --
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


-- TEACHER SIDE --
CREATE TABLE teacher
(
    id         INTEGER PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    first_name VARCHAR NOT NULL,
    last_name  VARCHAR NOT NULL,
    username   VARCHAR GENERATED ALWAYS AS ( CONCAT(first_name, '.', last_name) ) STORED,
    password   VARCHAR CHECK (LENGTH(password) > 6)
);

CREATE TABLE databases
(
    id   INTEGER PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    name VARCHAR NOT NULL,
    data bytea   NOT NULL
);

CREATE TABLE projects
(
    id          INTEGER PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    database_id INTEGER REFERENCES databases,
    name        VARCHAR NOT NULL
);

-- TODO: question removen, task kann beliebig kinder haben (task von task)
CREATE TABLE task
(
    project_id   INTEGER NOT NULL REFERENCES projects ON DELETE CASCADE,
    number       SMALLINT CHECK ( number > 0 ),
    name         VARCHAR NOT NULL,
    is_voluntary bool    NOT NULL DEFAULT FALSE,
    PRIMARY KEY (project_id, number)
);

CREATE TYPE question_type AS ENUM (
    --    'multiple_choice',
    --    'true/false',
    --    'sql-without-question', -- a sql question but as a question you get a query output that you have to come to
    'sql',
    'text'
    );

CREATE TABLE question
(
    project_id  INTEGER       NOT NULL REFERENCES projects ON DELETE CASCADE,
    task_number INTEGER REFERENCES task,
    number      SMALLINT CHECK ( number > 0 ),
    type        question_type NOT NULL,
    solution    VARCHAR       NOT NULL,
    PRIMARY KEY (project_id, number),
    FOREIGN KEY (project_id, task_number) REFERENCES task (project_id, number) ON DELETE CASCADE
);

-- STUDENT SIDE --
-- TODO: mehrere Lehrer für einen Kurs (vlt)
CREATE TABLE courses
(
    id         INTEGER PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    teacher_id INTEGER NOT NULL REFERENCES teacher ON DELETE CASCADE,
    name       VARCHAR NOT NULL CHECK ( LENGTH(name) > 2 )
);

-- TODO: ein schüler mehrere Kurse (vlt)
-- TODO: Link/QR-Code zum Beitrittund
CREATE TABLE participants
(
    id          INTEGER PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    course_id   INTEGER        NOT NULL REFERENCES courses ON DELETE CASCADE,
    name        VARCHAR,
    access_code CHAR(6) UNIQUE NOT NULL GENERATED ALWAYS AS ( utils.random_string(6) ) STORED
);

CREATE TYPE assignment_status AS ENUM (
    'finished',
    'open',
    'locked'
    );

CREATE TYPE assignment_solution_mode AS ENUM (
    'exam', -- no solutions, no query output, no submitting
    'tryout', -- on sql questions can see if the query wrong/the supposed query output is shown,
    -- after submitting, the solutions are shown, however the query can't be resubmitted
    'no-solutions-tryout', -- same as tryout, but after submitting the solutions are still not shown
    'voluntary' -- can always see if requested, no submitting
    );

-- TODO: TIMER FOR ASSIGNMENT
CREATE TABLE assignments
(
    project_id    INTEGER                  NOT NULL REFERENCES projects ON DELETE CASCADE,
    course_id     INTEGER                  NOT NULL REFERENCES courses ON DELETE CASCADE,
    status        assignment_status        NOT NULL,
    solution_mode assignment_solution_mode NOT NULL,
    sequence      SMALLINT                 NOT NULL CHECK ( sequence > 0),
    PRIMARY KEY (project_id, course_id)
);

CREATE TYPE correct AS ENUM (
    'unknown',
    'correct',
    'false'
    );

CREATE TABLE answers
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
    FOREIGN KEY (project_id, task_number, question_number) REFERENCES question (project_id, task_number, number)
);