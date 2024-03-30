package main

import (
	"testing"
	"time"
)

type Day struct {
	Name    string
	Date    time.Time
	Weekend bool
}

type Month struct {
	Name    string
	Month   int
	LastDay int
	Year    int
}

type Format struct {
	S    string
	Date time.Time
}

func TestPaydayLastDayOfMonth(t *testing.T) {
	var months [3]Month
	months[0] = Month{"February", 2, 28, 2023}
	months[1] = Month{"February", 2, 29, 2024}
	months[2] = Month{"March", 3, 31, 2024}

	for _, m := range months {
		var actual int = LastDayOfMonth(m.Year, m.Month)
		if actual != m.LastDay {
			t.Errorf("%s %d has %d days, not %d.", m.Name, m.Year, m.LastDay, actual)
		}
	}
}

func TestPaydayIsWeekend(t *testing.T) {
	var days [3]Day
	days[0] = Day{"Saturday", time.Date(2024, time.Month(3), 23, 0, 0, 0, 0, time.Local), true}
	days[1] = Day{"Sunday", time.Date(2024, time.Month(3), 24, 0, 0, 0, 0, time.Local), true}
	days[2] = Day{"Monday", time.Date(2024, time.Month(3), 25, 0, 0, 0, 0, time.Local), false}

	for _, d := range days {
		var verb string

		if d.Weekend {
			verb = "is"
		} else {
			verb = "is not"
		}

		if IsWeekend(d.Date) != d.Weekend {
			t.Errorf("%s %s a weekend day.", d.Name, verb)
		}
	}
}

func TestPaydayNextPayday(t *testing.T) {
	firstPayday := time.Date(2024, time.Month(3), 5, 0, 0, 0, 0, time.Local)
	firstActual := NextPayday(2024, 3, 1)
	if firstActual != firstPayday {
		t.Errorf("1. %s != %s", firstActual, firstPayday)
	}

	secondPayday := time.Date(2024, time.Month(3), 19, 0, 0, 0, 0, time.Local)
	secondActual := NextPayday(2024, 3, 5)
	if NextPayday(2024, 3, 15) != secondPayday {
		t.Errorf("2. %s != %s", secondActual, secondPayday)
	}

	thirdPayday := time.Date(2024, time.Month(4), 5, 0, 0, 0, 0, time.Local)
	thirdActual := NextPayday(2024, 3, 19)
	if NextPayday(2024, 3, 31) != thirdPayday {
		t.Errorf("3. %s != %s", thirdActual, thirdPayday)
	}
}

func TestPaydayDateFormat(t *testing.T) {
	var formats [2]Format
	formats[0] = Format{"2024-01-01", time.Date(2024, time.Month(1), 1, 0, 0, 0, 0, time.Local)}
	formats[1] = Format{"2024-12-31", time.Date(2024, time.Month(12), 31, 0, 0, 0, 0, time.Local)}

	for _, f := range formats {
		var actual string = DateFormat(f.Date)
		var expected string = f.S
		if actual != expected {
			t.Errorf("%s != %s", actual, expected)
		}
	}
}
