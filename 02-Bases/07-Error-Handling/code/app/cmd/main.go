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
			Rating: 9.2,
		},
		{
			Title: "The Dark Knight",
			Year: 2008,
			Rating: 9.0,
		},
		{
			Title: "Pulp Fiction",
			Year: 1994,
			Rating: 8.9,
		},
	}
	// save movies
	for _, mv := range movies {
		err := movieRepo.Save(&mv)
		if err != nil {
			switch {
				case errors.Is(err, repository.ErrLimitReached):
					fmt.Println("code 01: limit reached")
				case errors.Is(err, repository.ErrInvalidMovie):
					fmt.Println("code 02: invalid movie")
				case errors.Is(err, repository.ErrConstraintViolation):
					fmt.Println("code 03: constraint violation")
				default:
					fmt.Println("code 04: unknown error")
			}
			return
		}
		// err = errors.Unwrap(err)
		// err = errors.Unwrap(err)
		// err = errors.Unwrap(err)
		// err = errors.Unwrap(err)
		// err = errors.Unwrap(err)
		// err = errors.Unwrap(err)
		// err = errors.Unwrap(err)

		// if err != nil {
		// 	switch err {
		// 	case repository.ErrLimitReached:
		// 		fmt.Println("code 01: limit reached")
		// 	case repository.ErrInvalidMovie:
		// 		fmt.Println("code 02: invalid movie")
		// 	case repository.ErrConstraintViolation:
		// 		fmt.Println("code 03: constraint violation")
		// 	default:
		// 		fmt.Println("code 04: unknown error")
		// 	}
		// 	return
		// }

		fmt.Println(mv)
	}
}