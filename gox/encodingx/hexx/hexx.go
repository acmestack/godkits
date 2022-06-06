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

package hexx

import (
	"encoding/hex"
	"github.com/openingo/godkits/gox/stringsx"
)

// HexUpperTable upper case hex table
const HexUpperTable = "0123456789ABCDEF"

// HexLowerTable lower case hex table
const HexLowerTable = "0123456789abcdef"

// EncodeBytesToString convert bytes to string
func EncodeBytesToString(src []byte) string {
	return hex.EncodeToString(src)
}

// EncodeToBytes base method convert byte array encode to byte array
/*
 * <pre>
 * dst := new byte("I am moremind")
 * src := new byte("4920616d206d6f72656d696e64")
 * hextable := "0123456789abcdef"
 * EncodeToBytes(dst, src, hextable) = len(src) * 2 = 26
 * </pre>
 *
 */
func EncodeToBytes(dst, src []byte, hextable string) int {
	j := 0
	for _, v := range src {
		dst[j] = hextable[v>>4]
		dst[j+1] = hextable[v&0x0f]
		j += 2
	}
	return len(src) * 2
}

// EncodeToCaseBytes base method, encode bytes to upper or lower case bytes
/*
 * <pre>
 * src := []byte("nd");
 * EncodeToCaseBytes(src, true)  = []byte("6e64")
 * EncodeToCaseBytes(src, false) = []byte("6E64")
 * </pre>
 *
 */
func EncodeToCaseBytes(src []byte, toLowerCase bool) []byte {
	dst := make([]byte, hex.EncodedLen(len(src)))
	if toLowerCase {
		EncodeToBytes(dst, src, HexLowerTable)
	} else {
		EncodeToBytes(dst, src, HexUpperTable)
	}
	return dst
}

// EncodeToCaseString encode byte to upper or lower case string
func EncodeToCaseString(src []byte, toLowerCase bool) string {
	// TODO optimize use Pointer generate string
	return string(EncodeToCaseBytes(src, toLowerCase))
}

// EncodeToHexString encode src to upper or lower case string
func EncodeToHexString(src string, toLowerCase bool) string {
	// TODO optimize,use Pointer generate bytes
	dst := stringsx.ToBytes(src)
	return EncodeToCaseString(dst, toLowerCase)
}

// HexDecodeToBytes build string to hex bytes
func HexDecodeToBytes(src string) ([]byte, error) {
	return hex.DecodeString(src)
}
