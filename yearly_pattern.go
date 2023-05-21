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
	"fmt"
	"time"
)

type YearlyPattern struct {
	recurPattern *RecurringPattern
	err          error
}

func (pattern *YearlyPattern) getTimePattern() *TimePattern {
	return &TimePattern{minPattern: &MinutesPattern{
		secPattern: &SecondsPattern{recurPattern: pattern.recurPattern}}}
}

func (pattern *YearlyPattern) Every(delay time.Duration) *RecurringPattern {
	pattern.recurPattern.waitTime = &delay
	return pattern.recurPattern
}

func (pattern *YearlyPattern) Daily() *TimePattern {
	day := 24 * time.Hour
	pattern.recurPattern.Start = pattern.recurPattern.Start.Truncate(day)
	pattern.recurPattern.waitTime = &day
	return pattern.getTimePattern()
}

func (pattern *YearlyPattern) DailyWithin(from time.Weekday, to time.Weekday) *TimePattern {
	timePattern := pattern.Daily()
	recurPattern := timePattern.minPattern.secPattern.recurPattern
	recurPattern.AllowedWeekdays = []Weekday{}
	recurPattern.AllowedWeekdays = make([]Weekday, 1)
	recurPattern.AllowedWeekdays[0] = Weekday{Day: from}
	if from != to {
		nxtDay := from
		for nxtDay != to {
			nxtDay = nextDay(nxtDay)
			recurPattern.AllowedWeekdays = append(recurPattern.AllowedWeekdays, Weekday{Day: nxtDay})
		}
	}
	return timePattern
}

func (pattern *YearlyPattern) EveryDays(days int) *TimePattern {
	pattern.err = checkInput("days", &days, 1, 2147483647)
	if pattern.err == nil {
		waitTime := time.Duration(days*24) * time.Hour
		pattern.recurPattern.waitTime = &waitTime
		return pattern.getTimePattern()
	} else {
		return nil
	}
}

func (pattern *YearlyPattern) Weekly() *WeeklyPattern {
	week := time.Duration(7*24) * time.Hour
	pattern.recurPattern.waitTime = &week
	return &WeeklyPattern{recurPattern: pattern.recurPattern}
}

func (pattern *YearlyPattern) EveryWeeks(weeks int) *WeeklyPattern {
	pattern.err = checkInput("weeks", &weeks, 1, 2147483647)
	if pattern.err == nil {
		week := time.Duration(weeks*7*24) * time.Hour
		pattern.recurPattern.waitTime = &week
		return &WeeklyPattern{recurPattern: pattern.recurPattern}
	} else {
		return nil
	}
}

func (pattern *YearlyPattern) EveryHours(hours int) *MinutesPattern {
	pattern.err = checkInput("hours", &hours, 1, 2147483647)
	if pattern.err == nil {
		waitTime := time.Duration(hours) * time.Hour
		pattern.recurPattern.waitTime = &waitTime
		return &MinutesPattern{secPattern: &SecondsPattern{
			recurPattern: pattern.recurPattern}}
	} else {
		return nil
	}
}

func (pattern *YearlyPattern) EveryMinutes(minutes int) *SecondsPattern {
	pattern.err = checkInput("minutes", &minutes, 1, 2147483647)
	if pattern.err == nil {
		waitTime := time.Duration(minutes) * time.Minute
		pattern.recurPattern.waitTime = &waitTime
		return &SecondsPattern{recurPattern: pattern.recurPattern}
	} else {
		return nil
	}
}

func (pattern *YearlyPattern) EverySeconds(seconds int) *RecurringPattern {
	pattern.err = checkInput("seconds", &seconds, 1, 2147483647)
	if pattern.err == nil {
		waitTime := time.Second * time.Duration(seconds)
		pattern.recurPattern.waitTime = &waitTime
		return pattern.recurPattern
	} else {
		return nil
	}
}

func (pattern *YearlyPattern) Monthly() *MonthlyPattern {
	pattern.recurPattern.Start = pattern.recurPattern.Start.
		Truncate(24 * time.Hour)
	return &MonthlyPattern{monthdayPattern: &MonthdayPattern{
		lastDayPattern: &LastMonthdayPattern{timePattern: &TimePattern{
			minPattern: &MinutesPattern{secPattern: &SecondsPattern{
				recurPattern: pattern.recurPattern}}}}}}
}

func (pattern *YearlyPattern) OnMonths(months ...int) *MonthlyPattern {
	if months == nil || len(months) == 0 {
		pattern.err = checkInput("month", nil, 1, 12)
		return nil
	}
	pattern.recurPattern.AllowedMonths = []int{}
	for _, month := range months {
		pattern.err = checkInput("month", &month, 1, 12)
		if pattern.err != nil {
			return nil
		}
		pattern.recurPattern.AllowedMonths = append(pattern.recurPattern.AllowedMonths, month)
	}
	return pattern.Monthly()
}

func (pattern *YearlyPattern) EveryMonths(months int) *MonthlyPattern {
	pattern.err = checkInput("months", &months, 1, 300)
	if pattern.err == nil {
		pattern.recurPattern.WaitMonths = &months
		return pattern.Monthly()
	} else {
		fmt.Println(pattern.err)
		return nil
	}
}

func (pattern *YearlyPattern) Build() *RecurringPattern {
	recurPattern := pattern.OnMonths(1).Build()
	if recurPattern.WaitYears == nil {
		first := 1
		recurPattern.WaitYears = &first
	}
	return recurPattern
}
