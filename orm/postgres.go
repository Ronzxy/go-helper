/* Copyright 2016 Ron Zhang <ronzxy@mx.aketi.cn>. All rights reserved.
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
	_ "github.com/lib/pq"
	"github.com/ronzxy/go-helper"
	"time"
	"xorm.io/xorm"
)

type Postgres struct {
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
	IsCitus         bool   `yaml:"IsCitus"`
	// Instance
	Conn *xorm.Engine
}

func NewPostgres() *Postgres {
	return &Postgres{}
}

func (this *Postgres) init() {
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

func (this *Postgres) DriverName() string {
	return "postgres"
}

func (this *Postgres) ConnString() string {
	connString := fmt.Sprintf("%s://%s:%s@%s:%d/%s?sslmode=disable&application_name=%s&connect_timeout=%d",
		this.DriverName(),
		this.Username,
		this.Password,
		this.Host,
		this.Port,
		this.DbName,
		this.AppName,
		this.Timeout)

	return connString
}

func (this *Postgres) Connect() error {
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
