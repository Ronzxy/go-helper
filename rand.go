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
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

type Rand struct{}

func NewRand() *Rand {
	return &Rand{}
}

func (this *Rand) Num(length int) int64 {
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
func (this *Rand) Bytes(length int) []byte {
	buf := make([]byte, 0)
	str := []byte("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ+=/")
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		buf = append(buf, str[r.Intn(len(str))])
	}
	return buf
}

//生成随机字符串
func (this *Rand) String(length int) string {
	buf := this.Bytes(length)
	return string(buf)
}
