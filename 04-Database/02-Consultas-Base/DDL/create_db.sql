-- DDL (Data Definition Language) statements for creating the database and tables
DROP DATABASE IF EXISTS `employees_db`; -- Drop the database if it already exists
CREATE DATABASE `employees_db`;         -- Create the (empty) database
USE `employees_db`;                     -- Make sure we use the database

-- Create departments table with the following columns:
-- `id`, `name`, `description`, `active`.
CREATE TABLE `departments` (
    `id` int NOT NULL AUTO_INCREMENT,
    `name` varchar(50) NOT NULL,
    `description` text NOT NULL,
    `active` boolean NOT NULL,
    PRIMARY KEY (`id`)
);


-- Create employees table with the following columns:
CREATE TABLE `employees` (
    `id` int NOT NULL AUTO_INCREMENT,
    `first_name` varchar(50) NOT NULL,
    `last_name` varchar(50) NOT NULL,
    `salary` decimal(10,2) NOT NULL,
    `active` boolean NOT NULL,
    `department_id` int NOT NULL,
    PRIMARY KEY (`id`),
    KEY `idx_employees_department_id` (`department_id`),
    CONSTRAINT `fk_employees_department_id` FOREIGN KEY (`department_id`) REFERENCES `departments` (`id`)
);