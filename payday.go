package main

import (
	"fmt"
	"time"
)

func LastDayOfMonth(y int, m int) int {
	// Every month has at least 28 days, so we can start counting there.
	var lastDay int = 28
	var today time.Time = time.Date(y, time.Month(m), lastDay, 0, 0, 0, 0, time.Local)
	var currentDate = today

	for today.Month() == currentDate.Month() {
		currentDate = today
		today = today.AddDate(0, 0, 1)
	}

	return int(currentDate.Day())
}

func IsWeekend(theDate time.Time) bool {
	// If a day is Sunday (0) or Saturday (6), it is a weekend day.
	var theWeekday int = int(theDate.Weekday())
	return (theWeekday == 0 || theWeekday == 6)
}

func FindPayday(theDate time.Time) time.Time {
	var paydate time.Time = theDate
	for IsWeekend(paydate) {
		paydate = paydate.AddDate(0, 0, -1)
	}
	return paydate
}

func NextPayday(y int, m int, d int) time.Time {
	var midPayday int = 15
	var lastPayday int = LastDayOfMonth(y, m)
	var paydate time.Time

	if d < midPayday {
		paydate = FindPayday(time.Date(y, time.Month(m), midPayday, 0, 0, 0, 0, time.Local))
	} else if d < lastPayday {
		paydate = FindPayday(time.Date(y, time.Month(m), lastPayday, 0, 0, 0, 0, time.Local))
	} else {
		var nextMonth time.Time = time.Date(y, time.Month(m), lastPayday, 0, 0, 0, 0, time.Local).AddDate(0, 0, 1)
		paydate = NextPayday(nextMonth.Year(), int(nextMonth.Month()), nextMonth.Day())
	}

	return paydate
}

func DateFormat(theDate time.Time) string {
	// There is a time.Format method, but I want the displayed format to be explcit.
	return fmt.Sprintf("%d-%02d-%02d", theDate.Year(), theDate.Month(), theDate.Day())
}

func main() {
	currentTime := time.Now().Local()
	nextPaydate := NextPayday(currentTime.Year(), int(currentTime.Month()), currentTime.Day())
	var hoursToPayday int = int(nextPaydate.Sub(currentTime).Hours())
	var daysToPayday int = hoursToPayday / 24
	var moduloHours int = hoursToPayday % 24

	fmt.Printf("The next payday is %s, which is a %s.\n", DateFormat(nextPaydate), nextPaydate.Weekday())
	fmt.Printf("That is %d days and %d hours away.\n", daysToPayday, moduloHours)
}
