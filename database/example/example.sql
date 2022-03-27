-- insert example teachers (with the password '123456' in the default backend .env settings)
INSERT INTO teachers(email, password)
VALUES ('main@hello.world', '$2a$04$sA/YyRJDkxLg3jK6URBAdOxrRgKeDLTLgXcLsm9x7E8xtVQdDfQo6'),
       ('second@world.com', '$2a$04$sA/YyRJDkxLg3jK6URBAdOxrRgKeDLTLgXcLsm9x7E8xtVQdDfQo6');

-- insert public database
INSERT INTO databases(name, data, schema_svg_path)
VALUES ('public database', '', NULL);

-- insert non public database of different teachers
INSERT INTO databases(name, data, schema_svg_path, owner_id)
VALUES ('database of main', '', NULL, 1),
       ('database of second', '', NULL, 2);

-- insert public projects
INSERT INTO projects(database_id, name, documentation_md)
VALUES (1, 'public project 1', '# docs'),
       (NULL, 'public project 2', '# other docs');

-- insert non public projects
INSERT INTO projects (database_id, name, documentation_md, owner_id)
VALUES (2, 'project of main', '# docs', 1),
       (3, 'project of second (no tasks or questions)', '# other docs', 2);

-- insert tasks and questions for 'project of main'
INSERT INTO tasks(project_id, number, description, is_voluntary)
VALUES (3, 0, 'First task of project of main', FALSE),
       (3, 1, 'Second (voluntary) task of project of main', TRUE);

INSERT INTO questions(project_id, task_number, question, number, type, solution)
VALUES (3, 0, 'Why is this the first question of the first task?', 0, 'text', 'I don''t know...'),
       (3, 0, 'Why is this the second question of the first task?', 1, 'text', 'I don''t know...'),
       (3, 1, 'Why is this the first question of the first task?', 0, 'text', 'I don''t know...'),
       (3, 1, 'Why is this the second question of the first task?', 1, 'text', 'I don''t know...');

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
INSERT INTO assignments(project_id, course_id, status, solution_mode, number)
VALUES (1, 1, 'finished', 'tryout', 0),
       (3, 1, 'open', 'tryout', 1);

-- insert answer of 'student 1' into the second (and still open) assignment in the 'Main course'
INSERT INTO answers(course_id, project_id, participant_id, task_number, question_number, answer)
VALUES (1, 3, 1, 0, 0, 'Because it is that way!');
