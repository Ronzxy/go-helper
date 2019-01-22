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
	"github.com/go-xorm/xorm"
	_ "github.com/lib/pq"
	"github.com/skygangsta/go-helper"
)

type Postgres struct {
	Host         string
	Port         int
	DbName       string
	Dba          string
	Pwd          string
	MaxIdleConns int
	MaxOpenConns int
	Timeout      int
	AppName      string
	Conn         *xorm.Engine
}

func NewPostgres() *Postgres {
	return &Postgres{}
}

func (this *Postgres) DriverName() string {
	return "postgres"
}

func (this *Postgres) ConnString() string {
	connString := fmt.Sprintf("%s://%s:%s@%s:%d/%s?sslmode=disable&application_name=%s&connect_timeout=%d",
		this.DriverName(),
		this.Dba,
		this.Pwd,
		this.Host,
		this.Port,
		this.DbName,
		this.AppName,
		this.Timeout)

	return connString
}

func (this *Postgres) Init() error {
	var (
		err error
	)

	if this.Timeout == 0 {
		this.Timeout = 10
	}

	if this.AppName == "" {
		this.AppName = helper.NewPathHelper().WorkName()
	}

	this.Conn, err = xorm.NewEngine(this.DriverName(), this.ConnString())
	if err == nil {
		this.Conn.SetMaxIdleConns(this.MaxIdleConns)
		this.Conn.SetMaxOpenConns(this.MaxOpenConns)
	}

	return err
}
