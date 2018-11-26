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

// benchmark: go test -test.bench=".*"

// b.StopTimer() //调用该函数停止压力测试的时间计数
//做一些初始化的工作,例如读取文件数据,数据库连接之类的,
//这样这些时间不影响我们测试函数本身的性能

package cryptox

import (
	"fmt"
	"os"
	"testing"
)

func TestDES_HexString(t *testing.T) {
	DES, err := NewDES([]byte("01234567"), nil)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	output := DES.EncryptToHexString([]byte("我是DES"), 0)

	t.Log("DES.Encrypt", output)

	input, err := DES.DecryptFromHexString(output, 0)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	t.Log("DES.Decrypt", string(input))
}

func TestDES_Base64String(t *testing.T) {
	DES, err := NewDES([]byte("01234567"), nil)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	output := DES.EncryptToBase64String([]byte("我是DES"), 1)

	t.Log("DES.Encrypt", output)

	input, err := DES.DecryptFromBase64String(output, 1)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	t.Log("DES.Decrypt", string(input))
}

func BenchmarkDES_HexString(b *testing.B) {
	// b.StopTimer() //调用该函数停止压力测试的时间计数
	//做一些初始化的工作,例如读取文件数据,数据库连接之类的,
	//这样这些时间不影响我们测试函数本身的性能
	// b.StartTimer() //重新开始时间
	for i := 0; i < b.N; i++ {
		DES, err := NewDES([]byte("89abcdef"), nil)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		output := DES.EncryptToHexString([]byte("我是DES"), 0)
		// if err != nil {
		//  fmt.Println(err.Error())
		//  os.Exit(1)
		// }

		b.Log("DES.Encrypt", output)

		input, err := DES.DecryptFromHexString(output, 0)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		b.Log("DES.Decrypt", string(input))
	}
}

func BenchmarkDES_Base64String(b *testing.B) {
	// b.StopTimer() //调用该函数停止压力测试的时间计数
	//做一些初始化的工作,例如读取文件数据,数据库连接之类的,
	//这样这些时间不影响我们测试函数本身的性能
	// b.StartTimer() //重新开始时间
	for i := 0; i < b.N; i++ {
		DES, err := NewDES([]byte("89abcdef"), nil)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		output := DES.EncryptToBase64String([]byte("我是DES"), 1)
		// if err != nil {
		//  fmt.Println(err.Error())
		//  os.Exit(1)
		// }

		b.Log("DES.Encrypt", output)

		input, err := DES.DecryptFromBase64String(output, 1)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		b.Log("DES.Decrypt", string(input))
	}
}

func TestDES_CBC_HexString(t *testing.T) {
	DES, err := NewDES([]byte("01234567"), []byte("89abcdef"))
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	output := DES.CBCEncryptToHexString([]byte("我是DES_CBC"), 0)

	t.Log("DES.CBCEncrypt", output)

	input, err := DES.CBCDecryptFromHexString(output, 0)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	t.Log("DES.CBCDecrypt", string(input))
}

func TestDES_CBC_Base64String(t *testing.T) {
	DES, err := NewDES([]byte("01234567"), []byte("89abcdef"))
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	output := DES.CBCEncryptToBase64String([]byte("我是DES_CBC"), 1)

	t.Log("DES.CBCEncrypt", output)

	input, err := DES.CBCDecryptFromBase64String(output, 1)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	t.Log("DES.CBCDecrypt", string(input))
}

func BenchmarkDES_CBC_HexString(b *testing.B) {
	// b.StopTimer() //调用该函数停止压力测试的时间计数
	//做一些初始化的工作,例如读取文件数据,数据库连接之类的,
	//这样这些时间不影响我们测试函数本身的性能
	// b.StartTimer() //重新开始时间
	for i := 0; i < b.N; i++ {
		DES, err := NewDES([]byte("01234567"), []byte("89abcdef"))
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		output := DES.CBCEncryptToHexString([]byte("我是DES_CBC"), 0)

		b.Log("DES.CBCEncrypt", output)

		input, err := DES.CBCDecryptFromHexString(output, 0)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		b.Log("DES.CBCDecrypt", string(input))
	}
}

func BenchmarkDES_CBC_Base64String(b *testing.B) {
	// b.StopTimer() //调用该函数停止压力测试的时间计数
	//做一些初始化的工作,例如读取文件数据,数据库连接之类的,
	//这样这些时间不影响我们测试函数本身的性能
	// b.StartTimer() //重新开始时间
	for i := 0; i < b.N; i++ {
		DES, err := NewDES([]byte("89abcdef"), []byte("01234567"))
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		output := DES.CBCEncryptToBase64String([]byte("我是DES_CBC"), 1)

		b.Log("DES.CBCEncrypt", output)

		input, err := DES.CBCDecryptFromBase64String(output, 1)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		b.Log("DES.CBCDecrypt", string(input))
	}
}
