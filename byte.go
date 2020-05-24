/* Copyright 2015 Ron Zhang <ronzxy@hotmail.com>. All rights reserved.
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
	"strings"
)

type ByteHelper struct{}

func NewByteHelper() *ByteHelper {
	return &ByteHelper{}
}

func (this *ByteHelper) Byte2Hex(buf []byte) (hex []byte) {
	hex = make([]byte, len(buf)*2)
	for i := 0; i < len(buf); i++ {

		v1 := buf[i] / 16
		v2 := buf[i] % 16

		if v1 >= 0 && v1 <= 9 {
			hex[i*2] = 48 + v1
		} else {
			// hex[i*2] = 55 + v1 // upper
			hex[i*2] = 87 + v1 // lower
		}

		if v2 >= 0 && v2 <= 9 {
			hex[i*2+1] = 48 + v2
		} else {
			// hex[i*2+1] = 55 + v2 // upper
			hex[i*2+1] = 87 + v2 // lower
		}

	}

	// for i:=0; i < len(buf); i++ {
	//  hex[i*2] = fmt.Sprintf("%02x", buf[i])
	// }

	return
}

func (this *ByteHelper) Hex2Byte(hex []byte) (buf []byte) {
	buf = make([]byte, len(hex)/2)
	for i := 0; i < len(hex); i++ {
		s1 := strings.ToUpper(string(hex[i]))
		s2 := strings.ToUpper(string(hex[i+1]))

		byte1 := []byte(s1)[0]
		byte2 := []byte(s2)[0]

		var v1, v2 byte
		if byte1 >= 65 {
			v1 = byte1 - 55
		} else {
			v1 = byte1 - 48
		}

		if byte2 >= 65 {
			v2 = byte2 - 55
		} else {
			v2 = byte2 - 48
		}

		i += 1
		buf[(i-1)/2] = v1*16 + v2
	}

	return
}

func (this *ByteHelper) Byte2Bin(buf []byte) (bin []byte) {
	bin = make([]byte, len(buf)*8)
	mask := 1 << (8 - 1)
	for i := 0; i < len(buf); i++ {
		n := int(buf[i])
		for j := 0; j < 8; j++ {
			if (n & mask) == 0 {
				bin[i*8+j] = 0
			} else {
				bin[i*8+j] = 1
			}
			n <<= 1
		}
	}

	return
}

func (this *ByteHelper) Bin2Byte(bin []byte) (buf []byte) {
	if (len(bin) % 8) == 0 {
		buf = make([]byte, len(bin)/8)

		for i := 0; i < len(bin)/8; i++ {
			var b byte
			for j := 0; j < 8; j++ {
				if bin[i*8+j] == 1 {
					var c byte = 1
					for n := 0; n < 7-j; n++ {
						c = c * 2
					}
					b += c
				}
			}
			buf[i] = b
		}
	}

	return
}
