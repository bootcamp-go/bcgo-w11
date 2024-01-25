USE movies_db;

-- DDL
/* 
    create a temporary table as a copy of an existing table
    - keeps the same structure (primary key, foreign key, indexes, etc.) but does not check integrity
*/
CREATE TEMPORARY TABLE `movies_temporary` LIKE `movies`;

-- DML
-- insert data into the temporary table
INSERT INTO `movies_temporary`
SELECT * FROM `movies`;

-- DDL
/*
    create a temporary table as a usual table from scratch
    - customizable
*/
CREATE TEMPORARY TABLE `the_walking_dead` (
    `id` int unsigned NOT NULL AUTO_INCREMENT,
    `serie_title` varchar(255) NOT NULL,
    `season_title` varchar(255) NOT NULL,
    `episode_title` varchar(255) NOT NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- DML
-- insert data into the temporary table
INSERT INTO `the_walking_dead` (`serie_title`, `season_title`, `episode_title`)
SELECT se.`title`, ss.`title`, ep.`title` FROM `series` se
INNER JOIN `seasons` ss ON ss.`series_id` = se.`id` AND se.`title` = 'The Walking Dead'
INNER JOIN `episodes` ep ON ep.`season_id` = ss.`id`;

SELECT twd.* FROM `the_walking_dead` twd WHERE twd.`episode_title` LIKE 'Days Gone Bye';

-- DDL
/*
    create temporary table as a copy of the table of a query
*/
CREATE TEMPORARY TABLE `movies_temporary_2` AS
SELECT m.* FROM `movies` m;