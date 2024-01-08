package main

type Ticket struct {
	Hours Day
}

func NewDay(milliseconds, seconds, minutes, hours int64) *Day {
	// validation

	return &Day{
		Milliseconds: milliseconds,
		Seconds:      seconds,
		Minutes:      minutes,
		Hours:        hours,
	}
}

type Day struct {
	Milliseconds int64
	Seconds      int64
	Minutes      int64
	Hours        int64
}