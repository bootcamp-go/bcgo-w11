-- DDL (Data Definition Language)
DROP DATABASE IF EXISTS `tasks_complex_db`;

CREATE DATABASE `tasks_complex_db`;

USE `tasks_complex_db`;

-- Create owners table
CREATE TABLE `owners` (
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `name` varchar(255) NOT NULL,
    `email` varchar(255) NOT NULL,
    PRIMARY KEY (`id`)
);

-- Create tasks table
CREATE TABLE `tasks` (
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `name` varchar(255) NOT NULL,
    `description` text NOT NULL,
    `completed` boolean NOT NULL DEFAULT FALSE,
    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `owner_id` int(11) NOT NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_tasks_name` (`name`),
    KEY `idx_tasks_owner_id` (`owner_id`),
    CONSTRAINT `fk_tasks_owner_id` FOREIGN KEY (`owner_id`) REFERENCES `owners` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
);