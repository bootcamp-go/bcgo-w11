USE movies_db;

-- This query will return the frequency of the movie avatar based on the favourite movie of the actors
SELECT mo.title, mo.rating, mo.awards, ac.first_name, ac.last_name
FROM movies mo INNER JOIN actors ac 
ON mo.id = ac.favorite_movie_id
WHERE mo.title = "Avatar";

-- This query will return the frequency of the movie avatar based on the favourite movie of the actors grouped by title
SELECT mo.title, mo.rating, mo.awards, ac.first_name, ac.last_name
FROM movies mo INNER JOIN actors ac 
ON mo.id = ac.favorite_movie_id
GROUP BY mo.title
HAVING mo.title = "Avatar";


-- This query will return the each movie with the number of actors that have it as their favourite movie
SELECT COUNT(*), mo.title, mo.rating, mo.awards, ac.first_name, ac.last_name
FROM movies mo INNER JOIN actors ac 
ON mo.id = ac.favorite_movie_id
GROUP BY mo.title, mo.rating, mo.awards;

-- This query will return the each movie with the number of actors that have it as their favourite movie
-- since ac.first_name and ac.last_name are not gonna be unique, creates a group for each row
-- therefore, the count will be 1
SELECT COUNT(*), mo.title, mo.rating, mo.awards, ac.first_name, ac.last_name
FROM movies mo INNER JOIN actors ac 
ON mo.id = ac.favorite_movie_id
GROUP BY mo.title, mo.rating, mo.awards, ac.first_name, ac.last_name;