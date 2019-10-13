/* Copyright 2016 sky<skygangsta@hotmail.com>. All rights reserved.
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

package algo

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
)

type HMac struct {
	Key string
}

func NewHMac() *HMac {
	return &HMac{}
}

func (this *HMac) Sha256(text string) string {
	hMac := hmac.New(sha256.New, []byte(this.Key))
	hMac.Write([]byte(text))
	return base64.StdEncoding.EncodeToString(hMac.Sum(nil))
}

func (this *HMac) Md5(text string) string {
	hMac := hmac.New(md5.New, []byte(this.Key))
	hMac.Write([]byte(text))
	return base64.StdEncoding.EncodeToString(hMac.Sum(nil))
}

func (this *HMac) Sha512(text string) string {
	hMac := hmac.New(sha512.New, []byte(this.Key))
	hMac.Write([]byte(text))
	return base64.StdEncoding.EncodeToString(hMac.Sum(nil))
}
