/* Copyright 2018 Ron Zhang <ronzxy@hotmail.com>. All rights reserved.
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

var ConsoleColor = NewConsoleColorHelper()

type ConsoleColorHelper struct{}

func NewConsoleColorHelper() *ConsoleColorHelper {
	return &ConsoleColorHelper{}
}

func (this *ConsoleColorHelper) White() string {
	return "\033[30m"
}

func (this *ConsoleColorHelper) Red() string {
	return "\033[31m"
}

func (this *ConsoleColorHelper) Green() string {
	return "\033[32m"
}

func (this *ConsoleColorHelper) Yello() string {
	return "\033[33m"
}

func (this *ConsoleColorHelper) Blue() string {
	return "\033[34m"
}

func (this *ConsoleColorHelper) Magenta() string {
	return "\033[35m"
}

func (this *ConsoleColorHelper) Cyan() string {
	return "\033[36m"
}

func (this *ConsoleColorHelper) Clear() string {
	return "\033[0m"
}

func (this *ConsoleColorHelper) WhiteBackground() string {
	return "\033[97;40m"
}

func (this *ConsoleColorHelper) RedBackground() string {
	return "\033[97;41m"
}

func (this *ConsoleColorHelper) GreenBackground() string {
	return "\033[97;42m"
}

func (this *ConsoleColorHelper) YelloBackground() string {
	return "\033[97;43m"
}

func (this *ConsoleColorHelper) BlueBackground() string {
	return "\033[97;44m"
}

func (this *ConsoleColorHelper) MagentaBackground() string {
	return "\033[97;45m"
}

func (this *ConsoleColorHelper) CyanBackground() string {
	return "\033[97;46m"
}

// \033[90;47m
func (this *ConsoleColorHelper) GrayBackground() string {
	return "\033[97;47m"
}
