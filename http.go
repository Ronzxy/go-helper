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

package helper

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type HttpHelper struct {
	Method  string
	Addr    string
	Data    string // http POST data
	Timeout int    // http timeout second
	Header  struct {
		UserAgent   string
		ContentType string // "text/json; charset=utf-8"
	}
}

func NewHttpHelper() *HttpHelper {
	return NewHttpHelperWithUserAgent(fmt.Sprintf("%s/%d", NewPathHelper().WorkName(), time.Now().Year()))
}

func NewHttpHelperWithUserAgent(userAgent string) *HttpHelper {
	http := NewHttpHelper()
	http.Header.UserAgent = userAgent

	return http
}

func (this *HttpHelper) Do() (data []byte, err error) {
	var (
		req  *http.Request
		resp *http.Response

		client = &http.Client{
			Timeout: time.Duration(this.Timeout) * time.Second,
		}
	)

	req, err = http.NewRequest(this.Method, this.Addr, strings.NewReader(this.Data))
	if err != nil {
		err = errors.New(fmt.Sprintf("http.NewRequest - %s", err.Error()))
		return
	}

	if NewStringHelper().IsEmpty(this.Header.UserAgent) {
		req.Header.Set("User-Agent", fmt.Sprintf("%s/%d", NewPathHelper().WorkName(), time.Now().Year()))
	} else {
		req.Header.Set("User-Agent", this.Header.UserAgent)
	}

	if NewStringHelper().IsEmpty(this.Header.UserAgent) {
		req.Header.Set("Content-Type", "text/plain; charset=utf-8")
	} else {
		req.Header.Set("Content-Type", this.Header.ContentType)
	}

	resp, err = client.Do(req)
	if err != nil {
		err = errors.New(fmt.Sprintf("client.Do - %s", err.Error()))
		return
	}

	data, err = ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	return
}

func (this *HttpHelper) GET(addr string) ([]byte, error) {
	this.Method = "GET"
	this.Addr = addr

	return this.Do()
}

func (this *HttpHelper) POST(addr, data string) ([]byte, error) {
	this.Method = "POST"
	this.Addr = addr
	this.Data = data

	return this.Do()
}

func (this *HttpHelper) PostForm(addr string, data url.Values) ([]byte, error) {
	this.Method = "POST"
	this.Addr = addr
	this.Data = data.Encode()

	this.Header.ContentType = "application/x-www-form-urlencoded"

	return this.Do()
}

func (this *HttpHelper) RemoteAddr(req *http.Request) string {
	addr := req.Header.Get("X-Real-IP")
	if len(addr) == 0 {
		addr = req.Header.Get("X-Forwarded-For")
		if addr == "" {
			addr = req.RemoteAddr
			if i := strings.LastIndex(addr, ":"); i > -1 {
				addr = addr[:i]
			}
		} else {
			if i := strings.LastIndex(addr, ","); i > -1 {
				addr = addr[:i]
			}
		}
	}
	return addr
}
