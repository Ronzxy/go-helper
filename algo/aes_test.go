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

func TestAES_HexString(t *testing.T) {
	AES, err := NewAES([]byte("0123456789abcdef"), nil)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	output := AES.EncryptToHexString([]byte("我是AES"), 0)

	t.Log("AES.Encrypt", output)

	input, err := AES.DecryptFromHexString(output, 0)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	t.Log("AES.Decrypt", string(input))
}

func TestAES_Base64String(t *testing.T) {
	AES, err := NewAES([]byte("0123456789abcdef"), nil)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	output := AES.EncryptToBase64String([]byte("我是AES"), 1)

	t.Log("AES.Encrypt", output)

	input, err := AES.DecryptFromBase64String(output, 1)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	t.Log("AES.Decrypt", string(input))
}

func BenchmarkAES_HexString(b *testing.B) {
	// b.StopTimer() //调用该函数停止压力测试的时间计数
	//做一些初始化的工作,例如读取文件数据,数据库连接之类的,
	//这样这些时间不影响我们测试函数本身的性能
	// b.StartTimer() //重新开始时间
	for i := 0; i < b.N; i++ {
		AES, err := NewAES([]byte("0123456789abcdef"), nil)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		output := AES.EncryptToHexString([]byte("我是AES"), 0)
		// if err != nil {
		//  fmt.Println(err.Error())
		//  os.Exit(1)
		// }

		b.Log("AES.Encrypt", output)

		input, err := AES.DecryptFromHexString(output, 0)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		b.Log("AES.Decrypt", string(input))
	}
}

func BenchmarkAES_Base64String(b *testing.B) {
	// b.StopTimer() //调用该函数停止压力测试的时间计数
	//做一些初始化的工作,例如读取文件数据,数据库连接之类的,
	//这样这些时间不影响我们测试函数本身的性能
	// b.StartTimer() //重新开始时间
	for i := 0; i < b.N; i++ {
		AES, err := NewAES([]byte("0123456789abcdef"), nil)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		output := AES.EncryptToBase64String([]byte("我是AES"), 1)
		// if err != nil {
		//  fmt.Println(err.Error())
		//  os.Exit(1)
		// }

		b.Log("AES.Encrypt", output)

		input, err := AES.DecryptFromBase64String(output, 1)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		b.Log("AES.Decrypt", string(input))
	}
}

func TestAES_CBC_HexString(t *testing.T) {
	AES, err := NewAES([]byte("0123456789abcdef"), []byte("0123456789abcdef"))
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	output := AES.CBCEncryptToHexString([]byte("我是AES_CBC"), 0)

	t.Log("AES.CBCEncrypt", output)

	input, err := AES.CBCDecryptFromHexString(output, 0)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	t.Log("AES.CBCDecrypt", string(input))
}

func TestAES_CBC_Base64String(t *testing.T) {
	AES, err := NewAES([]byte("0123456789abcdef"), []byte("0123456789abcdef"))
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	output := AES.CBCEncryptToBase64String([]byte("我是AES_CBC"), 1)

	t.Log("AES.CBCEncrypt", output)

	input, err := AES.CBCDecryptFromBase64String(output, 1)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	t.Log("AES.CBCDecrypt", string(input))
}

func BenchmarkAES_CBC_HexString(b *testing.B) {
	// b.StopTimer() //调用该函数停止压力测试的时间计数
	//做一些初始化的工作,例如读取文件数据,数据库连接之类的,
	//这样这些时间不影响我们测试函数本身的性能
	// b.StartTimer() //重新开始时间
	for i := 0; i < b.N; i++ {
		AES, err := NewAES([]byte("0123456789abcdefabcdef0123456789"), []byte("0123456789abcdef"))
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		output := AES.CBCEncryptToHexString([]byte("我是AES_CBC"), 0)

		b.Log("AES.CBCEncrypt", output)

		input, err := AES.CBCDecryptFromHexString(output, 0)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		b.Log("AES.CBCDecrypt", string(input))
	}
}

func BenchmarkAES_CBC_Base64String(b *testing.B) {
	// b.StopTimer() //调用该函数停止压力测试的时间计数
	//做一些初始化的工作,例如读取文件数据,数据库连接之类的,
	//这样这些时间不影响我们测试函数本身的性能
	// b.StartTimer() //重新开始时间
	for i := 0; i < b.N; i++ {
		AES, err := NewAES([]byte("0123456789abcdefabcdef0123456789"), []byte("0123456789abcdef"))
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		output := AES.CBCEncryptToBase64String([]byte("我是AES_CBC"), 1)

		b.Log("AES.CBCEncrypt", output)

		input, err := AES.CBCDecryptFromBase64String(output, 1)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		b.Log("AES.CBCDecrypt", string(input))
	}
}
