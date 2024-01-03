package main

import "fmt"

type Microchip struct {
	Heatsink string
	Fan      string
}

func (m Microchip) On() {
	fmt.Println("Microchip is on")
}

type RAM struct {
	// declaring fields with its own types
	Brand string
	Size  int
	// Embedding
	Microchip
}

type Motherboard struct {
	Brand string
	// Embedding
	Microchip
}

// Pc represents a personal computer
type Pc struct {
	// Embedding
	RAM
	Motherboard
}

func main() {
	ram := RAM{
		Brand: "Kingston",
		Size:  16,
		Microchip: Microchip{
			Heatsink: "Aluminium",
			Fan:      "120mm",
		},
	}
	fmt.Printf("%#v\n", ram)

	ram.On()

	mobo := Motherboard{
		Brand: "Asus",
		Microchip: Microchip{
			Heatsink: "Copper",
			Fan:      "140mm",
		},
	}

	mobo.On()

	fmt.Printf("%#v\n", mobo)

	pc := Pc{
		RAM: ram,
		Motherboard: mobo,
	}

	pc.Motherboard.On()
	pc.RAM.On()
}