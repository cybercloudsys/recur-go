// Recur
// Copyright © 2023 Cyber Cloud Systems LLC

// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as
// published by the Free Software Foundation, either version 3 of the
// License, or (at your option) any later version.

// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.

// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package tests

import (
	"testing"
	"time"

	"cybercloudsys.com/recur"
)

func TestFirstWorkingDay(t *testing.T) {
	recurPattern := recur.Monthly().OnDay(1).OnWorkday().Build()
	recurPattern.Start = time.Date(2101, 1, 1, 0, 0, 0, 0, time.UTC)
	expected := recurPattern.Start.AddDate(0, 0, 2)
	if nextTime := recurPattern.NextTime(); !nextTime.Equal(expected) {
		t.Errorf("Expected %s, but got %s (Sat, first working day '1')", expected, nextTime)
	}
	recurPattern.Start = time.Date(2101, 5, 1, 0, 0, 0, 0, time.UTC)
	expected = recurPattern.Start.AddDate(0, 0, 1)
	if nextTime := recurPattern.NextTime(); !nextTime.Equal(expected) {
		t.Errorf("Expected %s, but got %s (Sun, first working day '2')", expected, nextTime)
	}
	recurPattern.Start = time.Date(2101, 8, 10, 0, 0, 0, 0, time.UTC)
	expected = time.Date(2101, 9, 1, 0, 0, 0, 0, time.UTC)
	if nextTime := recurPattern.NextTime(); !nextTime.Equal(expected) {
		t.Errorf("Expected %s, but got %s (Mon, first working day '3')", expected, nextTime)
	}
	recurPattern.Start = time.Date(2100, 12, 2, 0, 0, 0, 0, time.UTC)
	expected = time.Date(2101, 1, 3, 0, 0, 0, 0, time.UTC)
	if nextTime := recurPattern.NextTime(); !nextTime.Equal(expected) {
		t.Errorf("Expected %s, but got %s (Sat, first working day '4')", expected, nextTime)
	}
	recurPattern.Start = time.Date(2101, 4, 2, 0, 0, 0, 0, time.UTC)
	expected = time.Date(2101, 5, 2, 0, 0, 0, 0, time.UTC)
	if nextTime := recurPattern.NextTime(); !nextTime.Equal(expected) {
		t.Errorf("Expected %s, but got %s (Sun, first working day '5')", expected, nextTime)
	}
	recurPattern.Start = time.Date(2101, 7, 2, 0, 0, 0, 0, time.UTC)
	expected = time.Date(2101, 8, 1, 0, 0, 0, 0, time.UTC)
	if nextTime := recurPattern.NextTime(); !nextTime.Equal(expected) {
		t.Errorf("Expected %s, but got %s (Mon, first working day '6')", expected, nextTime)
	}
	expectedString := "Monthly, 1st day (Mon-Fri)"
	if patternString := recurPattern.String(); patternString != expectedString {
		t.Errorf("Expected '%s', but got '%s' (first working day monthly '7')", expectedString, patternString)
	}
}

func TestSecondWorkingDay(t *testing.T) {
	recurPattern := recur.Monthly().OnDay(2).OnWorkday().Build()
	recurPattern.Start = time.Date(2101, 6, 20, 0, 0, 0, 0, time.UTC)
	expected := time.Date(2101, 7, 1, 0, 0, 0, 0, time.UTC)
	if nextTime := recurPattern.NextTime(); !nextTime.Equal(expected) {
		t.Errorf("Expected %s, but got %s (Sat, 2nd working day '1')", expected, nextTime)
	}
	recurPattern.Start = time.Date(2101, 9, 25, 0, 0, 0, 0, time.UTC)
	expected = time.Date(2101, 10, 3, 0, 0, 0, 0, time.UTC)
	if nextTime := recurPattern.NextTime(); !nextTime.Equal(expected) {
		t.Errorf("Expected %s, but got %s (Sun, 2nd working day '2')", expected, nextTime)
	}
	expectedString := "Monthly, 2nd day (Mon-Fri)"
	if patternString := recurPattern.String(); patternString != expectedString {
		t.Errorf("Expected '%s', but got '%s' (2nd working day monthly '2')", expectedString, patternString)
	}
}

func TestThirteenthWorkingDay(t *testing.T) {
	recurPattern := recur.Monthly().OnDay(13).OnWorkday().Build()
	recurPattern.Start = time.Date(2101, 7, 20, 0, 0, 0, 0, time.UTC)
	expected := time.Date(2101, 8, 12, 0, 0, 0, 0, time.UTC)
	if nextTime := recurPattern.NextTime(); !nextTime.Equal(expected) {
		t.Errorf("Expected %s, but got %s (Sat, 13th working day '1')", expected, nextTime)
	}
	recurPattern.Start = time.Date(2101, 2, 25, 0, 0, 0, 0, time.UTC)
	expected = time.Date(2101, 3, 14, 0, 0, 0, 0, time.UTC)
	if nextTime := recurPattern.NextTime(); !nextTime.Equal(expected) {
		t.Errorf("Expected %s, but got %s (Sun, 13th working day '2')", expected, nextTime)
	}
	expectedString := "Monthly, 13th day (Mon-Fri)"
	if patternString := recurPattern.String(); patternString != expectedString {
		t.Errorf("Expected '%s', but got '%s' (13th working day monthly '2')", expectedString, patternString)
	}
}

func TestLastWorkingDay(t *testing.T) {
	recurPattern := recur.Monthly().FromLastDay().OnWorkday().Build()
	recurPattern.Start = time.Date(2101, 4, 1, 0, 0, 0, 0, time.UTC)
	expected := time.Date(2101, 4, 29, 0, 0, 0, 0, time.UTC)
	if nextTime := recurPattern.NextTime(); !nextTime.Equal(expected) {
		t.Errorf("Expected %s, but got %s (Sat, last working day '1')", expected, nextTime)
	}
	recurPattern.Start = time.Date(2101, 7, 1, 0, 0, 0, 0, time.UTC)
	expected = time.Date(2101, 7, 29, 0, 0, 0, 0, time.UTC)
	if nextTime := recurPattern.NextTime(); !nextTime.Equal(expected) {
		t.Errorf("Expected %s, but got %s (Sun, last working day '2')", expected, nextTime)
	}
	recurPattern.Start = time.Date(2101, 10, 1, 0, 0, 0, 0, time.UTC)
	expected = time.Date(2101, 10, 31, 0, 0, 0, 0, time.UTC)
	if nextTime := recurPattern.NextTime(); !nextTime.Equal(expected) {
		t.Errorf("Expected %s, but got %s (Mon, last working day '3')", expected, nextTime)
	}
	expectedString := "Monthly, last day (Mon-Fri)"
	if patternString := recurPattern.String(); patternString != expectedString {
		t.Errorf("Expected '%s', but got '%s' (last working day '4')", expectedString, patternString)
	}
}

func TestLastDay(t *testing.T) {
	recurPattern := recur.Monthly().FromLastDay().Build()
	recurPattern.Start = time.Date(2101, 1, 1, 0, 0, 0, 0, time.UTC)
	expected := time.Date(2101, 1, 31, 0, 0, 0, 0, time.UTC)
	if nextTime := recurPattern.NextTime(); !nextTime.Equal(expected) {
		t.Errorf("Expected %s, but got %s (last day '1')", expected, nextTime)
	}
	recurPattern.Start = time.Date(2101, 2, 3, 0, 0, 0, 0, time.UTC)
	expected = time.Date(2101, 2, 28, 0, 0, 0, 0, time.UTC)
	if nextTime := recurPattern.NextTime(); !nextTime.Equal(expected) {
		t.Errorf("Expected %s, but got %s (last day '2')", expected, nextTime)
	}
	recurPattern.Start = time.Date(2101, 4, 5, 0, 0, 0, 0, time.UTC)
	expected = time.Date(2101, 4, 30, 0, 0, 0, 0, time.UTC)
	if nextTime := recurPattern.NextTime(); !nextTime.Equal(expected) {
		t.Errorf("Expected %s, but got %s (last day '3')", expected, nextTime)
	}
	expectedString := "Monthly, last day"
	if patternString := recurPattern.String(); patternString != expectedString {
		t.Errorf("Expected '%s', but got '%s' (last day '4')", expectedString, patternString)
	}
}

func TestBeforeLastDay(t *testing.T) {
	recurPattern := recur.Monthly().OnDay(2).FromLastDay().Build()
	recurPattern.Start = time.Date(2101, 1, 1, 0, 0, 0, 0, time.UTC)
	expected := time.Date(2101, 1, 29, 0, 0, 0, 0, time.UTC)
	if nextTime := recurPattern.NextTime(); !nextTime.Equal(expected) {
		t.Errorf("Expected %s, but got %s (2 days before last day '1')", expected, nextTime)
	}
	recurPattern.Start = time.Date(2101, 2, 3, 0, 0, 0, 0, time.UTC)
	expected = time.Date(2101, 2, 26, 0, 0, 0, 0, time.UTC)
	if nextTime := recurPattern.NextTime(); !nextTime.Equal(expected) {
		t.Errorf("Expected %s, but got %s (2 days before last day '2')", expected, nextTime)
	}
	recurPattern.Start = time.Date(2101, 4, 5, 0, 0, 0, 0, time.UTC)
	expected = time.Date(2101, 4, 28, 0, 0, 0, 0, time.UTC)
	if nextTime := recurPattern.NextTime(); !nextTime.Equal(expected) {
		t.Errorf("Expected %s, but got %s (2 days before last day '3')", expected, nextTime)
	}
	expectedString := "Monthly, 2 days before last day"
	if patternString := recurPattern.String(); patternString != expectedString {
		t.Errorf("Expected '%s', but got '%s' (2 days before last day '4')", expectedString, patternString)
	}
}

func TestOnLastDayOfQuarter(t *testing.T) {
	recurPattern := recur.OnMonths(3, 6, 9, 12).FromLastDay().Build()
	recurPattern.Start = time.Date(2101, 1, 1, 0, 0, 0, 0, time.UTC)
	expected := time.Date(2101, 3, 31, 0, 0, 0, 0, time.UTC)
	if nextTime := recurPattern.NextTime(); !nextTime.Equal(expected) {
		t.Errorf("Expected %s, but got %s (last day in quarter '1')", expected, nextTime)
	}
	recurPattern.Start = time.Date(2101, 4, 3, 0, 0, 0, 0, time.UTC)
	expected = time.Date(2101, 6, 30, 0, 0, 0, 0, time.UTC)
	if nextTime := recurPattern.NextTime(); !nextTime.Equal(expected) {
		t.Errorf("Expected %s, but got %s (last day in quarter '2')", expected, nextTime)
	}
	recurPattern.Start = time.Date(2101, 11, 5, 0, 0, 0, 0, time.UTC)
	expected = time.Date(2101, 12, 31, 0, 0, 0, 0, time.UTC)
	if nextTime := recurPattern.NextTime(); !nextTime.Equal(expected) {
		t.Errorf("Expected %s, but got %s (last day in quarter '3')", expected, nextTime)
	}
	recur.OnMonths()
	expectedString := "[Mar,Jun,Sep,Dec] last day"
	if patternString := recurPattern.String(); patternString != expectedString {
		t.Errorf("Expected '%s', but got '%s' (last day in quarter '4')", expectedString, patternString)
	}
}

func TestDayOnWeek(t *testing.T) {
	recurPattern := recur.Monthly().OnWeek(1, time.Saturday).Build()
	recurPattern.Start = time.Date(2100, 12, 31, 0, 0, 0, 0, time.UTC)
	expected := time.Date(2101, 1, 1, 0, 0, 0, 0, time.UTC)
	if nextTime := recurPattern.NextTime(); !nextTime.Equal(expected) {
		t.Errorf("Expected %s, but got %s (first Saterday '1')", expected, nextTime)
	}
	recurPattern.Start = time.Date(2101, 4, 30, 0, 0, 0, 0, time.UTC)
	expected = time.Date(2101, 5, 7, 0, 0, 0, 0, time.UTC)
	if nextTime := recurPattern.NextTime(); !nextTime.Equal(expected) {
		t.Errorf("Expected %s, but got %s (first Saterday '2')", expected, nextTime)
	}
	recurPattern.Start = time.Date(2101, 9, 1, 0, 0, 0, 0, time.UTC)
	expected = time.Date(2101, 9, 3, 0, 0, 0, 0, time.UTC)
	if nextTime := recurPattern.NextTime(); !nextTime.Equal(expected) {
		t.Errorf("Expected %s, but got %s (first Saterday '3')", expected, nextTime)
	}
	expectedString := "Monthly, on 1st Sat"
	if patternString := recurPattern.String(); patternString != expectedString {
		t.Errorf("Expected '%s', but got '%s' (first Saterday '4')", expectedString, patternString)
	}
	recurPattern = recur.Monthly().OnWeek(3, time.Monday).Build()
	recurPattern.Start = time.Date(2101, 7, 1, 0, 0, 0, 0, time.UTC)
	expected = time.Date(2101, 7, 18, 0, 0, 0, 0, time.UTC)
	if nextTime := recurPattern.NextTime(); !nextTime.Equal(expected) {
		t.Errorf("Expected %s, but got %s (3rd Monday '1')", expected, nextTime)
	}
	recurPattern.Start = time.Date(2101, 7, 19, 0, 0, 0, 0, time.UTC)
	expected = time.Date(2101, 8, 15, 0, 0, 0, 0, time.UTC)
	if nextTime := recurPattern.NextTime(); !nextTime.Equal(expected) {
		t.Errorf("Expected %s, but got %s (3rd Monday '2')", expected, nextTime)
	}
	expectedString = "Monthly, on 3rd Mon"
	if patternString := recurPattern.String(); patternString != expectedString {
		t.Errorf("Expected '%s', but got '%s' (3rd Monday '3')", expectedString, patternString)
	}
}

func TestDayOnLastWeek(t *testing.T) {
	recurPattern := recur.Monthly().OnLastWeek(time.Friday).Build()
	recurPattern.Start = time.Date(2101, 1, 1, 0, 0, 0, 0, time.UTC)
	expected := time.Date(2101, 1, 28, 0, 0, 0, 0, time.UTC)
	if nextTime := recurPattern.NextTime(); !nextTime.Equal(expected) {
		t.Errorf("Expected %s, but got %s (last Friday '1')", expected, nextTime)
	}
	recurPattern.Start = time.Date(2101, 4, 1, 0, 0, 0, 0, time.UTC)
	expected = time.Date(2101, 4, 29, 0, 0, 0, 0, time.UTC)
	if nextTime := recurPattern.NextTime(); !nextTime.Equal(expected) {
		t.Errorf("Expected %s, but got %s (last Friday '2')", expected, nextTime)
	}
	recurPattern.Start = time.Date(2101, 9, 1, 0, 0, 0, 0, time.UTC)
	expected = time.Date(2101, 9, 30, 0, 0, 0, 0, time.UTC)
	if nextTime := recurPattern.NextTime(); !nextTime.Equal(expected) {
		t.Errorf("Expected %s, but got %s (last Friday '3')", expected, nextTime)
	}
	expectedString := "Monthly, on last Fri"
	if patternString := recurPattern.String(); patternString != expectedString {
		t.Errorf("Expected '%s', but got '%s' (last Friday '4')", expectedString, patternString)
	}
}

func TestWeekly(t *testing.T) {
	recurPattern := recur.Weekly().OnDayOfWeek(time.Monday).Build()
	if nextTime := recurPattern.NextTime().Weekday(); nextTime != time.Monday {
		t.Errorf("Expected %s, but got %s (weekly '1')", time.Monday, nextTime)
	}
	expectedString := "Every 7 days"
	if patternString := recurPattern.String(); patternString != expectedString {
		t.Errorf("Expected '%s', but got '%s' (weekly '2')", expectedString, patternString)
	}
}

func TestDailyFromTo(t *testing.T) {
	recurPattern := recur.DailyWithin(time.Monday, time.Friday).Build()
	recurPattern.Start = time.Date(2101, 1, 1, 0, 0, 0, 0, time.UTC)
	expected := time.Date(2101, 1, 3, 0, 0, 0, 0, time.UTC)
	if nextTime := recurPattern.NextTime(); !nextTime.Equal(expected) {
		t.Errorf("Expected %s, but got %s (Mon - Fri '1')", expected, nextTime)
	}
	recurPattern.Start = time.Date(2101, 1, 6, 0, 0, 0, 0, time.UTC)
	expected = time.Date(2101, 1, 6, 0, 0, 0, 0, time.UTC)
	if nextTime := recurPattern.NextTime(); !nextTime.Equal(expected) {
		t.Errorf("Expected %s, but got %s (Mon - Fri '2')", expected, nextTime)
	}
	recurPattern.Start = time.Date(2101, 1, 8, 0, 0, 0, 0, time.UTC)
	expected = time.Date(2101, 1, 10, 0, 0, 0, 0, time.UTC)
	if nextTime := recurPattern.NextTime(); !nextTime.Equal(expected) {
		t.Errorf("Expected %s, but got %s (Mon - Fri '3')", expected, nextTime)
	}
	expectedString := "Every 1 day (Mon-Fri)"
	if patternString := recurPattern.String(); patternString != expectedString {
		t.Errorf("Expected '%s', but got '%s' (Mon - Fri '4')", expectedString, patternString)
	}
}

func TestEveryMonths(t *testing.T) {
	recurPattern := recur.EveryMonths(2).Build()
	recurPattern.Start = time.Date(2101, 1, 20, 0, 0, 0, 0, time.UTC)
	expected := time.Date(2101, 3, 1, 0, 0, 0, 0, time.UTC)
	if nextTime := recurPattern.NextTime(); !nextTime.Equal(expected) {
		t.Errorf("Expected %s, but got %s (Every 2 months '1')", expected, nextTime)
	}
	expectedString := "Every 2 months, 1st day"
	if patternString := recurPattern.String(); patternString != expectedString {
		t.Errorf("Expected '%s', but got '%s' (Every 2 months '2')", expectedString, patternString)
	}
	recurPattern = recur.EveryMonths(2).OnDay(15).Build()
	recurPattern.Start = time.Date(2101, 1, 16, 0, 0, 0, 0, time.UTC)
	expected = time.Date(2101, 3, 15, 0, 0, 0, 0, time.UTC)
	if nextTime := recurPattern.NextTime(); !nextTime.Equal(expected) {
		t.Errorf("Expected %s, but got %s (Every 2 months '3')", expected, nextTime)
	}
	expectedString = "Every 2 months, 15th day"
	if patternString := recurPattern.String(); patternString != expectedString {
		t.Errorf("Expected '%s', but got '%s' (Every 2 months '4')", expectedString, patternString)
	}
}

func TestEveryMonthsLastDay(t *testing.T) {
	recurPattern := recur.EveryMonths(2).OnDay(1).FromLastDay().Build()
	recurPattern.Start = time.Date(2101, 1, 5, 0, 0, 0, 0, time.UTC)
	expected := time.Date(2101, 2, 27, 0, 0, 0, 0, time.UTC)
	if nextTime := recurPattern.NextTime(); !nextTime.Equal(expected) {
		t.Errorf("Expected %s, but got %s (Every 2 months before last day '1')", expected, nextTime)
	}
	expectedString := "Every 2 months, 1 day before last day"
	if patternString := recurPattern.String(); patternString != expectedString {
		t.Errorf("Expected '%s', but got '%s' (Every 2 months before last day '2')", expectedString, patternString)
	}
}

func TestYearly(t *testing.T) {
	recurPattern := recur.Yearly().Build()
	expected := time.Date(recurPattern.Start.Year()+1, 1, 1, 0, 0, 0, 0, time.UTC)
	if nextTime := recurPattern.NextTime(); !nextTime.Equal(expected) {
		t.Errorf("Expected %s, but got %s (Yearly '1')", expected, nextTime)
	}
	expectedString := "Every 1 year, [Jan] 1st day"
	if patternString := recurPattern.String(); patternString != expectedString {
		t.Errorf("Expected '%s', but got '%s' (Yearly '2')", expectedString, patternString)
	}
}

func TestEveryYears(t *testing.T) {
	recurPattern := recur.EveryYears(2).Build()
	expected := time.Date(recurPattern.Start.Year()+2, 1, 1, 0, 0, 0, 0, time.UTC)
	if nextTime := recurPattern.NextTime(); !nextTime.Equal(expected) {
		t.Errorf("Expected %s, but got %s (Every 2 years '1')", expected, nextTime)
	}
	expectedString := "Every 2 years, [Jan] 1st day"
	if patternString := recurPattern.String(); patternString != expectedString {
		t.Errorf("Expected '%s', but got '%s' (Every 2 years '2')", expectedString, patternString)
	}
	recurPattern = recur.EveryYears(2).OnMonths(4).OnDay(5).Build()
	expected = time.Date(recurPattern.Start.Year()+2, 4, 5, 0, 0, 0, 0, time.UTC)
	if nextTime := recurPattern.NextTime(); !nextTime.Equal(expected) {
		t.Errorf("Expected %s, but got %s (Every 2 years '3')", expected, nextTime)
	}
	expectedString = "Every 2 years, [Apr] 5th day"
	if patternString := recurPattern.String(); patternString != expectedString {
		t.Errorf("Expected '%s', but got '%s' (Every 2 years '4')", expectedString, patternString)
	}
}

func TestEverySeconds(t *testing.T) {
	recurPattern := recur.EverySeconds(7)
	recurPattern.Start = recurPattern.Start.Add(-time.Second)
	seconds := int(recurPattern.NextTime().Sub(recurPattern.Start).Seconds())
	if seconds != 7 {
		t.Errorf("Expected %d, but got %d (Every 7 seconds '1')", 7, seconds)
	}
	expectedString := "Every 7s"
	if patternString := recurPattern.String(); patternString != expectedString {
		t.Errorf("Expected '%s', but got '%s' (Every 7 seconds '2')", expectedString, patternString)
	}
}

func TestEveryMinutes(t *testing.T) {
	recurPattern := recur.EveryMinutes(7).Build()
	recurPattern.Start = recurPattern.Start.Add(-time.Minute)
	minutes := int(recurPattern.NextTime().Sub(recurPattern.Start).Minutes())
	if minutes != 7 {
		t.Errorf("Expected %d, but got %d (Every 7 minutes '1')", 7, minutes)
	}
	expectedString := "Every 7m0s"
	if patternString := recurPattern.String(); patternString != expectedString {
		t.Errorf("Expected '%s', but got '%s' (Every 7 minutes '2')", expectedString, patternString)
	}
}

func TestEveryHours(t *testing.T) {
	recurPattern := recur.EveryHours(7).Build()
	recurPattern.Start = recurPattern.Start.Add(-time.Hour)
	hours := int(recurPattern.NextTime().Sub(recurPattern.Start).Hours())
	if hours != 7 {
		t.Errorf("Expected %d, but got %d (Every 7 hours '1')", 7, hours)
	}
	expectedString := "Every 7h0m0s"
	if patternString := recurPattern.String(); patternString != expectedString {
		t.Errorf("Expected '%s', but got '%s' (Every 7 hours '2')", expectedString, patternString)
	}
}

func TestEveryDays(t *testing.T) {
	recurPattern := recur.EveryDays(25).Build()
	days := int(recurPattern.NextTime().Sub(recurPattern.Start).Hours() / 24)
	if days == 7 {
		t.Errorf("Expected %d, but got %d (Every 25 days '1')", 7, days)
	}
	expectedString := "Every 25 days"
	if patternString := recurPattern.String(); patternString != expectedString {
		t.Errorf("Expected '%s', but got '%s' (Every 25 days '2')", expectedString, patternString)
	}
}
