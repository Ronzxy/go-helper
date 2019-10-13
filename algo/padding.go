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

package algo

import (
	"bytes"
)

func ZeroPadding(buf []byte, blockSize int) []byte {
	count := blockSize - len(buf)%blockSize

	return append(buf, bytes.Repeat([]byte{0}, count)...)
}

func UnZeroPadding(buf []byte) []byte {
	length := len(buf)

	for i := 0; i < len(buf); i++ {
		if buf[length-1-i] != byte(0) {
			length = length - i
			break
		}
	}

	return buf[:length]
}

// PKCS5是PKCS7的子集？
// 1、如果输入的数据不足16个字节，需要补齐（填充）到16个字节，比如：10个字节就补齐6个6,11个字节就补齐5个5，以此类推；
// 2、如果输入的数据是16的整数倍个字节，需要在数据后面填充16个0×10（也可以是10进制的16）；
// 3、如果输入的数据大于16且不是16的倍数，你需要把字符串补齐到16位，比如：如果少4位，就补充4个4， 如果少5位就补充5个5，少n位，补充n个n。
//
// 需要补齐（填充）到16整数倍个字节，填充1个字符就全0x01, 填充2个字符就全0x02, 填充3个字符就全0x03
// 不需要补齐时需要增加一个块，填充块长度，块长为8就填充0x08（10进制的8），块长为16就填充0x10（10进制的16）

/* PKCS5 is a subset of PKCS7. Handle the bytes creates and returns a new TripleDES. The key argument should be the TripleDES key,
 * either 8, 16, or 24 bytes to select TripleDES with 1 key, TripleDES with 2 key,
 * or TripleDES with 3 key. The length of iv must be the same as the Block's block size.
 * If an 8 bytes key, and it is the same as DES.
 *
 * @param text   the TripleDES key
 * @param iv    the iv key
 *
 * @return *TripleDES   TripleDES instance
 * @return error        Errors encountered
 */
func PKCS5Padding(buf []byte, blockSize int) []byte {
	padding := blockSize - len(buf)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(buf, padText...)
}

func UnPKCS5Padding(buf []byte) []byte {
	length := len(buf)
	// 去掉最后一个字节 unpadding 次
	unpadding := int(buf[length-1])
	return buf[:(length - unpadding)]
}

func PKCS7Padding(buf []byte, blockSize int) (padded []byte) {
	// block size must be bigger or equal 2
	if blockSize < 1<<1 {
		panic("block size is too small (minimum is 2 bytes)")
	}
	// block size up to 255 requires 1 byte padding
	if blockSize < 1<<8 {
		// calculate padding length
		slice_length := len(buf)
		padlen := blockSize - slice_length%blockSize
		if padlen == 0 {
			padlen = blockSize
		}

		// define PKCS7 padding block
		padding := bytes.Repeat([]byte{byte(padlen)}, padlen)

		// apply padding
		padded = append(buf, padding...)
		return padded
	}
	// block size bigger or equal 256 is not currently supported
	panic("unsupported block size")
}
