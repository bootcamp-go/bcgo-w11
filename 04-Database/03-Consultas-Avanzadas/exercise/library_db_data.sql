USE library;

-- DML: Data Manipulation Language
-- Insert data into authors
INSERT INTO authors (name, nationality) VALUES
('J.K. Rowling', 'British'),
('George Orwell', 'British'),
('Haruki Murakami', 'Japanese'),
('Mark A Garlick', 'British'),
('Oliver Barry', 'British'),
('Valerie Stimac', 'American'),
('Gabriel Garcia Marquez', 'Colombian');  -- non asocciated data

-- Insert data into books
INSERT INTO books (title, publisher, area) VALUES
('Harry Potter and the Sorcerer Stone', 'Bloomsbury', 'Fantasy'),
('1984', 'Secker & Warburg', 'Dystopian'),
('Kafka on the Shore', 'Shinchosha', 'Magical Realism'),
('The Wind-Up Bird Chronicle', 'Shinchosha', 'Magic Realism'),
('1Q84', 'Shinchosha', 'Magic Realism'),
('El Universo: Guía de viaje', 'Planeta', 'Science'),
('random book', 'random publisher', 'random area'); -- non asocciated data

-- Insert data into authors_books
INSERT INTO authors_books (author_id, book_id) VALUES
(1, 1),
(2, 2),
(3, 3),
(3, 4),
(3, 5),
(4, 6),
(5, 6),
(6, 6);

-- Insert data into students
INSERT INTO students (first_name, last_name, address, carreer, age) VALUES
('John', 'Doe', '123 Main St', 'Computer Science', 20),
('Jane', 'Doe', '123 Main St', 'Software Engineer', 21),
('Juan', 'Perez', '123 Calle Principal', 'Data Scientist', 22),
('Maria', 'Perez', '123 Calle Principal', 'Backend Programmer', 23),
('Filippo', 'Galli', '123 Via Principale', 'Frontend Programmer', 24);

-- Insert data into borrow_students
INSERT INTO borrow_students (book_id, student_id, borrow_date, return_date, returned) VALUES
(1, 1, '2021-01-01', '2021-07-16', true),
(1, 2, '2021-02-01', '2021-07-06', true),
(2, 5, '2022-02-01', '2022-05-03', false),
(3, 5, '2022-02-01', '2022-05-03', false),
(4, 5, '2022-02-01', '2022-05-03', false);