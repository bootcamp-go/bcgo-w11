package main

func main() {
	i := 42

	// pass by value
	Increment(&i) // & takes the address of an instance -> 0xc0000b4008

	println(i)
}

func Increment(i *int) {
	// i -> 0xc0000b4008
	(*i)++
}