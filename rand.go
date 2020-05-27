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

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

var Rand = NewRandHelper()

type RandHelper struct{}

func NewRandHelper() *RandHelper {
	return &RandHelper{}
}

func (this *RandHelper) Num(length int) int64 {
	var s string
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	p := r.Perm(length)
	for i := 0; i < length; i++ {
		s = fmt.Sprintf("%s%d", s, p[i])
	}

	num, _ := strconv.ParseInt(s, 10, 64)

	return num
}

//生成随机字符串
func (this *RandHelper) Bytes(length int) []byte {
	buf := make([]byte, 0)
	str := []byte("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ+=/")
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		buf = append(buf, str[r.Intn(len(str))])
	}
	return buf
}

//生成随机字符串
func (this *RandHelper) String(length int) string {
	buf := this.Bytes(length)
	return string(buf)
}
