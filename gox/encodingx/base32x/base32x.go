/*
 * Licensed to the AcmeStack under one or more contributor license
 * agreements. See the NOTICE file distributed with this work for
 * additional information regarding copyright ownership.
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

package base32x

import "encoding/base32"

// Base32EncodeToString base32x convert bytes to string
func Base32EncodeToString(src []byte) string {
	return base32.StdEncoding.EncodeToString(src)
}

// Base32DecodeToBytes base32x convert string to bytes
func Base32DecodeToBytes(src string) ([]byte, error) {
	return base32.StdEncoding.DecodeString(src)
}

// Base32HexEncodeToString base32x hex convert bytes to string
func Base32HexEncodeToString(src []byte) string {
	return base32.HexEncoding.EncodeToString(src)
}

// Base32HexDecodeToBytes base32x hex convert string to bytes
func Base32HexDecodeToBytes(src string) ([]byte, error) {
	return base32.HexEncoding.DecodeString(src)
}
