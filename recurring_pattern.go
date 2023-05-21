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
	"encoding/json"
	"math"
	"strconv"
	"strings"
	"time"
)

type RecurringPattern struct {
	Start           time.Time
	waitTime        *time.Duration
	WaitSeconds     *int       `json:",omitempty"`
	WaitMonths      *int       `json:",omitempty"`
	WaitYears       *int       `json:",omitempty"`
	AllowedWeekdays []Weekday  `json:",omitempty"`
	Weeks           []int      `json:",omitempty"`
	AllowedDays     []Monthday `json:",omitempty"`
	AllowedMonths   []int      `json:",omitempty"`
	WeeklyOffDays   []time.Weekday
}

func (pattern *RecurringPattern) Build() *RecurringPattern {
	return pattern
}

func NewRecurringPattern() *RecurringPattern {
	pattern := RecurringPattern{}
	pattern.Start = Now()
	pattern.SetWeeklyOffDays(Settings.GetWeeklyOffDays()...)
	return &pattern
}

func (recurPattern *RecurringPattern) GetWaitTime() *time.Duration {
	return recurPattern.waitTime
}

func (recurPattern *RecurringPattern) NextTime() time.Time {
	nextDate := Now()
	if nextDate.Before(recurPattern.Start) {
		nextDate = recurPattern.Start
	}
	for !recurPattern.IsMatching(nextDate) {
		if recurPattern.waitTime != nil {
			nextDate = nextDate.Add(time.Duration(recurPattern.waitTime.Seconds()-
				math.Mod(math.Ceil(nextDate.Sub(recurPattern.Start).Seconds()),
					recurPattern.waitTime.Seconds())) * time.Second)
		} else if recurPattern.AllowedDays != nil && len(recurPattern.AllowedDays) > 0 ||
			recurPattern.AllowedWeekdays != nil && len(recurPattern.AllowedWeekdays) > 0 {
			if nextDate.Hour() != recurPattern.Start.Hour() || nextDate.Minute() != recurPattern.Start.Minute() ||
				nextDate.Second() != recurPattern.Start.Second() || nextDate.Nanosecond() > 0 {
				newDate := nextDate.Truncate(24 * time.Hour).Add(recurPattern.Start.Sub(recurPattern.Start.Truncate(24 * time.Hour)))
				if newDate.After(nextDate) {
					nextDate = newDate
				} else {
					nextDate = newDate.AddDate(0, 0, 1)
				}
			} else {
				nextDate = nextDate.AddDate(0, 0, 1)
			}
		}
	}
	return nextDate
}

func nextDay(day time.Weekday) time.Weekday {
	if day == time.Saturday {
		return time.Sunday
	} else {
		return day + 1
	}
}

func previousDay(day time.Weekday) time.Weekday {
	if day == time.Sunday {
		return time.Saturday
	} else {
		return day - 1
	}
}

func (recurPattern *RecurringPattern) getDay(dayOfMonth *Monthday, chkTime time.Time) int {
	theday := chkTime.Truncate(24 * time.Hour)
	if dayOfMonth.IsWorkday {
		day := 0
		if dayOfMonth.Day != nil {
			day = *dayOfMonth.Day
		}
		lastDay := theday.AddDate(0, 1, -chkTime.Day()).Day()
		if dayOfMonth.IsLastDay {
			day = lastDay - day
		}
		weekday := theday.AddDate(0, 0, day-chkTime.Day()).Weekday()
		if contains(recurPattern.WeeklyOffDays, weekday) {
			prev := contains(recurPattern.WeeklyOffDays, previousDay(weekday))
			if contains(recurPattern.WeeklyOffDays, nextDay(nextDay((weekday)))) {
				if day > 1 {
					return day - 1
				} else {
					return 4
				}
			} else if contains(recurPattern.WeeklyOffDays, nextDay(weekday)) {
				if prev {
					if day > 2 {
						return day - 2
					} else {
						return 3
					}
				} else {
					if day > 1 {
						return day - 1
					} else {
						return 3
					}
				}
			} else if contains(recurPattern.WeeklyOffDays, previousDay(previousDay((weekday)))) {
				if day < lastDay-1 {
					return day + 2
				} else {
					return day - 3
				}
			} else if prev {
				if day < lastDay {
					return day + 1
				} else {
					return day - 2
				}
			}
		}
		return day
	} else if dayOfMonth.IsLastDay {
		lastDay := theday.AddDate(0, 1, -chkTime.Day()).Day()
		if dayOfMonth.Day != nil {
			return lastDay - *dayOfMonth.Day
		}
		return lastDay
	} else {
		return *dayOfMonth.Day
	}
}

func (recurPattern *RecurringPattern) IsMatching(chkTime time.Time) bool {
	return (chkTime.Equal(recurPattern.Start) || chkTime.After(recurPattern.Start)) &&
		(recurPattern.waitTime == nil || math.Mod(math.Floor(chkTime.Sub(recurPattern.Start).Seconds()), recurPattern.waitTime.Seconds()) == 0) &&
		(recurPattern.AllowedDays == nil || (recurPattern.anyDay(recurPattern.AllowedDays, chkTime.Day(), chkTime) &&
			math.Floor(recurPattern.Start.Sub(recurPattern.Start.Truncate(24*time.Hour)).Seconds()) ==
				math.Floor(chkTime.Sub(chkTime.Truncate(24*time.Hour)).Seconds()))) &&
		(recurPattern.AllowedWeekdays == nil || recurPattern.anyWeekday(recurPattern.AllowedWeekdays, chkTime.Weekday(), chkTime)) &&
		(recurPattern.AllowedMonths == nil || contains(recurPattern.AllowedMonths, int(chkTime.Month()))) &&
		(recurPattern.WaitMonths == nil || recurPattern.waitTime != nil ||
			math.Mod(math.Ceil(chkTime.Sub(recurPattern.Start).Hours()/(24*30)), float64(*recurPattern.WaitMonths)) == 0) &&
		(recurPattern.WaitYears == nil || recurPattern.WaitMonths != nil || recurPattern.waitTime != nil ||
			math.Mod(math.Ceil(chkTime.Sub(recurPattern.Start).Hours()/(24*365)), float64(*recurPattern.WaitYears)) == 0)
}

func (recurPattern *RecurringPattern) anyDay(array []Monthday, item int, chkTime time.Time) bool {
	for _, v := range array {
		if recurPattern.getDay(&v, chkTime) == item {
			return true
		}
	}
	return false
}

func (recurPattern *RecurringPattern) anyWeekday(array []Weekday, item time.Weekday, chkTime time.Time) bool {
	for _, d := range array {
		if d.Day == item && (d.WeekOfMonth == nil || *d.WeekOfMonth == ((chkTime.Day()-1)/7+1)) &&
			(!d.IsLastWeek || chkTime.AddDate(0, 1, -chkTime.Day()).Day()-chkTime.Day() < 7) &&
			math.Floor(float64(recurPattern.Start.Sub(recurPattern.Start.Truncate(24*time.Hour)).Seconds())) ==
				math.Floor(float64(chkTime.Sub(chkTime.Truncate(24*time.Hour)).Seconds())) {
			return true
		}
	}
	return false
}

func (recurPattern *RecurringPattern) SetWeeklyOffDays(offDays ...time.Weekday) error {
	length := len(offDays)
	err := checkInput("weeklyOffDays", &length, 1, 3)
	if err == nil {
		recurPattern.WeeklyOffDays = []time.Weekday{}
		for _, day := range offDays {
			if !contains(recurPattern.WeeklyOffDays, day) {
				recurPattern.WeeklyOffDays = append(recurPattern.WeeklyOffDays, day)
			}
		}
		return nil
	} else {
		return err
	}
}

func contains[T comparable](arr []T, value T) bool {
	for _, item := range arr {
		if item == value {
			return true
		}
	}
	return false
}

func containsLower(arr []string, value string) bool {
	for _, item := range arr {
		if strings.ToLower(item) == strings.ToLower(value) {
			return true
		}
	}
	return false
}

func (recurPattern *RecurringPattern) String() string {
	var output strings.Builder
	if recurPattern.waitTime != nil {
		output.WriteString("Every ")
		wtime := *recurPattern.waitTime
		days := int(math.Floor(wtime.Hours() / 24))
		if days >= 1 {
			output.WriteString(strconv.Itoa(days))
			output.WriteString(" day")
			if days > 1 {
				output.WriteString("s")
			}
			wtime = wtime - time.Duration(days*24*int(time.Hour))
			if wtime.Seconds() > 0 {
				output.WriteString(" ")
			}
		}
		if wtime.Seconds() > 0 {
			output.WriteString(wtime.String())
		}
		ln := len(recurPattern.AllowedWeekdays)
		if recurPattern.AllowedWeekdays != nil && ln > 0 {
			weekdays := make([]time.Weekday, ln)
			for i := 0; i < ln; i++ {
				weekdays[i] = recurPattern.AllowedWeekdays[i].Day
			}
			writeStringOfWeekdays(weekdays, &output)
		}
	} else {
		if recurPattern.AllowedMonths == nil && recurPattern.WaitMonths == nil &&
			recurPattern.WaitYears == nil && (recurPattern.AllowedDays != nil ||
			recurPattern.AllowedWeekdays != nil) {
			output.WriteString("Monthly, ")
		} else if recurPattern.AllowedMonths != nil {
			if recurPattern.WaitYears != nil {
				output.WriteString("Every ")
				output.WriteString(strconv.Itoa(*recurPattern.WaitYears))
				output.WriteString(" year")
				if *recurPattern.WaitYears > 1 {
					output.WriteString("s")
				}
				output.WriteString(", ")
			} else if len(recurPattern.AllowedMonths) == 1 {
				output.WriteString("Yearly, ")
			}
		}
		if recurPattern.AllowedDays != nil {
			if recurPattern.WaitMonths != nil {
				output.WriteString("Every ")
				output.WriteString(strconv.Itoa(*recurPattern.WaitMonths))
				output.WriteString(" month")
				if *recurPattern.WaitMonths > 1 {
					output.WriteString("s")
				}
				output.WriteString(", ")
			} else if recurPattern.AllowedMonths != nil {
				output.WriteString("[")
				ln := len(recurPattern.AllowedMonths) - 1
				for i, month := range recurPattern.AllowedMonths {
					output.WriteString(time.Month(month).String()[:3])
					if i < ln {
						output.WriteString(",")
					}
				}
				output.WriteString("] ")
			}
			for _, day := range recurPattern.AllowedDays {
				if day.IsLastDay {
					if day.Day != nil {
						output.WriteString(strconv.Itoa(*day.Day))
						output.WriteString(" day")
						if *day.Day > 1 {
							output.WriteString("s")
						}
						output.WriteString(" before last day")
					} else {
						output.WriteString("last day")
					}
				} else {
					writeNth(*day.Day, &output)
					output.WriteString(" day")
				}
				if day.IsWorkday {
					weekdays := []time.Weekday{}
					for day := time.Sunday; day <= time.Saturday; day++ {
						if !contains(recurPattern.WeeklyOffDays, day) {
							weekdays = append(weekdays, day)
						}
					}
					writeStringOfWeekdays(weekdays, &output)
				}
			}
		}
		if recurPattern.AllowedWeekdays != nil && len(recurPattern.AllowedWeekdays) > 0 {
			output.WriteString("on ")
			for _, weekday := range recurPattern.AllowedWeekdays {
				if weekday.IsLastWeek {
					output.WriteString("last")
				} else {
					writeNth(*weekday.WeekOfMonth, &output)
				}
				output.WriteString(" ")
				output.WriteString(weekday.Day.String()[:3])
			}
		}
	}
	return output.String()
}

func writeStringOfWeekdays(days []time.Weekday, output *strings.Builder) string {
	sortedDays := make([]time.Weekday, len(days))
	i := 0
	var offDay *time.Weekday
	for day := time.Sunday; day <= time.Saturday; day++ {
		if contains(days, day) {
			if offDay != nil {
				sortedDays[i] = day
				i++
			}
		} else if offDay == nil {
			newOffDay := day
			offDay = &newOffDay
		}
	}
	for day := time.Sunday; day <= time.Saturday; day++ {
		if contains(days, day) && (offDay == nil || day < *offDay) {
			sortedDays[i] = day
			i++
		}
	}
	output.WriteString(" (")
	started := false
	for _, day := range sortedDays {
		prev := contains(days, previousDay(day))
		if started && !prev {
			output.WriteString(",")
		}
		if !prev || (offDay == nil && day == time.Sunday) {
			output.WriteString(day.String()[:3])
			started = true
		} else if started && (!contains(days, nextDay(day)) ||
			(offDay == nil && day == time.Saturday)) {
			output.WriteString("-")
			output.WriteString(day.String()[:3])
		}
	}
	output.WriteString(")")
	return output.String()
}

func writeNth(number int, output *strings.Builder) {
	result := strconv.Itoa(number)
	output.WriteString(result)
	if strings.HasSuffix(result, "1") && !strings.HasSuffix(result, "11") {
		output.WriteString("st")
	} else if strings.HasSuffix(result, "2") && !strings.HasSuffix(result, "12") {
		output.WriteString("nd")
	} else if strings.HasSuffix(result, "3") && !strings.HasSuffix(result, "13") {
		output.WriteString("rd")
	} else {
		output.WriteString("th")
	}
}

func (recurPattern *RecurringPattern) setWaitTime() {
	if recurPattern.WaitSeconds != nil {
		sec := time.Duration(*recurPattern.WaitSeconds) * time.Second
		recurPattern.waitTime = &sec
	}
}

func (recurPattern *RecurringPattern) setWaitSeconds() {
	if recurPattern.waitTime != nil {
		sec := int(recurPattern.waitTime.Seconds())
		recurPattern.WaitSeconds = &sec
	}
}

func (pattern *RecurringPattern) MarshalJSON() ([]byte, error) {
	pattern.setWaitSeconds()
	type RecurPattern RecurringPattern
	recurPattern := &struct{ *RecurPattern }{RecurPattern: (*RecurPattern)(pattern)}
	return json.Marshal(recurPattern)
}

func (pattern *RecurringPattern) UnmarshalJSON(data []byte) error {
	type RecurPattern RecurringPattern
	recurPattern := &struct{ *RecurPattern }{RecurPattern: (*RecurPattern)(pattern)}
	if err := json.Unmarshal(data, recurPattern); err != nil {
		return err
	}
	pattern.setWaitTime()
	return nil
}
