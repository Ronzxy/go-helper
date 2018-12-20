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
	"github.com/go-xorm/xorm"
	_ "github.com/mattn/go-sqlite3"
)

type Sqlite3 struct {
	Datafile string
	Conn     *xorm.Engine
}

func NewSqlite3() *Sqlite3 {
	return &Sqlite3{}
}

func (this *Sqlite3) Init() error {
	var (
		err error
	)

	this.Conn, err = xorm.NewEngine("sqlite3", this.Datafile)

	return err
}
