-- DDL (Data Definition Language)
DROP DATABASE IF EXISTS `task_db`;

CREATE DATABASE `task_db`;

USE `task_db`;

-- Create table
CREATE TABLE `tasks` (
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `name` varchar(255) NOT NULL,
    `description` text NOT NULL,
    `completed` boolean NOT NULL DEFAULT FALSE,
    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_task_name` (`name`)
);