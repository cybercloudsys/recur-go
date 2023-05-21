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

import "time"

type MonthlyPattern struct {
	monthdayPattern *MonthdayPattern
	err             error
}

func (pattern *MonthlyPattern) OnDay(day int) *MonthdayPattern {
	pattern.err = checkInput("day", &day, 1, 31)
	if pattern.err == nil {
		pattern.monthdayPattern.lastDayPattern.timePattern.minPattern.secPattern.
			recurPattern.AllowedDays = []Monthday{{Day: &day}}
		return pattern.monthdayPattern
	} else {
		return nil
	}
}

func (pattern *MonthlyPattern) OnWeek(weekOfMonth int, dayOfWeek time.Weekday) *TimePattern {
	pattern.err = checkInput("weekOfMonth", &weekOfMonth, 1, 5)
	if pattern.err == nil {
		pattern.monthdayPattern.lastDayPattern.timePattern.minPattern.secPattern.
			recurPattern.AllowedWeekdays = []Weekday{{Day: dayOfWeek, WeekOfMonth: &weekOfMonth}}
		return pattern.monthdayPattern.lastDayPattern.timePattern
	} else {
		return nil
	}
}

func (pattern *MonthlyPattern) OnLastWeek(dayOfWeek time.Weekday) *TimePattern {
	pattern.monthdayPattern.lastDayPattern.timePattern.minPattern.secPattern.
		recurPattern.AllowedWeekdays = []Weekday{{Day: dayOfWeek, IsLastWeek: true}}
	return pattern.monthdayPattern.lastDayPattern.timePattern
}

func (pattern *MonthlyPattern) FromLastDay() *LastMonthdayPattern {
	return pattern.monthdayPattern.FromLastDay()
}

func (pattern *MonthlyPattern) OnWorkday() *TimePattern {
	return pattern.monthdayPattern.OnWorkday()
}

func (pattern *MonthlyPattern) AtHour(hour int) *MinutesPattern {
	return pattern.monthdayPattern.AtHour(hour)
}

func (pattern *MonthlyPattern) AtMinute(minute int) *SecondsPattern {
	return pattern.monthdayPattern.AtMinute(minute)
}

func (pattern *MonthlyPattern) AtSecond(second int) *RecurringPattern {
	return pattern.monthdayPattern.AtSecond(second)
}

func (pattern *MonthlyPattern) Build() *RecurringPattern {
	first := 1
	pattern.monthdayPattern.lastDayPattern.timePattern.minPattern.secPattern.
		recurPattern.AllowedDays = []Monthday{{Day: &first}}
	return pattern.monthdayPattern.Build()
}
