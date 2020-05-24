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

package orm

import (
	"github.com/go-xorm/xorm"
	_ "github.com/mattn/go-sqlite3"
)

type SQLite3 struct {
	Datafile string
	Conn     *xorm.Engine
}

func NewSqlite3() *SQLite3 {
	return &SQLite3{}
}

func (this *SQLite3) DriverName() string {
	return "sqlite3"
}

func (this *SQLite3) Connect() error {
	var (
		err error
	)

	this.Conn, err = xorm.NewEngine(this.DriverName(), this.Datafile)

	return err
}
