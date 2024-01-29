USE `tasks_complex_db`;

-- DML (Data Manipulation Language)
INSERT INTO `owners` (name, email) VALUES
('Owner 1', 'owner_1@something.com'),
('Owner 2', 'owner_2@something.com'),
('Owner 3', 'owner_3@something.com');

INSERT INTO `tasks` (name, description, completed, owner_id) VALUES
('Task 1', 'Description for task 1', FALSE, 1),
('Task 2', 'Description for task 2', FALSE, 1),
('Task 3', 'Description for task 3', TRUE, 2),
('Task 4', 'Description for task 4', TRUE, 2),
('Task 5', 'Description for task 5', FALSE, 3);