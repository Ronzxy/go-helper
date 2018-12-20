/* Copyright 2015 sky<skygangsta@hotmail.com>. All rights reserved.
 *
 * Licensed under the Apache License, version 2.0 (the "License").
 * You may not use this work except in compliance with the License, which is
 * available at www.apache.org/licenses/LICENSE-2.0
 *
 * This software is distributed on an "AS IS" basis, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied, as more fully set forth in the License.
 *
 * See the NOTICE file distributed with this work for information regarding copyright ownership.
 */

package util

import (
	"strings"
	"time"
)

type Date struct{}

func NewDate() *Date {
	return &Date{}
}

// Mon Jan 2 2006-01-02 15:04:05.999999999 -0700 MST
func (this *Date) Replacer() *strings.Replacer {
	patterns := []string{
		// less second
		"ms", "000", // millisecond
		"mi", "000000", // microseconds
		"ns", "000000000", // nanoseconds

		// year
		"Y", "2006", // A full numeric representation of a year, 4 digits	Examples: 1999 or 2003
		"yyyy", "2006", // A full numeric representation of a year, 4 digits	Examples: 1999 or 2003
		"yy", "06", // A two digit representation of a year	Examples: 99 or 03
		"y", "6", // A two digit representation of a year	Examples: 99 or 03

		// month
		"mm", "01", // Numeric representation of a month, with leading zeros	01 through 12
		"m", "1", // Numeric representation of a month, without leading zeros	1  through 12
		"JJ", "January", // Full string representation of a month
		"J", "Jan", // Abbreviated string representation of a month

		// week
		"ww", "Monday", // Full string representation of a week
		"w", "Mon", // Abbreviated string representation of a week

		// day
		"dd", "02", // Day of the month, 2 digits with leading zeros	01 to 31
		"d", "2", // Day of the month without leading zeros	1 to 31

		// hours
		"HH", "15", // 24-hour format of an hour with leading zeros	00 through 23
		"H", "15", // 24-hour format of an hour with leading zeros	00 through 23
		"hh", "03", // 12-hour format of an hour with leading zeros	 0 through 12
		"h", "3", // 12-hour format of an hour with leading zeros	00 through 12

		// minutes
		"MM", "04", // Minutes with leading zeros	00 to 59
		"M", "4", // Minutes with leading zeros	00 to 59

		// second
		"SS", "05", // Seconds, with leading zeros	00 through 59
		"S", "5", // Seconds, with leading zeros	00 through 59

		"tt", "PM", // Upper representation of a abbreviated time
		"t", "pm", // Lower representation of a abbreviated time

		// time zone
		"zz", "-0700", // Full numeric representation of time zone
		"z", "-07", // Abbreviated numeric representation of time zone
		"GMT", "MST", // Full string representation of time zone
		"G", "MST", // Abbreviated string representation of time zone
	}

	return strings.NewReplacer(patterns...)
}

// Format time with easy-to-understand strings, e.g. yyyy-mm-dd HH:MM:SS.ms
func (this *Date) Format(t time.Time, format string) string {
	format = this.Replacer().Replace(format)

	return t.Format(format)
}

// Use yyyy-mm-dd HH:MM:SS.ms format the time
func (this *Date) DefaultFormat(t time.Time) string {
	return this.Format(t, "yyyy-mm-dd HH:MM:SS.ms")
}

// Parse time with easy-to-understand strings, e.g. yyyy-mm-dd HH:MM:SS.ms
func (this *Date) Parse(format, value string) (time.Time, error) {
	format = this.Replacer().Replace(format)

	return time.Parse(format, value)
}
