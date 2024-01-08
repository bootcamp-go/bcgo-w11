package service

import (
	"app/desafio-cierre/internal/tickets/storage"
	"errors"
	"time"
)

var (
	// ErrInvalidPeriod is returned when the period is invalid
	ErrInvalidPeriod = errors.New("invalid period")
)

// Period is an alias for a period of time
type Period [2]time.Time
var (
	// EarlyMorning is the period from 00:00 to 06:00
	EarlyMorning = Period{time.Date(0, 0, 0, 0, 0, 0, 0, time.UTC), time.Date(0, 0, 0, 5, 59, 59, 999, time.UTC)}
	// Morning is the period from 06:00 to 12:00
	Morning = Period{time.Date(0, 0, 0, 6, 0, 0, 0, time.UTC), time.Date(0, 0, 0, 11, 59, 59, 999, time.UTC)}
	// Afternoon is the period from 12:00 to 20:00
	Afternoon = Period{time.Date(0, 0, 0, 12, 0, 0, 0, time.UTC), time.Date(0, 0, 0, 18, 59, 59, 999, time.UTC)}
	// Night is the period from 20:00 to 24:00
	Night = Period{time.Date(0, 0, 0, 19, 0, 0, 0, time.UTC), time.Date(0, 0, 0, 23, 59, 59, 999, time.UTC)}
)

// NewDefault returns a new Default instance
func NewDefault(st storage.Reader) *Default {
	return &Default{st: st}
}

// Default is an struct that represents a service for tickets
type Default struct {
	// st is the storage of tickets
	st storage.Reader
}

// GetTotal returns the total of tickets
func (d *Default) GetTotal() (total int, err error) {
	// get all the tickets
	t, err := d.st.Read()
	if err != nil {
		return
	}

	// return the total of tickets
	total = len(t)
	return
}

// GetCountByPeriod returns the count of tickets by period
func (d *Default) GetCountByPeriod(period Period) (count int, err error) {
	// verify the period
	if period != EarlyMorning && period != Morning && period != Afternoon && period != Night {
		err = ErrInvalidPeriod
		return
	}

	// get all the tickets
	t, err := d.st.Read()
	if err != nil {
		return
	}

	// iterate over the tickets
	for _, v := range t {
		// check if the departure is between the period
		if (v.Departure.Equal(period[0]) || v.Departure.After(period[0])) && (v.Departure.Equal(period[1]) || v.Departure.Before(period[1])) {
			// increment the count
			count++
		}
	}

	// return the count
	return
}

// GetPercentageByDestination returns the percentage of tickets by destination
func (d *Default) GetPercentageByDestination(destination string) (percentage float64, err error) {
	// get all the tickets
	t, err := d.st.Read()
	if err != nil {
		return
	}
	// check if there are tickets
	if len(t) == 0 {
		return
	}

	// iterate over the tickets
	for _, v := range t {
		// check if the destination is the same
		if v.Destination == destination {
			// increment the count
			percentage++
		}
	}

	// calculate the percentage
	percentage = percentage / float64(len(t))

	// return the percentage
	return
}