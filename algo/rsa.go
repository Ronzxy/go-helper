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

package algo

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
)

type RSA struct {
	PublicKeyBytes  []byte
	PrivateKeyBytes []byte

	privateKey *rsa.PrivateKey
	publicKey  *rsa.PublicKey
}

func NewRSA() *RSA {
	return &RSA{}
}

// 公钥加密
func (this *RSA) PublicKeyEncrypt(buf []byte) (string, error) {
	block, _ := pem.Decode(this.PublicKeyBytes)
	if block == nil {
		return "", errors.New("public key error")
	}

	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return "", err
	}

	publicKey := pubInterface.(*rsa.PublicKey)

	cipher, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, buf)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(cipher), nil
}

// 私钥解密
func (this *RSA) PrivateKeyDecrypt(cipher string) ([]byte, error) {
	if this.privateKey == nil {
		err := this.ParsePKCS1PrivateKey()
		if err != nil {
			return nil, err
		}
	}

	data, err := base64.StdEncoding.DecodeString(cipher)
	if err != nil {
		return nil, err
	}

	return rsa.DecryptPKCS1v15(rand.Reader, this.privateKey, data)
}

// // 私钥加密
// func (this *RSA) PrivateKeyEncrypt(buf []byte) (string, error) {
// 	if this.privateKey == nil {
// 		err := this.ParsePKCS1PrivateKey()
// 		if err != nil {
// 			return "", err
// 		}
// 	}

// 	cipher, err := privateKeyEncrypt(rand.Reader, this.privateKey, buf)
// 	if err != nil {
// 		return "", err
// 	}

// 	return base64.StdEncoding.EncodeToString(cipher), nil
// }

// // 公钥解密
// func (this *RSA) PublicKeyDecrypt(cipher string) ([]byte, error) {
// 	if this.publicKey == nil {
// 		err := this.ParsePKCS1PublicKey()
// 		if err != nil {
// 			return nil, err
// 		}
// 	}

// 	data, err := base64.StdEncoding.DecodeString(cipher)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return publicKeyDecrypt(this.publicKey, data)
// }

// 签名
func (this *RSA) Sign(src []byte, hash crypto.Hash) (string, error) {
	if this.privateKey == nil {
		err := this.ParsePKCS1PrivateKey()
		if err != nil {
			return "", err
		}
	}

	h := hash.New()
	h.Write(src)
	hashed := h.Sum(nil)
	data, err := rsa.SignPKCS1v15(rand.Reader, this.privateKey, hash, hashed)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(data), nil
}

// // 验签
// func (this *RSA) Verify(src []byte, s string, hash crypto.Hash) error {
// 	if this.publicKey == nil {
// 		err := this.ParsePKCS1PublicKey()
// 		if err != nil {
// 			return err
// 		}
// 	}

// 	sign, err := base64.StdEncoding.DecodeString(s)
// 	if err != nil {
// 		return err
// 	}

// 	h := hash.New()
// 	h.Write(src)
// 	hashed := h.Sum(nil)
// 	return rsa.VerifyPKCS1v15(this.publicKey, hash, hashed, sign)
// }

func (this *RSA) ParsePKCS1PublicKey() error {
	block, _ := pem.Decode(this.PublicKeyBytes)
	if block == nil {
		return errors.New("public key error")
	}

	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err == nil {
		this.publicKey = pubInterface.(*rsa.PublicKey)
	}

	return err
}

func (this *RSA) ParsePKCS1PrivateKey() error {
	var (
		err error
	)
	block, _ := pem.Decode(this.PrivateKeyBytes)
	if block == nil {
		return errors.New("Parse PrivateKey in Error!")
	}

	this.privateKey, err = x509.ParsePKCS1PrivateKey(block.Bytes)

	return err
}

func (this *RSA) GenerateKeyWithBits(bits int) ([]byte, []byte, error) {
	var (
		privateKeyBytes []byte
		publicKeyBytes  []byte
		err             error
	)

	privateKeyBytes, err = this.GeneratePrivateKey(bits)
	if err != nil {
		return nil, nil, err
	}

	publicKeyBytes, err = this.GeneratePublicKey()
	if err != nil {
		return nil, nil, err
	}

	return privateKeyBytes, publicKeyBytes, nil
}

func (this *RSA) GeneratePrivateKey(bits int) ([]byte, error) {
	var (
		privateKeyBytes []byte
		err             error
	)
	// 生成私钥
	this.privateKey, err = rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return nil, err
	}

	pkcs1 := x509.MarshalPKCS1PrivateKey(this.privateKey)
	block := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: pkcs1,
	}

	// 生成私钥数据
	privateKeyBytes = pem.EncodeToMemory(block)

	return privateKeyBytes, nil
}

func (this *RSA) GeneratePublicKey() ([]byte, error) {
	var (
		publicKeyBytes []byte
	)

	if this.privateKey == nil {
		err := this.ParsePKCS1PrivateKey()
		if err != nil {
			return nil, err
		}
	}

	// 生成公钥
	this.publicKey = &this.privateKey.PublicKey
	pkix, err := x509.MarshalPKIXPublicKey(this.publicKey)
	if err != nil {
		return nil, err
	}

	block := &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: pkix,
	}

	// 生成公钥数据
	publicKeyBytes = pem.EncodeToMemory(block)

	return publicKeyBytes, nil
}
