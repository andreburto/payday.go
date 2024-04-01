package main

import (
	"fmt"
	"os"
	"time"

	"gopkg.in/yaml.v3"
)

type PaydaySettings struct {
	First int `yaml:"first_payday"`
	Next  string `yaml:"next_payday"`
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func GetSettingsFileName() string {
	filename, ok := os.LookupEnv("PAYDAY_SETTINGS")
	if !ok {
		filename = "payday.yml"
	}
	return filename
}

func LoadSettings() (int, string) {
	workingPath, err := os.Getwd()
	check(err)

	settingFilePath := fmt.Sprintf("%s/%s", workingPath, GetSettingsFileName())

	dat, err := os.ReadFile(settingFilePath)
	check(err)

	settings := PaydaySettings{}
	yaml.Unmarshal(dat, &settings)

	return settings.First, settings.Next
}

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
	var firstPayday int
	var nextPaydayMarker string 
	firstPayday, nextPaydayMarker = LoadSettings()

	var paydate time.Time
	if d < firstPayday {
		paydate = FindPayday(time.Date(y, time.Month(m), firstPayday, 0, 0, 0, 0, time.Local))
	} else {
		var nextPayday int

		switch nextPaydaySwitch := nextPaydayMarker; nextPaydaySwitch {
		case "TWO_WEEKS":
			nextPayday = firstPayday + 14
		case "LAST_DAY":
			nextPayday = LastDayOfMonth(y, m)
		default:
			panic("Next payday setting not found.")
		}

		paydate = FindPayday(time.Date(y, time.Month(m), nextPayday, 0, 0, 0, 0, time.Local))
		if d >= paydate.Day() {
			var nextMonth time.Time = time.Date(y, time.Month(m), 1, 0, 0, 0, 0, time.Local).AddDate(0, 1, 0)
			paydate = NextPayday(nextMonth.Year(), int(nextMonth.Month()), nextMonth.Day())
		}
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
