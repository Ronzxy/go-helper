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

func (this *Date) Format(t time.Time, format string) string {
	//patterns := []string{
	//	// year
	//	"Y", "2006", // A full numeric representation of a year, 4 digits	Examples: 1999 or 2003
	//	"y", "06", //A two digit representation of a year	Examples: 99 or 03
	//
	//	// month
	//	"m", "01", // Numeric representation of a month, with leading zeros	01 through 12
	//	"n", "1", // Numeric representation of a month, without leading zeros	1 through 12
	//	"M", "Jan", // A short textual representation of a month, three letters	Jan through Dec
	//	"F", "January", // A full textual representation of a month, such as January or March	January through December
	//
	//	// day
	//	"d", "02", // Day of the month, 2 digits with leading zeros	01 to 31
	//	"j", "2", // Day of the month without leading zeros	1 to 31
	//
	//	// week
	//	"D", "Mon", // A textual representation of a day, three letters	Mon through Sun
	//	"l", "Monday", // A full textual representation of the day of the week	Sunday through Saturday
	//
	//	// time
	//	"g", "3", // 12-hour format of an hour without leading zeros	1 through 12
	//	"G", "15", // 24-hour format of an hour without leading zeros	0 through 23
	//	"h", "03", // 12-hour format of an hour with leading zeros	01 through 12
	//	"H", "15", // 24-hour format of an hour with leading zeros	00 through 23
	//
	//	"a", "pm", // Lowercase Ante meridiem and Post meridiem	am or pm
	//	"A", "PM", // Uppercase Ante meridiem and Post meridiem	AM or PM
	//
	//	"i", "04", // Minutes with leading zeros	00 to 59
	//	"s", "05", // Seconds, with leading zeros	00 through 59
	//
	//	"S", "000", // millisecond
	//}
	patterns := []string{
		// two letter
		"ms", "000", 		// millisecond
		"mi", "00000", 		// microseconds
		"ns", "00000000", 	// nanoseconds

		// year
		"Y", "2006", 	// A full numeric representation of a year, 4 digits	Examples: 1999 or 2003
		"y", "06", 		//A two digit representation of a year	Examples: 99 or 03

		// month
		"m", "01", 		// Numeric representation of a month, with leading zeros	01 through 12
		"n", "1", 		// Numeric representation of a month, without leading zeros	1 through 12

		// day
		"d", "02", 		// Day of the month, 2 digits with leading zeros	01 to 31
		"j", "2", 		// Day of the month without leading zeros	1 to 31

		// time
		"H", "15", 		// 24-hour format of an hour with leading zeros	00 through 23
		"M", "04", 		// Minutes with leading zeros	00 to 59
		"S", "05", 		// Seconds, with leading zeros	00 through 59
	}

	replacer := strings.NewReplacer(patterns...)
	format = replacer.Replace(format)

	return t.Format(format)
}
