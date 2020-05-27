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

package algo

import (
	"crypto/cipher"
	"crypto/des"
	"encoding/base64"
	"encoding/hex"
	"errors"
)

type DES struct {
	block cipher.Block
	iv    []byte
}

/* creates and returns a new DES. The key argument should be the 8 bytes DES key.
 * The length of iv must be the same as the Block's block size.
 *
 * @param key   the AES key
 * @param iv    the iv key
 *
 * @return *AES		AES instance
 * @return error	Errors encountered
 */
func NewDES(key, iv []byte) (*DES, error) {
	var (
		this *DES = &DES{}
		err  error
	)

	err = this.DESInit(key, iv)

	return this, err
}

/* 初始化 DES 实例，密钥必须 16 字节
 *
 * @param key 		DES 加解密密钥
 * @param iv 		8 字节 CBC 初始化向量
 *
 * @return error	遇到的错误
 *
 */
func (this *DES) DESInit(key, iv []byte) error {
	var (
		err error
	)

	this.block, err = des.NewCipher(key)
	if err != nil {
		return err
	}

	if len(iv) > 0 {
		if len(iv) != this.block.BlockSize() {
			return errors.New("IV length must equal block size")
		}

		this.iv = iv
	}

	return nil
}

/* 获取 AES 使用的 cipher.Block
 *
 * @return cipher.Block 	AES使用的block
 *
 */
func (this *DES) GetBlock() cipher.Block {
	return this.block
}

/* 获取 DES 加密块长度
 *
 * @return int 				DES 加密块长度
 *
 */
func (this *DES) GetBlockSize() int {
	return this.block.BlockSize()
}

/* 加密明文
 * @param inputBuf 	要加密的明文数据
 * @param p 		数据填充方式
 *
 * @return []byte 明文
 */
func (this *DES) Encrypt(inputBuf []byte, p int) []byte {
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
func (this *DES) Decrypt(inputBuf []byte, p int) []byte {
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
func (this *DES) EncryptToHexString(inputBuf []byte, p int) string {
	buf := this.Encrypt(inputBuf, p)
	outputStr := hex.EncodeToString(buf)

	return outputStr
}

/** 解密密文
 * @param inputStr 	要解密的Hex编码密文字符
 * @param p 		数据填充方式 */
func (this *DES) DecryptFromHexString(inputStr string, p int) ([]byte, error) {
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
func (this *DES) EncryptToBase64String(inputBuf []byte, p int) string {
	buf := this.Encrypt(inputBuf, p)
	outputStr := base64.StdEncoding.EncodeToString(buf)

	return outputStr
}

/** 解密密文
 * @param inputStr 	要解密的Base64编码密文字符
 * @param p 		数据填充方式 */
func (this *DES) DecryptFromBase64String(inputStr string, p int) ([]byte, error) {
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
func (this *DES) CBCEncrypt(inputBuf []byte, p int) []byte {
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
func (this *DES) CBCDecrypt(inputBuf []byte, p int) []byte {
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
func (this *DES) CBCEncryptToHexString(inputBuf []byte, p int) string {
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
func (this *DES) CBCDecryptFromHexString(inputStr string, p int) ([]byte, error) {
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
func (this *DES) CBCEncryptToBase64String(inputBuf []byte, p int) string {
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
func (this *DES) CBCDecryptFromBase64String(inputStr string, p int) ([]byte, error) {
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
