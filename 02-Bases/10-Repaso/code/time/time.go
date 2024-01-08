package main

import (
	"fmt"
	"time"
)

const CustomDate = "2006/01-02"

func main() {
	timeString := "2022/01-01"
	// parse time
	time2022, err := time.Parse(CustomDate, timeString)
	if err != nil {
		println(err.Error())
		return
	}

	// declare times
	time2023 := time.Date(2023, time.January, 1, 0, 0, 0, 0, time.UTC)
	time2024 := time.Date(2024, time.January, 1, 0, 0, 0, 0, time.UTC)

	fmt.Println(time2022.Before(time2023))
	fmt.Println(time2024.After(time2023))

	// parse time to string
	fmt.Println(time2022.Format(time.DateOnly))
}
