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

type Strings struct {}

func NewStrings() *Strings {
	return &Strings{}
}

func (this *Strings) Empty() string {
	return ""
}

func (this *Strings) IsEmpty(str string) bool {
	return this.Empty() == str
}
