/*
 * Copyright (c) 2022, AcmeStack
 * All rights reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package aesx

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
)

// EncryptString0 the origin data with the string key
func EncryptString0(str string, strKey string) ([]byte, error) {
	return Encrypt([]byte(str), bytesKey(strKey))
}

// EncryptString the origin data with the key
func EncryptString(str string, key []byte) ([]byte, error) {
	return Encrypt([]byte(str), key)
}

// Encrypt0 the origin data with the string key
func Encrypt0(originData []byte, strKey string) ([]byte, error) {
	return Encrypt(originData, bytesKey(strKey))
}

// Encrypt the origin data with the key
func Encrypt(originData []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	blockSize := block.BlockSize()
	originData = pkcs5Padding(originData, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize])
	encrypted := make([]byte, len(originData))
	blockMode.CryptBlocks(encrypted, originData)
	return encrypted, nil
}

// Decrypt0 the encrypted data with the string key
func Decrypt0(encrypted []byte, strKey string) ([]byte, error) {
	return Decrypt(encrypted, bytesKey(strKey))
}

// Decrypt the encrypted data with the key
func Decrypt(encrypted []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
	origData := make([]byte, len(encrypted))
	blockMode.CryptBlocks(origData, encrypted)
	origData = pkcs5UnPadding(origData)
	return origData, nil
}

func bytesKey(strKey string) []byte {
	return []byte(strKey)
}

func pkcs5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padText...)
}

func pkcs5UnPadding(origData []byte) []byte {
	length := len(origData)
	unPadding := int(origData[length-1])
	return origData[:(length - unPadding)]
}
