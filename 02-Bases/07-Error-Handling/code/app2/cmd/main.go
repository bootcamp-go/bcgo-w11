package main

import (
	"app/internal"
	"app/internal/repository"
	"errors"
	"fmt"
)

func main() {
	// dependencies
	movieRepo := repository.NewMovieMap(3)

	// movies
	movies := []internal.Movie{
		{
			Title: "The Matrix",
			Year: 1999,
			Rating: 8.7,
		},
		{
			Title: "The Godfather",
			Year: 1972,
			Rating: 9.6,
		},
		{
			Title: "The Lord of the Rings: The Return of the King",
			Year: 2003,
			Rating: 8.9,
		},
		{
			Title: "The Dark Knight",
			Year: 2008,
			Rating: 9.0,
		},
	}
	// save movies
	for _, mv := range movies {
		err := movieRepo.Save(&mv)

		// unwrap error
		// err = errors.Unwrap(err)
		if err != nil {
			var errLimitReached *repository.ErrorMovieLimit
			if errors.As(err, &errLimitReached) {
				fmt.Println("code 02: limit reached", errLimitReached.Limit)
				return
			}
			// errLimitReached, ok := err.(*repository.ErrorMovieLimit)
			// if ok {
			// 	fmt.Println("code 02: limit reached", errLimitReached.Limit)
			// 	return
			// }

			var errInvalidMovie *repository.ErrorMovieInvalid
			if errors.As(err, &errInvalidMovie) {
				fmt.Println("code 01: invalid movie", errInvalidMovie.Feature)
				return
			}
			// errInvalidMovie, ok := err.(*repository.ErrorMovieInvalid)
			// if ok {
			// 	fmt.Println("code 01: invalid movie", errInvalidMovie.Feature)
			// 	return
			// }

			fmt.Println("code 03: unknown error")
		}
	}


		// if err != nil {
		// 	switch err {
		// 	case repository.ErrLimitReached:
		// 		fmt.Println("code 01: limit reached")
		// 	case repository.ErrInvalidMovieTitle:
		// 		fmt.Println("code 02: invalid movie")
		// 	case repository.ErrConstraintViolation:
		// 		fmt.Println("code 03: constraint violation")
		// 	default:
		// 		fmt.Println("code 04: unknown error")
		// 	}
		// 	return
		// }
}