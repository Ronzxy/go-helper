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

package cryptox

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/hex"
	"errors"
)

type AES struct {
	block cipher.Block
	iv    []byte
}

/* creates and returns a new AES. The key argument should be the AES key,
 * either 16, 24, or 32 bytes to select AES-128, AES-192, or AES-256.
 * The length of iv must be the same as the Block's block size.
 *
 * @param key   	The AES key
 * @param iv    	The iv key
 *
 * @return *AES		AES instance
 * @return error	Errors encountered
 *
 */
func NewAES(key, iv []byte) (*AES, error) {
	var (
		this = &AES{}
		err  error
	)

	err = this.AESInit(key, iv)

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
func (this *AES) AESInit(key, iv []byte) error {
	var (
		//this *AES = &AES{}
		err error
	)

	this.block, err = aes.NewCipher(key)
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
func (this *AES) GetBlock() cipher.Block {
	return this.block
}

/* 获取 AES 加密块长度
 *
 * @return int 				AES 加密块长度
 *
 */
func (this *AES) GetBlockSize() int {
	return this.block.BlockSize()
}

/* 加密
 *
 * @param inputBuf 	要加密的明文数据
 * @param p 		数据填充方式
 *
 * @return []byte 	加密后的数据
 *
 */
func (this *AES) Encrypt(inputBuf []byte, p int) []byte {
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

/* 解密
 *
 * @param inputBuf	要解密的密文数据
 * @param p			数据填充方式
 *
 * @return []byte 	解密后的数据
 *
 */
func (this *AES) Decrypt(inputBuf []byte, p int) []byte {
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

/* 加密数据并返回 hex 格式的密文
 *
 * @param inputBuf 	要加密的数据
 * @param p 		数据填充方式
 *
 * @return string 	hex 格式的密文
 *
 */
func (this *AES) EncryptToHexString(inputBuf []byte, p int) string {
	buf := this.Encrypt(inputBuf, p)
	outputStr := hex.EncodeToString(buf)

	return outputStr
}

/* 解密 hex 格式的密文
 *
 * @param inputBuf 	要解密的Hex编码密文字符
 * @param p 		数据填充方式
 *
 * @return []byte   解密后的数据
 * @return error	解密中出现的错误信息
 *
 */
func (this *AES) DecryptFromHexString(inputStr string, p int) ([]byte, error) {
	buf, err := hex.DecodeString(inputStr)
	if err != nil {
		return nil, err
	}
	outputBuf := this.Decrypt(buf, p)

	return outputBuf, nil
}

/* 加密数据并返回 base64 格式的密文
 *
 * @param inputBuf 	要加密的数据
 * @param p 		数据填充方式
 *
 * @return string 	base64 格式的密文
 *
 */
func (this *AES) EncryptToBase64String(inputBuf []byte, p int) string {
	buf := this.Encrypt(inputBuf, p)
	outputStr := base64.StdEncoding.EncodeToString(buf)

	return outputStr
}

/* 解密 base64 格式的密文
 *
 * @param inputBuf 	要解密的 base64 编码密文字符
 * @param p 		数据填充方式
 *
 * @return []byte   解密后的数据
 * @return error	解密中出现的错误信息
 *
 */
func (this *AES) DecryptFromBase64String(inputStr string, p int) ([]byte, error) {
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

/* CBC 加密
 *
 * @param inputBuf 	要加密的明文数据
 * @param p 		数据填充方式
 *
 * @return []byte 	加密后的数据
 *
 */
func (this *AES) CBCEncrypt(inputBuf []byte, p int) []byte {
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

/* CBC 解密
 *
 * @param inputBuf	要解密的密文数据
 * @param p			数据填充方式
 *
 * @return []byte 	解密后的数据
 *
 */
func (this *AES) CBCDecrypt(inputBuf []byte, p int) []byte {
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

/* CBC 加密数据并返回 hex 格式的密文
 *
 * @param inputBuf 	要加密的数据
 * @param p 		数据填充方式
 *
 * @return string 	hex 格式的密文
 *
 */
func (this *AES) CBCEncryptToHexString(inputBuf []byte, p int) string {
	var (
		buf       []byte
		outputStr string
	)

	buf = this.CBCEncrypt(inputBuf, p)
	outputStr = hex.EncodeToString(buf)

	return outputStr
}

/* CBC 解密 hex 格式的密文
 *
 * @param inputBuf 	要解密的Hex编码密文字符
 * @param p 		数据填充方式
 *
 * @return []byte   解密后的数据
 * @return error	解密中出现的错误信息
 *
 */
func (this *AES) CBCDecryptFromHexString(inputStr string, p int) ([]byte, error) {
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

/* CBC 加密数据并返回 base64 格式的密文
 *
 * @param inputBuf 	要加密的数据
 * @param p 		数据填充方式
 *
 * @return string 	base64 格式的密文
 *
 */
func (this *AES) CBCEncryptToBase64String(inputBuf []byte, p int) string {
	var (
		buf       []byte
		outputStr string
	)

	buf = this.CBCEncrypt(inputBuf, p)
	outputStr = base64.StdEncoding.EncodeToString(buf)

	return outputStr
}

/* CBC 解密 base64 格式的密文
 *
 * @param inputBuf 	要解密的 base64 编码密文字符
 * @param p 		数据填充方式
 *
 * @return []byte   解密后的数据
 * @return error	解密中出现的错误信息
 *
 */
func (this *AES) CBCDecryptFromBase64String(inputStr string, p int) ([]byte, error) {
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
