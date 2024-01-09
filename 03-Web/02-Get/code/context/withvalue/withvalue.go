package main

import (
	"context"
	"fmt"
)

type User struct {
	Name  string
	Other string
}

func main() {
	ctx := context.Background()
	auth(ctx)
}

// auth create a user and call functionOne with a derived context with value
func auth(ctx context.Context) {
	// Create a user
	user := User{
		Name:  "John Doe",
		Other: "Other data",
	}
	// Add the user to the context
	ctx = context.WithValue(ctx, "user", user)
	// Call functionOne with the context
	functionOne(ctx)
}

// functionOne add data to the user in the context and call functionTwo
func functionOne(ctx context.Context) {
	fmt.Println("Hola desde functionOne💖✨")

	// Add data to the user
	if v := ctx.Value("user"); v != nil {
		// - type assertion
		user, ok := v.(User)
		if ok {
			// - modify the user
			user.Other = "Data added in functionOne"
			// - add the user to the context
			ctx = context.WithValue(ctx, "user", user)
		}
	}
	// Call functionTwo with the context
	functionTwo(ctx)
}

// functionTwo print the user from the context
func functionTwo(ctx context.Context) {
	fmt.Println("Hola desde functionTwo✔️👍")

	// Create a default user for response
	contextUser := &User{
		Name:  "Anonymous",
		Other: "No data",
	}

	// Get the user from the context
	if v := ctx.Value("user"); v != nil {
		// - type assertion
		user, ok := v.(User)
		if ok {
			// - modify the user for response
			contextUser = &user
		}
	}

	// Print the user
	fmt.Printf("Hola %s, %s😎\n", contextUser.Name, contextUser.Other)
}
