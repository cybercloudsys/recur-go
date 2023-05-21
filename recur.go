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

package recur

import (
	"errors"
	"strconv"
	"time"
)

type PatternBuilder interface {
	Build() *RecurringPattern
}

func checkInput(name string, value *int, min int, max int) error {
	if value == nil || *value < min || *value > max {
		var val string
		if value != nil {
			val = strconv.Itoa(*value)
		} else {
			val = "nil"
		}
		return errors.New("Recurring pattern is invalid - input field " + name + ": " +
			val + " must be between (" + strconv.Itoa(min) + "-" + strconv.Itoa(max) + ")")
	}
	return nil
}

type RecurSettings struct {
	DateTimeKind  *time.Location
	weeklyOffDays []time.Weekday
}

var Settings = RecurSettings{DateTimeKind: time.UTC, weeklyOffDays: []time.Weekday{time.Saturday, time.Sunday}}

func (settings *RecurSettings) SetWeeklyOffDays(offDays ...time.Weekday) error {
	length := len(offDays)
	err := checkInput("weeklyOffDays", &length, 1, 3)
	if err == nil {
		settings.weeklyOffDays = []time.Weekday{}
		for _, day := range offDays {
			if !contains(settings.weeklyOffDays, day) {
				settings.weeklyOffDays = append(settings.weeklyOffDays, day)
			}
		}
		return nil
	} else {
		return err
	}
}

func (settings *RecurSettings) GetWeeklyOffDays() []time.Weekday {
	return settings.weeklyOffDays
}

func Now() time.Time {
	return time.Now().In(Settings.DateTimeKind).Truncate(time.Second)
}

func Every(delay time.Duration) *RecurringPattern {
	p := YearlyPattern{recurPattern: NewRecurringPattern()}
	return p.Every(delay)
}

func Daily() *TimePattern {
	p := YearlyPattern{recurPattern: NewRecurringPattern()}
	return p.Daily()
}

func DailyWithin(from time.Weekday, to time.Weekday) *TimePattern {
	p := YearlyPattern{recurPattern: NewRecurringPattern()}
	return p.DailyWithin(from, to)
}

func EveryDays(days int) *TimePattern {
	p := YearlyPattern{recurPattern: NewRecurringPattern()}
	return p.EveryDays(days)
}

func Weekly() *WeeklyPattern {
	p := YearlyPattern{recurPattern: NewRecurringPattern()}
	return p.Weekly()
}

func EveryWeeks(weeks int) *WeeklyPattern {
	p := YearlyPattern{recurPattern: NewRecurringPattern()}
	return p.EveryWeeks(weeks)
}

func EveryHours(hours int) *MinutesPattern {
	p := YearlyPattern{recurPattern: NewRecurringPattern()}
	return p.EveryHours(hours)
}

func EveryMinutes(minutes int) *SecondsPattern {
	p := YearlyPattern{recurPattern: NewRecurringPattern()}
	return p.EveryMinutes(minutes)
}

func EverySeconds(seconds int) *RecurringPattern {
	p := YearlyPattern{recurPattern: NewRecurringPattern()}
	return p.EverySeconds(seconds)
}

func Monthly() *MonthlyPattern {
	p := YearlyPattern{recurPattern: NewRecurringPattern()}
	return p.Monthly()
}

func OnMonths(months ...int) *MonthlyPattern {
	p := YearlyPattern{recurPattern: NewRecurringPattern()}
	return p.OnMonths(months...)
}

func EveryMonths(months int) *MonthlyPattern {
	p := YearlyPattern{recurPattern: NewRecurringPattern()}
	return p.EveryMonths(months)
}

func Yearly() *YearlyPattern {
	p := Pattern{yearPattern: &YearlyPattern{recurPattern: NewRecurringPattern()}}
	return p.Yearly()
}

func EveryYears(years int) *YearlyPattern {
	p := Pattern{yearPattern: &YearlyPattern{recurPattern: NewRecurringPattern()}}
	return p.EveryYears(years)
}
