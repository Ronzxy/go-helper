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

package algo

import (
	"fmt"
	"os"
	"testing"
)

func TestTripleDES_HexString(t *testing.T) {
	TripleDES, err := NewTripleDES([]byte("0123456789abcdef"), nil)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	output := TripleDES.EncryptToHexString([]byte("我是TripleDES"), 0)

	t.Log("TripleDES.Encrypt", output)

	input, err := TripleDES.DecryptFromHexString(output, 0)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	t.Log("TripleDES.Decrypt", string(input))
}

func TestTripleDES_Base64String(t *testing.T) {
	TripleDES, err := NewTripleDES([]byte("0123456789abcdef"), nil)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	output := TripleDES.EncryptToBase64String([]byte("我是TripleDES"), 1)

	t.Log("TripleDES.Encrypt", output)

	input, err := TripleDES.DecryptFromBase64String(output, 1)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	t.Log("TripleDES.Decrypt", string(input))
}

func BenchmarkTripleDES_HexString(b *testing.B) {
	// b.StopTimer() //调用该函数停止压力测试的时间计数
	//做一些初始化的工作,例如读取文件数据,数据库连接之类的,
	//这样这些时间不影响我们测试函数本身的性能
	// b.StartTimer() //重新开始时间
	for i := 0; i < b.N; i++ {
		TripleDES, err := NewTripleDES([]byte("0123456789abcdef"), nil)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		output := TripleDES.EncryptToHexString([]byte("我是TripleDES"), 0)
		// if err != nil {
		//  fmt.Println(err.Error())
		//  os.Exit(1)
		// }

		b.Log("TripleDES.Encrypt", output)

		input, err := TripleDES.DecryptFromHexString(output, 0)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		b.Log("TripleDES.Decrypt", string(input))
	}
}

func BenchmarkTripleDES_Base64String(b *testing.B) {
	// b.StopTimer() //调用该函数停止压力测试的时间计数
	//做一些初始化的工作,例如读取文件数据,数据库连接之类的,
	//这样这些时间不影响我们测试函数本身的性能
	// b.StartTimer() //重新开始时间
	for i := 0; i < b.N; i++ {
		TripleDES, err := NewTripleDES([]byte("0123456789abcdef"), nil)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		output := TripleDES.EncryptToBase64String([]byte("我是TripleDES"), 1)
		// if err != nil {
		//  fmt.Println(err.Error())
		//  os.Exit(1)
		// }

		b.Log("TripleDES.Encrypt", output)

		input, err := TripleDES.DecryptFromBase64String(output, 1)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		b.Log("TripleDES.Decrypt", string(input))
	}
}

func TestTripleDES_CBC_HexString(t *testing.T) {
	TripleDES, err := NewTripleDES([]byte("0123456789abcdef"), []byte("01234567"))
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	output := TripleDES.CBCEncryptToHexString([]byte("我是TripleDES_CBC"), 0)

	t.Log("TripleDES.CBCEncrypt", output)

	input, err := TripleDES.CBCDecryptFromHexString(output, 0)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	t.Log("TripleDES.CBCDecrypt", string(input))
}

func TestTripleDES_CBC_Base64String(t *testing.T) {
	TripleDES, err := NewTripleDES([]byte("0123456789abcdef"), []byte("01234567"))
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	output := TripleDES.CBCEncryptToBase64String([]byte("我是TripleDES_CBC"), 1)

	t.Log("TripleDES.CBCEncrypt", output)

	input, err := TripleDES.CBCDecryptFromBase64String(output, 1)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	t.Log("TripleDES.CBCDecrypt", string(input))
}

func BenchmarkTripleDES_CBC_HexString(b *testing.B) {
	// b.StopTimer() //调用该函数停止压力测试的时间计数
	//做一些初始化的工作,例如读取文件数据,数据库连接之类的,
	//这样这些时间不影响我们测试函数本身的性能
	// b.StartTimer() //重新开始时间
	for i := 0; i < b.N; i++ {
		TripleDES, err := NewTripleDES([]byte("0123456789abcdefabcdef0123456789"), []byte("01234567"))
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		output := TripleDES.CBCEncryptToHexString([]byte("我是TripleDES_CBC"), 0)

		b.Log("TripleDES.CBCEncrypt", output)

		input, err := TripleDES.CBCDecryptFromHexString(output, 0)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		b.Log("TripleDES.CBCDecrypt", string(input))
	}
}

func BenchmarkTripleDES_CBC_Base64String(b *testing.B) {
	// b.StopTimer() //调用该函数停止压力测试的时间计数
	//做一些初始化的工作,例如读取文件数据,数据库连接之类的,
	//这样这些时间不影响我们测试函数本身的性能
	// b.StartTimer() //重新开始时间
	for i := 0; i < b.N; i++ {
		TripleDES, err := NewTripleDES([]byte("0123456789abcdefabcdef0123456789"), []byte("01234567"))
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		output := TripleDES.CBCEncryptToBase64String([]byte("我是TripleDES_CBC"), 1)

		b.Log("TripleDES.CBCEncrypt", output)

		input, err := TripleDES.CBCDecryptFromBase64String(output, 1)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		b.Log("TripleDES.CBCDecrypt", string(input))
	}
}
