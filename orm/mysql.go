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

package orm

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/skygangsta/go-helper"
	"time"
)

type MySQL struct {
	Host            string `yaml:"Host"`
	Port            int    `yaml:"Port"`
	DbName          string `yaml:"DbName"`
	Username        string `yaml:"Username"`
	Password        string `yaml:"Password"`
	MaxIdleConns    int    `yaml:"MaxIdleConns"`
	MaxOpenConns    int    `yaml:"MaxOpenConns"`
	ConnMaxLifetime int    `yaml:"ConnMaxLifetime"`
	Timeout         int    `yaml:"Timeout"`
	AppName         string `yaml:"AppName"`
	// Instance
	Conn *xorm.Engine
}

func NewMySQL() *MySQL {
	return &MySQL{}
}

func (this *MySQL) init() {
	if this.Timeout <= 0 {
		this.Timeout = 10
	}

	if this.AppName == "" {
		this.AppName = helper.NewPathHelper().WorkName()
	}

	if this.MaxIdleConns <= 0 {
		this.MaxIdleConns = 5
	}

	if this.MaxOpenConns <= 0 {
		this.MaxOpenConns = 1024
	}

	if this.ConnMaxLifetime <= 0 {
		// default 60s
		this.ConnMaxLifetime = 60000
	}
}

func (this *MySQL) DriverName() string {
	return "mysql"
}

func (this *MySQL) ConnString() string {
	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?timeout=%ds&collation=utf8mb4_unicode_ci&autocommit=false&parseTime=true",
		this.Username,
		this.Password,
		this.Host,
		this.Port,
		this.DbName,
		this.Timeout)

	return connString
}

func (this *MySQL) Connect() error {
	var (
		err error
	)

	this.init()

	this.Conn, err = xorm.NewEngine(this.DriverName(), this.ConnString())
	if err == nil {
		this.Conn.SetMaxIdleConns(this.MaxIdleConns)
		this.Conn.SetMaxOpenConns(this.MaxOpenConns)
		this.Conn.SetConnMaxLifetime(time.Duration(this.ConnMaxLifetime))
	}

	return err
}
