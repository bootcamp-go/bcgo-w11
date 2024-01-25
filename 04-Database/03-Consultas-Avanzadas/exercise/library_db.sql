-- DDL
DROP DATABASE IF EXISTS library;
CREATE DATABASE library;
USE library;

-- Table Authors
CREATE TABLE authors (
    `id` int AUTO_INCREMENT,
    `name` varchar(100) NOT NULL,
    `nationality` varchar(100) NOT NULL,
    PRIMARY KEY (`id`)
);

-- Table Books
CREATE TABLE books (
    `id` int AUTO_INCREMENT,
    `title` varchar(100) NOT NULL,
    `publisher` varchar(100) NOT NULL,
    `area` varchar(100) NOT NULL,
    PRIMARY KEY (`id`)
);

-- Table Students
CREATE TABLE students (
    `id` int AUTO_INCREMENT,
    `first_name` varchar(100) NOT NULL,
    `last_name` varchar(100) NOT NULL,
    `address` varchar(100) NOT NULL,
    `carreer` varchar(100) NOT NULL,
    `age` int NOT NULL,
    PRIMARY KEY (`id`)
);


-- Intermediate table
-- Table: authors_books
CREATE TABLE authors_books (
    `author_id` int NOT NULL,
    `book_id` int NOT NULL,
    PRIMARY KEY (`author_id`,`book_id`),  -- composite primary key: unique combination of both columns
    KEY `idx_authors_books_author_id` (`author_id`),
    KEY `idx_authors_books_book_id` (`book_id`),
    CONSTRAINT `fk_authors_books_author_id` FOREIGN KEY (`author_id`) REFERENCES `authors` (`id`),
    CONSTRAINT `fk_authors_books_book_id` FOREIGN KEY (`book_id`) REFERENCES `books` (`id`)
);

-- Table: borrow_students
CREATE TABLE borrow_students (
    `id` int AUTO_INCREMENT,    -- not composite primary key: unique id for each row, can be repeated combination of book_id and student_id
    `book_id` int NOT NULL,
    `student_id` int NOT NULL,
    `borrow_date` date NOT NULL,
    `return_date` date NOT NULL,
    `returned` boolean NOT NULL,
    PRIMARY KEY (`id`),
    KEY `idx_books_students_book_id` (`book_id`),
    KEY `idx_books_students_student_id` (`student_id`),
    CONSTRAINT `fk_books_students_book_id` FOREIGN KEY (`book_id`) REFERENCES `books` (`id`),
    CONSTRAINT `fk_books_students_student_id` FOREIGN KEY (`student_id`) REFERENCES `students` (`id`)
);