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

type TimePattern struct {
	minPattern *MinutesPattern
	err        error
}

func (pattern *TimePattern) AtHour(hour int) *MinutesPattern {
	pattern.err = checkInput("hour", &hour, 0, 59)
	if pattern.err == nil {
		pattern.minPattern.secPattern.recurPattern.Start = pattern.minPattern.secPattern.recurPattern.
			Start.Truncate(24 * time.Hour).Add(time.Duration(hour) * time.Hour)
		return pattern.minPattern
	} else {
		return nil
	}
}

func (pattern *TimePattern) AtMinute(minute int) *SecondsPattern {
	return pattern.minPattern.AtMinute(minute)
}

func (pattern *TimePattern) AtSecond(second int) *RecurringPattern {
	return pattern.minPattern.AtSecond(second)
}

func (pattern *TimePattern) Build() *RecurringPattern {
	return pattern.AtHour(0).Build()
}
