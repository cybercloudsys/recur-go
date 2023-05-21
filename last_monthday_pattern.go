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

type LastMonthdayPattern struct {
	timePattern *TimePattern
}

func (pattern *LastMonthdayPattern) OnWorkday() *TimePattern {
	pattern.timePattern.minPattern.secPattern.recurPattern.AllowedDays[0].IsWorkday = true
	return pattern.timePattern
}

func (pattern *LastMonthdayPattern) AtHour(hour int) *MinutesPattern {
	return pattern.timePattern.AtHour(hour)
}

func (pattern *LastMonthdayPattern) AtMinute(minute int) *SecondsPattern {
	return pattern.timePattern.AtMinute(minute)
}

func (pattern *LastMonthdayPattern) AtSecond(second int) *RecurringPattern {
	return pattern.timePattern.AtSecond(second)
}

func (pattern *LastMonthdayPattern) Build() *RecurringPattern {
	return pattern.timePattern.Build()
}
