-- insert example teachers (with the password '123456' in the default backend .env settings)
INSERT INTO teachers(email, password)
VALUES ('main@hello.world', '$2a$04$sA/YyRJDkxLg3jK6URBAdOxrRgKeDLTLgXcLsm9x7E8xtVQdDfQo6'),
       ('second@world.com', '$2a$04$sA/YyRJDkxLg3jK6URBAdOxrRgKeDLTLgXcLsm9x7E8xtVQdDfQo6');

-- insert (public) database templates
INSERT INTO database_templates(name, description, sql)
VALUES ('cars', 'different car manufacturers', 'create table cars(name varchar);
insert into cars values (''mazda''), (''ferrari''), (''rolls royce'')'),
       ('programming languages', 'important programming languages',
        'create table programming_languages(name varchar);
insert into programming_languages ' ||
        'values (''assembly''), (''c++''),(''perl'')');

-- insert public projects
INSERT INTO projects(name, documentation_md, db_sql)
VALUES ('public project 1', '# docs', NULL),
       ('public project 2', '# other docs', (SELECT sql FROM database_templates WHERE id = 1));

-- insert non public projects
INSERT INTO projects(name, documentation_md, db_sql, owner_id)
VALUES ('private project 1', '# docs', (SELECT sql FROM database_templates WHERE id = 1), 1),
       ('private project 2', '# other docs', (SELECT sql FROM database_templates WHERE id = 1), 2);

-- insert tasks and questions for project 'public project 1'
INSERT INTO tasks(project_id, number, description, is_voluntary)
VALUES (1, 0, 'First task of project public project 1', FALSE),
       (1, 1, 'Second (voluntary) task of project public project 1', TRUE),
       (2, 0, 'First task of project public project 2', FALSE),
       (2, 1, 'Second (voluntary) task of project public project 2', TRUE);

INSERT INTO questions(project_id, task_number, question, number, type, solution)
VALUES
       (1, 0, 'Why is this the first question of the first task?', 0, 'text', 'I don''t know...'),
       (1, 0, 'Why is this the second question of the first task?', 1, 'text', 'I don''t know...'),
       (1, 1, 'Why is this the first question of the second task?', 0, 'text', 'I don''t know...'),
       (1, 1, 'Why is this the second question of the second task?', 1, 'text', 'I don''t know...'),
       (2, 0, 'Why is this the first question of the first task?', 0, 'text', 'I don''t know...'),
       (2, 0, 'Why is this the second question of the first task?', 1, 'text', 'I don''t know...'),
       (2, 1, 'Why is this the first question of the second task?', 0, 'text', 'I don''t know...'),
       (2, 1, 'Why is this the second question of the second task?', 1, 'text', 'I don''t know...');

-- insert tasks and questions for teacher 'main'
INSERT INTO tasks(project_id, number, description, is_voluntary)
VALUES (3, 0, 'First task of project of main', FALSE),
       (3, 1, 'Second (voluntary) task of project of main', TRUE);

INSERT INTO questions(project_id, task_number, question, number, type, solution)
VALUES (3, 0, 'Why is this the first question of the first task?', 0, 'text', 'I don''t know...'),
       (3, 0, 'Why is this the second question of the first task?', 1, 'text', 'I don''t know...'),
       (3, 1, 'Why is this the first question of the second task?', 0, 'text', 'I don''t know...'),
       (3, 1, 'Why is this the second question of the second task?', 1, 'text', 'I don''t know...');

-- insert course of 'main' teacher
INSERT INTO courses(teacher_id, name)
VALUES (1, 'Main course');

-- insert participants
INSERT INTO participants(course_id, name, access_code)
VALUES (1, 'student 1', '123456'),
       (1, 'student 2', 'ABCDEF');

-- insert participants with a random access code
INSERT INTO participants(course_id, name)
VALUES (1, 'student 3'),
       (1, 'student 4');

-- insert one public and one private project as assignments to 'Main course'
INSERT INTO assignments_data(name, comment, course_id, project_id, submission_date, manual_status, solution_mode)
VALUES ('first assigment', 'manually finished', 1, 3, null, 'finished', 'tryout'),
('second assigment', 'automatically finished', 1, 3, now(), 'open', 'tryout'),
('third assigment', 'still open', 1, 3, now() + interval '55 years', 'open', 'tryout');

-- insert answer of 'student 1' into the second (and still open) assignment in the 'Main course'
INSERT INTO answers(course_id, project_id, participant_id, task_number, question_number, answer)
VALUES (1, 3, 1, 0, 0, 'Because it is that way!');
