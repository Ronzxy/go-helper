/* Copyright 2016 Ron Zhang <ronzxy@hotmail.com>. All rights reserved.
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
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"
)

type Sha struct{}

func NewSha() *Sha {
	return &Sha{}
}

// md5
func (this *Sha) Md5(buf []byte) []byte {
	ctx := md5.New()
	ctx.Write(buf)

	return ctx.Sum(nil)
}

func (this *Sha) Md5ToHexString(text string) string {
	return hex.EncodeToString(this.Md5([]byte(text)))
}

func (this *Sha) Md5ToBase64String(text string) string {
	return base64.StdEncoding.EncodeToString(this.Md5([]byte(text)))
}

// sha1
func (this *Sha) Sha1(buf []byte) []byte {
	ctx := sha1.New()
	ctx.Write(buf)

	return ctx.Sum(nil)
}

func (this *Sha) Sha1ToHexString(text string) string {
	return hex.EncodeToString(this.Sha1([]byte(text)))
}

func (this *Sha) Sha1ToBase64String(text string) string {
	return base64.StdEncoding.EncodeToString(this.Sha1([]byte(text)))
}

// sha256
func (this *Sha) Sha256(buf []byte) []byte {
	ctx := sha256.New()
	ctx.Write(buf)

	return ctx.Sum(nil)
}

func (this *Sha) Sha256ToHexString(text string) string {
	return hex.EncodeToString(this.Sha256([]byte(text)))
}

func (this *Sha) Sha256ToBase64String(text string) string {
	return base64.StdEncoding.EncodeToString(this.Sha256([]byte(text)))
}

// sha512
func (this *Sha) Sha512(buf []byte) []byte {
	ctx := sha512.New()
	ctx.Write(buf)

	return ctx.Sum(nil)
}

func (this *Sha) Sha512ToHexString(text string) string {
	return hex.EncodeToString(this.Sha512([]byte(text)))
}

func (this *Sha) Sha512ToBase64String(text string) string {
	return base64.StdEncoding.EncodeToString(this.Sha512([]byte(text)))
}
