package main

import (
	"fmt"
	"os"
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

func _GetFilePath() (string, string) {
	currentPath, err := os.Getwd()
	check(err)
	testFileName := "test_payday.yml"
	testYamlFile := fmt.Sprintf("%s/%s", currentPath, testFileName)
	return testFileName, testYamlFile
}

func _SetupSettings() {
	testYamlValues := "---\n\n" +
										"first_payday: 10\n" +
										"next_payday: TWO_WEEKS\n"
	testFileName, testYamlFile := _GetFilePath()
	os.WriteFile(testYamlFile, []byte(testYamlValues), 0777)
	os.Setenv("PAYDAY_SETTINGS", testFileName)
}

func _TeardownSettings() {
	_, testYamlFile := _GetFilePath()
	defer os.Unsetenv("PAYDAY_SETTINGS")
	os.Remove(testYamlFile)
}

func TestGetSettingsFileName(t *testing.T) {
	var firstExpected string = "payday.yml"
	var secondExpected string = "test.yml"

	firstActual := GetSettingsFileName()

	os.Setenv("PAYDAY_SETTINGS", secondExpected)
	secondActual := GetSettingsFileName()

	if firstActual != firstExpected {
		t.Error("Default value test failed.")
	}

	if secondActual != secondExpected {
		t.Error("Environment value test failed.")
	}

	defer os.Unsetenv("PAYDAY_SETTINGS")
}

func TestLoadSettings(t *testing.T) {
	_SetupSettings()
	actualFirstPayday, actualNextPaydayMarker := LoadSettings()
	if actualFirstPayday != 10 || actualNextPaydayMarker != "TWO_WEEKS" {
		t.Error("Expected settings did not load.")
	}
	_TeardownSettings()
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

func TestPaydayNextPaydayTwoWeeks(t *testing.T) {
	_SetupSettings()

	firstPayday := time.Date(2024, time.Month(3), 8, 0, 0, 0, 0, time.Local)
	firstActual := NextPayday(2024, 3, 1)
	if firstActual != firstPayday {
		t.Errorf("1. %s != %s", firstActual, firstPayday)
	}

	secondPayday := time.Date(2024, time.Month(3), 22, 0, 0, 0, 0, time.Local)
	secondActual := NextPayday(2024, 3, 10)
	if secondActual != secondPayday {
		t.Errorf("2. %s != %s", secondActual, secondPayday)
	}

	thirdPayday := time.Date(2024, time.Month(4), 10, 0, 0, 0, 0, time.Local)
	thirdActual := NextPayday(2024, 3, 22)
	if thirdActual != thirdPayday {
		t.Errorf("3. %s != %s", thirdActual, thirdPayday)
	}

	_TeardownSettings()
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
