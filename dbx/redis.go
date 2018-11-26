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

package dbx

// 如果编译时间过长，不使用时可注释
import (
	"fmt"

	"github.com/xuyu/goredis"
)

type Redis struct {
	Auth    string `json:"Auth"`
	Host    string `json:"Host"`
	Port    int    `json:"Port"`
	DbIndex int    `json:"DbIndex"`
	Timeout string `json:"Timeout"`
	Maxidle int    `json:"Maxidle"`
	Conn    *goredis.Redis
}

func NewRedis() *Redis {
	return &Redis{}
}

func (this *Redis) Init() error {
	var (
		err error
	)

	this.Conn, err = goredis.DialURL(fmt.Sprintf("tcp://auth:%s@%s:%d/%d?timeout=%s&maxidle=%d",
		this.Auth,
		this.Host,
		this.Port,
		this.DbIndex,
		this.Timeout,
		this.Maxidle))

	return err
}
