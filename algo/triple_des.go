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
	"crypto/cipher"
	"crypto/des"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
)

type TripleDES struct {
	block cipher.Block
	iv    []byte
}

/* creates and returns a new TripleDES. The key argument should be the TripleDES key,
 * either 8, 16, or 24 bytes to select TripleDES with 1 key, TripleDES with 2 key,
 * or TripleDES with 3 key. The length of iv must be the same as the Block's block size.
 * If an 8 bytes key, and it is the same as DES.
 *
 * @param key   the TripleDES key
 * @param iv    the iv key
 *
 * @return *TripleDES   TripleDES instance
 * @return error        Errors encountered
 */
func NewTripleDES(key, iv []byte) (*TripleDES, error) {
	var (
		this *TripleDES = &TripleDES{}
		err  error
	)

	err = this.TripleDESInit(key, iv)

	return this, err
}

/* 初始化 AES 实例，16, 24, 或 32 字节来选择 AES-128, AES-192, 或 AES-256
 *
 * @param key 		AES 加解密密钥
 * @param iv 		16 字节 CBC 初始化向量
 *
 * @return error	遇到的错误
 *
 */
func (this *TripleDES) TripleDESInit(key, iv []byte) error {
	var (
		err error
	)

	switch {
	case len(key) >= 24:
		{
			key = key[:24]
		}
	case len(key) >= 16:
		{
			key = append(key[:16], key[:8]...)
		}
	case len(key) >= 8:
		{
			key = append(key[:8], key[:8]...)
			key = append(key, key[:8]...)
		}
	default:
		{
			return errors.New(fmt.Sprintf("crypto/des: invalid key size %d, you can set key size with 8, 16, 24.",
				len(key)))
		}
	}

	this.block, err = des.NewTripleDESCipher(key)
	if err != nil {
		return err
	}

	if len(iv) > 0 {
		if len(iv) != this.block.BlockSize() {
			return errors.New(fmt.Sprintf("IV length %d must equal block size %d",
				len(iv),
				this.block.BlockSize()))
		}

		this.iv = iv
	}

	return nil
}

/* 获取 TripleDES 使用的 cipher.Block
 *
 * @return cipher.Block 	TripleDES使用的block
 *
 */
func (this *TripleDES) GetBlock() cipher.Block {
	return this.block
}

/* 获取 TripleDES 加密块长度
 *
 * @return int 				TripleDES 加密块长度
 *
 */
func (this *TripleDES) GetBlockSize() int {
	return this.block.BlockSize()
}

/* 加密明文
 * @param inputBuf 	要加密的明文数据
 * @param p 		数据填充方式
 *
 * @return []byte 明文
 */
func (this *TripleDES) Encrypt(inputBuf []byte, p int) []byte {
	var (
		buf       []byte
		outputBuf []byte
		blockSize int
	)

	blockSize = this.block.BlockSize()

	switch p {
	case PADDING_ZERO:
		inputBuf = ZeroPadding(inputBuf, blockSize)
	case PADDING_PKCS5:
		inputBuf = PKCS5Padding(inputBuf, blockSize)
	default:
		inputBuf = ZeroPadding(inputBuf, blockSize)
	}

	for i := 0; i < len(inputBuf)/blockSize; i++ {
		p := i * blockSize
		buf = make([]byte, blockSize)
		this.block.Encrypt(buf, inputBuf[p:p+blockSize])
		outputBuf = append(outputBuf, buf...)
	}

	return outputBuf
}

/* 解密密文
 * @param inputBuf	要解密的密文数据
 * @param p			数据填充方式
 *
 * @return []byte 明文
 */
func (this *TripleDES) Decrypt(inputBuf []byte, p int) []byte {
	var (
		length    int
		buf       []byte
		outputBuf []byte
		blockSize int
	)

	blockSize = this.block.BlockSize()

	// 按块计算长度
	length = blockSize * ((len(inputBuf) + blockSize - 1) / blockSize)

	if length != len(inputBuf) {
		return nil
	}

	for i := 0; i < len(inputBuf)/blockSize; i++ {
		p := i * blockSize
		buf = make([]byte, blockSize)
		this.block.Decrypt(buf, inputBuf[p:p+blockSize])
		outputBuf = append(outputBuf, buf...)
	}

	switch p {
	case PADDING_ZERO:
		outputBuf = UnZeroPadding(outputBuf)
	case PADDING_PKCS5:
		outputBuf = UnPKCS5Padding(outputBuf)
	default:
		outputBuf = UnZeroPadding(outputBuf)
	}

	return outputBuf
}

/** 解密字符密文
 * @param inputBuf 	要加密的数据
 * @param p 		数据填充方式 */
func (this *TripleDES) EncryptToHexString(inputBuf []byte, p int) string {
	buf := this.Encrypt(inputBuf, p)
	outputStr := hex.EncodeToString(buf)

	return outputStr
}

/** 解密密文
 * @param inputStr 	要解密的Hex编码密文字符
 * @param p 		数据填充方式 */
func (this *TripleDES) DecryptFromHexString(inputStr string, p int) ([]byte, error) {
	buf, err := hex.DecodeString(inputStr)
	if err != nil {
		return nil, err
	}
	outputBuf := this.Decrypt(buf, p)

	return outputBuf, nil
}

/** 解密字符密文
 * @param inputBuf 要加密的数据
 * @param p 	数据填充方式 */
func (this *TripleDES) EncryptToBase64String(inputBuf []byte, p int) string {
	buf := this.Encrypt(inputBuf, p)
	outputStr := base64.StdEncoding.EncodeToString(buf)

	return outputStr
}

/** 解密密文
 * @param inputStr 	要解密的Base64编码密文字符
 * @param p 		数据填充方式 */
func (this *TripleDES) DecryptFromBase64String(inputStr string, p int) ([]byte, error) {
	buf, err := base64.StdEncoding.DecodeString(inputStr)
	if err != nil {
		return nil, err
	}
	outputBuf := this.Decrypt(buf, p)

	return outputBuf, nil
}

//
// AES_CBC
//

/** 解密密文
 * @param inputBuf 	要解密的密文数据
 * @param p			数据填充方式
 *
 * @return []byte 密文
 */
func (this *TripleDES) CBCEncrypt(inputBuf []byte, p int) []byte {
	var (
		outputBuf []byte
	)
	if len(this.iv) > 0 {
		blockSize := this.block.BlockSize()

		switch p {
		case PADDING_ZERO:
			inputBuf = ZeroPadding(inputBuf, blockSize)
		case PADDING_PKCS5:
			inputBuf = PKCS5Padding(inputBuf, blockSize)
		default:
			inputBuf = ZeroPadding(inputBuf, blockSize)
		}

		outputBuf = make([]byte, len(inputBuf))
		blockMode := cipher.NewCBCEncrypter(this.block, this.iv)
		blockMode.CryptBlocks(outputBuf, inputBuf)
	}

	return outputBuf
}

/* 解密密文
 * @param inputBuf	要解密的密文数据
 * @param p			数据填充方式
 *
 * @return []byte 明文
 */
func (this *TripleDES) CBCDecrypt(inputBuf []byte, p int) []byte {
	var (
		outputBuf []byte
	)
	if len(this.iv) > 0 {
		outputBuf = make([]byte, len(inputBuf))
		blockMode := cipher.NewCBCDecrypter(this.block, this.iv)
		blockMode.CryptBlocks(outputBuf, inputBuf)
		switch p {
		case PADDING_ZERO:
			outputBuf = UnZeroPadding(outputBuf)
		case PADDING_PKCS5:
			outputBuf = UnPKCS5Padding(outputBuf)
		default:
			outputBuf = UnZeroPadding(outputBuf)
		}
	}

	return outputBuf
}

/** 解密字符密文
 * @param inputBuf 	要加密的数据
 * @param p 		数据填充方式 */
func (this *TripleDES) CBCEncryptToHexString(inputBuf []byte, p int) string {
	var (
		buf       []byte
		outputStr string
	)

	buf = this.CBCEncrypt(inputBuf, p)
	outputStr = hex.EncodeToString(buf)

	return outputStr
}

/** 解密密文
 * @param inputStr 	要解密的Hex编码密文数据
 * @param p 		数据填充方式 */
func (this *TripleDES) CBCDecryptFromHexString(inputStr string, p int) ([]byte, error) {
	var (
		buf       []byte
		outputBuf []byte
		err       error
	)

	buf, err = hex.DecodeString(inputStr)
	if err != nil {
		return nil, err
	}

	outputBuf = this.CBCDecrypt(buf, p)

	return outputBuf, nil
}

/** 解密字符密文
 * @param inputBuf 	要加密的数据
 * @param p 		数据填充方式 */
func (this *TripleDES) CBCEncryptToBase64String(inputBuf []byte, p int) string {
	var (
		buf       []byte
		outputStr string
	)

	buf = this.CBCEncrypt(inputBuf, p)
	outputStr = base64.StdEncoding.EncodeToString(buf)

	return outputStr
}

/** 解密密文
 * @param inputStr 	要解密的Base64编码密文数据
 * @param p 		数据填充方式 */
func (this *TripleDES) CBCDecryptFromBase64String(inputStr string, p int) ([]byte, error) {
	var (
		buf       []byte
		outputBuf []byte
		err       error
	)

	buf, err = base64.StdEncoding.DecodeString(inputStr)
	if err != nil {
		return nil, err
	}
	outputBuf = this.CBCDecrypt(buf, p)

	return outputBuf, nil
}
