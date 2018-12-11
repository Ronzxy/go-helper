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

/**
 * 使用易于理解的字符串格式化时间，Y-m-d H:M:S.ms
 */
func (this *Date) Format(t time.Time, format string) string {
	patterns := []string{
		// two letter
		"ms", "000", // millisecond
		"mi", "000000", // microseconds
		"ns", "000000000", // nanoseconds

		// year
		"Y", "2006", // A full numeric representation of a year, 4 digits	Examples: 1999 or 2003
		"y", "06", // A two digit representation of a year	Examples: 99 or 03

		// month
		"m", "01", // Numeric representation of a month, with leading zeros	01 through 12
		"n", "1", // Numeric representation of a month, without leading zeros	1 through 12

		// day
		"d", "02", // Day of the month, 2 digits with leading zeros	01 to 31
		"e", "2", // Day of the month without leading zeros	1 to 31

		// time
		"H", "15", // 24-hour format of an hour with leading zeros	00 through 23
		"M", "04", // Minutes with leading zeros	00 to 59
		"S", "05", // Seconds, with leading zeros	00 through 59
	}

	replacer := strings.NewReplacer(patterns...)
	format = replacer.Replace(format)

	return t.Format(format)
}

func (this *Date) DefaultFormat(t time.Time) string {
	return this.Format(t, "Y-m-d H:M:S.ms")
}
