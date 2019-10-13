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
	"crypto/tls"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"
)

var Http = NewHttpHelper(90, false)

type HttpHelper struct {
	client    *http.Client
	UserAgent string
}

func NewHttpHelper(timeout int, insecure bool) *HttpHelper {
	this := &HttpHelper{}

	this.client = &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyFromEnvironment,
			DialContext: (&net.Dialer{
				Timeout:   time.Duration(timeout) * time.Second,
				KeepAlive: time.Duration(timeout) * time.Second,
			}).DialContext,
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: insecure,
			},
			MaxIdleConns:        64,
			MaxIdleConnsPerHost: 64,
			IdleConnTimeout:     time.Duration(timeout) * time.Second,
		},
	}

	return this
}

func (this *HttpHelper) Do(r *http.Request) (*http.Response, error) {

	return this.client.Do(r)
}

func (this *HttpHelper) NewRequest(method, addr string, body io.Reader) (*http.Request, error) {
	var (
		r   *http.Request
		err error
	)

	r, err = http.NewRequest(method, addr, body)
	if err != nil {
		return nil, err
	}

	if NewStringHelper().IsNotEmpty(this.UserAgent) {
		r.Header.Set("User-Agent", this.UserAgent)
	}

	return r, nil
}

func (this *HttpHelper) Get(addr string) ([]byte, error) {
	var (
		request  *http.Request
		response *http.Response
		err      error
	)

	request, err = this.NewRequest("GET", addr, nil)
	if err != nil {
		return nil, err
	}

	response, err = this.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("server return status code %d", response.StatusCode)
	}

	return ioutil.ReadAll(response.Body)
}

func (this *HttpHelper) Post(addr string, body io.Reader) ([]byte, error) {
	var (
		request  *http.Request
		response *http.Response
		err      error
	)

	request, err = this.NewRequest("POST", addr, body)
	if err != nil {
		return nil, err
	}

	response, err = this.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("server return status code %d", response.StatusCode)
	}

	return ioutil.ReadAll(response.Body)
}

func (this *HttpHelper) PostForm(addr string, data url.Values) ([]byte, error) {
	var (
		request  *http.Request
		response *http.Response
		err      error
	)

	request, err = this.NewRequest("POST", addr, strings.NewReader(data.Encode()))
	if err != nil {
		return nil, err
	}

	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	response, err = this.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("server return status code %d", response.StatusCode)
	}

	return ioutil.ReadAll(response.Body)
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
