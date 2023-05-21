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

type SecondsPattern struct {
	recurPattern *RecurringPattern
	err          error
}

func (pattern *SecondsPattern) AtSecond(second int) *RecurringPattern {
	pattern.err = checkInput("second", &second, 0, 59)
	if pattern.err == nil {
		pattern.recurPattern.Start = pattern.recurPattern.Start.Truncate(time.Minute).
			Add(time.Duration(second) * time.Second)
		return pattern.recurPattern
	} else {
		return nil
	}
}

func (pattern *SecondsPattern) Build() *RecurringPattern {
	return pattern.AtSecond(0)
}
