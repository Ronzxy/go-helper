/* Copyright 2015 Ron Zhang <ronzxy@mx.aketi.cn>. All rights reserved.
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

var String = NewStringHelper()

type StringHelper struct{}

func NewStringHelper() *StringHelper {
	return &StringHelper{}
}

func (this *StringHelper) Empty() string {
	return ""
}

func (this *StringHelper) IsEmpty(str string) bool {
	return this.Empty() == str
}

func (this *StringHelper) IsNotEmpty(str string) bool {
	return this.Empty() != str
}

func (this *StringHelper) IsEqual(str, equal string) bool {
	if len(str) != len(equal) {
		return false
	}

	return str == equal
}

func (this *StringHelper) IsNotEqual(str, equal string) bool {
	return !this.IsEqual(str, equal)
}
