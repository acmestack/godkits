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

package ascii85x

import (
	"encoding/ascii85"
)

// Ascii85EncodeToInt ascii85x convert bytes to int
func Ascii85EncodeToInt(dst, src []byte) int {
	return ascii85.Encode(dst, src)
}

// Ascii85Decode ascii85x convert bytes and the number  to number of bytes
func Ascii85Decode(dst, src []byte, flush bool) (ndst, nsrc int, err error) {
	return ascii85.Decode(dst, src, flush)
}

// Ascii85MaxEncodedLenToInt ascii85x convert int to int
func Ascii85MaxEncodedLenToInt(n int) int {
	return ascii85.MaxEncodedLen(n)
}
