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

-- TODO: subtask removen, task kann beliebig kinder haben (task von task)
CREATE TABLE task
(
    project_id   INTEGER NOT NULL REFERENCES projects ON DELETE CASCADE,
    number       SMALLINT CHECK ( number > 0 ),
    name         VARCHAR NOT NULL,
    is_voluntary bool    NOT NULL DEFAULT FALSE,
    PRIMARY KEY (project_id, number)
);

CREATE TABLE subtask
(
    project_id  INTEGER NOT NULL REFERENCES projects ON DELETE CASCADE,
    task_number INTEGER REFERENCES task,
    number      SMALLINT CHECK ( number > 0 ),
    is_sql      BOOLEAN NOT NULL,
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
-- TODO: Link/QR-Code zum Beitritt
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

CREATE TABLE assignments
(
    project_id INTEGER           NOT NULL REFERENCES projects ON DELETE CASCADE,
    course_id  INTEGER           NOT NULL REFERENCES courses ON DELETE CASCADE,
    status     assignment_status NOT NULL,
    sequence   SMALLINT          NOT NULL CHECK ( sequence > 0),
    PRIMARY KEY (project_id, course_id)
);

CREATE TABLE answers
(
    id             INTEGER PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    course_id      INTEGER   NOT NULL REFERENCES courses ON DELETE CASCADE,
    participant_id INTEGER   NOT NULL REFERENCES participants ON DELETE CASCADE, -- TODO: Ohne constraints?
    project_id     INTEGER   NOT NULL REFERENCES projects ON DELETE CASCADE,
    task_number    SMALLINT  NOT NULL,
    subtask_number SMALLINT  NOT NULL,
    created_at     TIMESTAMP NOT NULL DEFAULT NOW(),
    FOREIGN KEY (project_id, task_number, subtask_number) REFERENCES subtask (project_id, task_number, number)
);