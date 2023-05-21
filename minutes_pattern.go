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

type MinutesPattern struct {
	secPattern *SecondsPattern
	err        error
}

func (pattern *MinutesPattern) AtMinute(minute int) *SecondsPattern {
	pattern.err = checkInput("minute", &minute, 0, 59)
	if pattern.err == nil {
		pattern.secPattern.recurPattern.Start = pattern.secPattern.recurPattern.Start.
			Truncate(time.Hour).Add(time.Duration(minute) * time.Minute)
		return pattern.secPattern
	} else {
		return nil
	}
}

func (pattern *MinutesPattern) AtSecond(second int) *RecurringPattern {
	return pattern.secPattern.AtSecond(second)
}

func (pattern *MinutesPattern) Build() *RecurringPattern {
	return pattern.AtMinute(0).Build()
}
