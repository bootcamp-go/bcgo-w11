package average

func Average(grades ...uint) (result uint, err string) {
	// check if len is 0
	if len(grades) == 0 {
		err = "No grades provided"
		return
	}

	// calculate average
	var sum uint
	for _, grade := range grades {
		sum += grade
	}
	result = sum / uint(len(grades))
	return
}