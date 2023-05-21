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

type MonthdayPattern struct {
	lastDayPattern *LastMonthdayPattern
}

func (pattern *MonthdayPattern) FromLastDay() *LastMonthdayPattern {
	recurPattern := pattern.lastDayPattern.timePattern.minPattern.secPattern.recurPattern
	if recurPattern.AllowedDays == nil || len(recurPattern.AllowedDays) == 0 {
		recurPattern.AllowedDays = []Monthday{{IsLastDay: true, IsWorkday: false}}
	} else {
		recurPattern.AllowedDays[0].IsLastDay = true
	}
	return pattern.lastDayPattern
}

func (pattern *MonthdayPattern) OnWorkday() *TimePattern {
	return pattern.lastDayPattern.OnWorkday()
}

func (pattern *MonthdayPattern) AtHour(hour int) *MinutesPattern {
	return pattern.lastDayPattern.AtHour(hour)
}

func (pattern *MonthdayPattern) AtMinute(minute int) *SecondsPattern {
	return pattern.lastDayPattern.AtMinute(minute)
}

func (pattern *MonthdayPattern) AtSecond(second int) *RecurringPattern {
	return pattern.lastDayPattern.AtSecond(second)
}

func (pattern *MonthdayPattern) Build() *RecurringPattern {
	return pattern.lastDayPattern.Build()
}
