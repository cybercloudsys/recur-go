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

import "errors"

type Pattern struct {
	yearPattern *YearlyPattern
	err         error
}

func (pattern *Pattern) Yearly() *YearlyPattern {
	return pattern.EveryYears(1)
}

func (pattern *Pattern) EveryYears(years int) *YearlyPattern {
	pattern.err = checkInput("years", &years, 1, 50)
	if pattern.err == nil {
		pattern.yearPattern.recurPattern.WaitYears = &years
		return &YearlyPattern{recurPattern: pattern.yearPattern.recurPattern}
	} else {
		return nil
	}
}

func (pattern *Pattern) Build() *RecurringPattern {
	pattern.err = errors.New("Recurring pattern is invalid - Time period must be specified.")
	return nil
}
