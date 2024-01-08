package pets

// PetType is a type of pet
type PetType int
const (
	Dog PetType = iota
	Cat
	Hamster
	Tarantula
)

// PetGramsIntake is a function that receives the quantity of the pet and returns the grams of food
type PetGramsIntake func(int) float64

// Animal is a function that receives the quantity of dogs and returns the grams of food
func DogGramsIntake(quantity int) float64 {
	return float64(quantity) * 10000
}

// CatGramsIntake is a function that receives the quantity of cats and returns the grams of food
func CatGramsIntake(quantity int) float64 {
	return float64(quantity) * 5000
}

// HamsterGramsIntake is a function that receives the quantity of hamsters and returns the grams of food
func HamsterGramsIntake(quantity int) float64 {
	return float64(quantity) * 250
}

// TarantulaGramsIntake is a function that receives the quantity of tarantulas and returns the grams of food
func TarantulaGramsIntake(quantity int) float64 {
	return float64(quantity) * 150
}


// Orchestrator is a function that receives the quantity of dogs and returns the grams of food
func Orchestrator(pet PetType) (gr PetGramsIntake, err string) {
	switch pet {
	case Dog:
		gr = DogGramsIntake
	case Cat:
		gr = CatGramsIntake
	case Hamster:
		gr = HamsterGramsIntake
	case Tarantula:
		gr = TarantulaGramsIntake
	default:
		err = "Unknown pet"
	}
	return
}
