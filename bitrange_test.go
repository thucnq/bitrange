package bitrange_test

import (
	"fmt"
	"testing"

	"bitrange"
)

const (
	minuteBlockRange = 5
)

type Assigment struct {
	ID           int64
	Shifts       []bitrange.Range
	Weekdays     [7]bool
	HashedShifts [7]bitrange.BitRange
}

func TestNew(t *testing.T) {
	nwd := 3 // Monday
	as := []Assigment{
		{
			ID: 1,
			Shifts: []bitrange.Range{
				{"08:30", "12:00"},
				{"13:00", "18:00"},
			},
			Weekdays: [7]bool{false, true, false, true, false, true, false},
		},
		{
			ID: 2,
			Shifts: []bitrange.Range{
				{"18:00", "22:00"},
			},
			Weekdays: [7]bool{false, true, false, true, false, true, false},
		},
		{
			ID: 3,
			Shifts: []bitrange.Range{
				{"10:00", "22:00"},
			},
			Weekdays: [7]bool{true, false, true, false, true, false, true},
		},
		{
			ID: 4,
			Shifts: []bitrange.Range{
				{"17:00", "22:00"},
			},
			Weekdays: [7]bool{false, true, false, true, false, true, false},
		},
	}

	for idx := 0; idx < len(as); idx++ {
		for wIdx, weekday := range as[idx].Weekdays {
			if !weekday {
				continue
			}
			as[idx].HashedShifts[wIdx], _ = bitrange.GetBitRanges(
				as[idx].Shifts, minuteBlockRange,
			)
			fmt.Printf("%+v\n", as[idx].HashedShifts[wIdx])
		}
	}

	hss := make([]bitrange.BitRange, 0)
	for _, item := range as {
		fmt.Printf("%+v\n", item.HashedShifts)
		hss = append(hss, item.HashedShifts[nwd])
	}

	if bitrange.HaveOverlappedBit(hss) {
		panic("Duplicate")
	}

	fmt.Printf("OK")

}
