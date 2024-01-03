package main

import (
	"fmt"
	"time"
)

type Breakfast struct {
	// ...
}

type Lunch struct {
	// ...
}

type Meal struct {
	// ...
}

type CookBreakfast func(ingridient string) Breakfast

type CookLunch func(ingridient string) Lunch

type CookMeal func(ingridient string) Meal

// Cook -> is like an interface
type Cook struct {
	BreakFast CookBreakfast
	Lunch     CookLunch
	Meal      CookMeal
}

// Interface
// - matters "what" you it does, not "how" it does
// - it's a contract
type Cooker interface {
	CookBreakfast(ingridient string) Breakfast
	CookLunch(ingridient string) Lunch
	CookMeal(ingridient string) Meal
}

// Implementing the interface
// Gordom Ramsay
type CookerRamsey struct {
	FirstName string
	LastName  string
	Speed     int	// seconds
}

func (c CookerRamsey) CookBreakfast(ingridient string) Breakfast {
	time.Sleep(time.Second * time.Duration(c.Speed))
	return Breakfast{}
}

func (c CookerRamsey) CookLunch(ingridient string) Lunch {
	time.Sleep(time.Second * time.Duration(c.Speed))
	return Lunch{}
}

func (c CookerRamsey) CookMeal(ingridient string) Meal {
	time.Sleep(time.Second * time.Duration(c.Speed))
	return Meal{}
}

type Apple struct {
	Cook Cooker
}

func main() {
	var cooker Cooker = CookerRamsey{
		FirstName: "Gordon",
		LastName:  "Ramsay",
		Speed:     5,
	}

	// i'm hungry, its early in the morning
	// i want to eat breakfast
	breakfast := cooker.CookBreakfast("eggs")
	// i want to eat lunch
	lunch := cooker.CookLunch("chicken")
	// i want to eat a meal
	meal := cooker.CookMeal("chicken")

	fmt.Println("yummy", breakfast, lunch, meal)
}