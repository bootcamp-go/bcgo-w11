/* Ejercicio 2

	En la base de datos “movies”, seleccionar una tabla donde crear un índice y luego chequear la creación del mismo.
	Analizar por qué crearía un índice en la tabla indicada y con qué criterio se elige/n el/los campos.
*/

USE movies_db;

-- DDL
/*
CREATE TEMPORARY TABLE movies_temporary (
	`id` int unsigned NOT NULL AUTO_INCREMENT,
    `title` varchar(50) NOT NULL,
    PRIMARY KEY (`id`),
    KEY `idx_movies_temporary_title` (`title`)
);
*/
DROP TABLE IF EXISTS movies_temporary;
CREATE TEMPORARY TABLE movies_temporary LIKE movies;
ALTER TABLE movies_temporary ADD INDEX `idx_movies_temporary_title` (`title`);
ALTER TABLE movies_temporary ADD INDEX `idx_movies_temporary_rating` (`rating`);

INSERT INTO movies_temporary SELECT * FROM movies;

-- DQL
SELECT * FROM movies_temporary LIMIT 4;

SELECT m.title, m.rating FROM movies_temporary m WHERE m.title = "Avatar";

SELECT m.title, m.rating FROM movies_temporary m WHERE m.rating BETWEEN 4.0 AND 7.0;

-- Query with indexed field and non-indexed field
SELECT m.title, m.rating FROM movies_temporary m
WHERE m.title = "Titanic"
AND m.awards > 1;

SELECT m.title, m.rating FROM movies_temporary m
WHERE m.title = "Titanic"
OR m.awards > 1;