package main

// Func -> is like an interface
type PrepareFood func(ingridient string) bool

func PrepareChicken(ingridient string) bool {
	return true
}

func PrepareBeef(ingridient string) bool {
	return false
}

func main() {

}