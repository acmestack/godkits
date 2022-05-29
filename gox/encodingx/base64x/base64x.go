/*
 * Copyright (c) 2022, OpeningO
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

package base64x

import "encoding/base64"

// Base64EncodeToString convert byte to string
// wrap base64.StdEncoding.EncodeToString
func Base64EncodeToString(input []byte) string {
	return base64.StdEncoding.EncodeToString(input)
}

// Base64DecodeToBytes convert byte to string byte
// warp base64.StdEncoding.DecodeString
func Base64DecodeToBytes(src string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(src)
}

// Base64URLEncodeToString convert url byte to string
// wrap base64.URLEncoding.EncodeToString
func Base64URLEncodeToString(src []byte) string {
	return base64.URLEncoding.EncodeToString(src)
}

// Base64URLDecodeToByte convert string to byte
func Base64URLDecodeToByte(src string) ([]byte, error) {
	return base64.URLEncoding.DecodeString(src)
}

// Base64RawEncodeToString convert raw to string
func Base64RawEncodeToString(src []byte) string {
	return base64.RawStdEncoding.EncodeToString(src)
}

// Base64RawDecodeToByte convert raw string to byte
func Base64RawDecodeToByte(src string) ([]byte, error) {
	return base64.RawStdEncoding.DecodeString(src)
}
