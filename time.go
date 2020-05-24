/* Copyright 2015 Ron Zhang <ronzxy@hotmail.com>. All rights reserved.
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

package helper

import (
	"strings"
	"time"
)

var Time = NewTimeHelper()

type TimeHelper struct{}

func NewTimeHelper() *TimeHelper {
	return &TimeHelper{}
}

// Mon Jan 2 2006-01-02 15:04:05.999999999 -0700 MST
//
// Thu, 10 Oct 2019 11:01:41 +0800 	--- w, dd J yyyy HH:MM:SS zz
// 2017-05-11 20:29:16 +0800 CST	---	yyyy-mm-dd HH:MM:SS zz G
func (this *TimeHelper) replacer() *strings.Replacer {
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
func (this *TimeHelper) Format(format string, t time.Time) string {
	format = this.replacer().Replace(format)

	return t.Format(format)
}

// Use yyyy-mm-dd HH:MM:SS.ms format the time
func (this *TimeHelper) DefaultFormat(t time.Time) string {
	return this.Format("yyyy-mm-dd HH:MM:SS.ms", t)
}

// Parse time with easy-to-understand strings, e.g. yyyy-mm-dd HH:MM:SS.ms
func (this *TimeHelper) Parse(format, value string) (time.Time, error) {
	format = this.replacer().Replace(format)

	return time.Parse(format, value)
}

// Convert time to the given location
func (this *TimeHelper) ConvertLocation(t time.Time, name string) (time.Time, error) {
	var (
		z   *time.Location
		err error
	)

	z, err = time.LoadLocation(name)
	if err != nil {
		return time.Time{}, err
	}

	// Convert to Unix time for formatted time format
	t = time.Unix(0, t.UnixNano())
	str := this.Format("yyyy-mm-dd HH:MM:SS.ns zz G", t)
	return time.ParseInLocation(this.replacer().Replace("yyyy-mm-dd HH:MM:SS.ns zz G"), str, z)
}
