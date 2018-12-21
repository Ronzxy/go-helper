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

package helper

import (
	"fmt"
	"reflect"
)

type StructMapHelper struct{}

func NewStructMapHelper() *StructMapHelper {
	return &StructMapHelper{}
}

func (this *StructMapHelper) MapString() string {
	return "map"
}

func (this *StructMapHelper) ToMap(obj interface{}, tag string) map[string]interface{} {
	m := make(map[string]interface{})

	v := reflect.ValueOf(obj)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	// we only accept structs
	if v.Kind() != reflect.Struct {
		fmt.Println("ToMap only accepts structs; got", v)
		return nil
	}

	t := v.Type()
	for i := 0; i < v.NumField(); i++ {
		// 获取字段Tag
		tag := t.Field(i).Tag.Get(tag)
		if tag != "" {
			// set key of map to value in struct field
			m[tag] = v.Field(i).Interface()
		}
	}
	return m
}

func (this *StructMapHelper) FromMap(obj interface{}, m map[string]interface{}, tag string) {
	v := reflect.ValueOf(obj)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	// we only accept structs
	if v.Kind() != reflect.Struct {
		fmt.Println("ToMap only accepts structs; got", v)
		return
	}

	t := v.Type()
	for i := 0; i < v.NumField(); i++ {
		// 获取字段名称
		fieldName := t.Field(i).Name

		value := reflect.ValueOf(m[fieldName])

		v.Field(i).Set(value)
	}
}

func (this *StructMapHelper) ToMapString(obj interface{}, tag string) map[string]string {
	m := make(map[string]string)

	v := reflect.ValueOf(obj)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	// we only accept structs
	if v.Kind() != reflect.Struct {
		fmt.Println("ToMap only accepts structs; got", v)
		return nil
	}

	t := v.Type()
	for i := 0; i < v.NumField(); i++ {
		// 获取字段Tag
		tag := t.Field(i).Tag.Get(tag)
		if tag != "" {
			// set key of map to value in struct field
			m[tag] = fmt.Sprintf("%v", v.Field(i).Interface())
		}
	}

	return m
}

func (this *StructMapHelper) FromMapString(obj interface{}, m map[string]string, tag string) {
	v := reflect.ValueOf(obj)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	// we only accept structs
	if v.Kind() != reflect.Struct {
		fmt.Println("ToMap only accepts structs; got", v)
		return
	}

	t := v.Type()
	for i := 0; i < v.NumField(); i++ {
		// 获取字段名称
		fieldName := t.Field(i).Name

		v.Field(i).SetString(m[fieldName])
	}
}
