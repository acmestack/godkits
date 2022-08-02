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

package bytesx

// SetAll set all bytes of the slice to value
// SetAll set all bytes of the slice to value
//  @param b byte array
//  @param v param
func SetAll(b []byte, v byte) {
	for i := 0; i < len(b); i++ {
		b[i] = v
	}
}

// ZeroAll set all bytes of the slice to 0
//  @param b byte array
func ZeroAll(b []byte) {
	SetAll(b, 0)
}

// FromUint16LE convert uint16 to []byte in little endian
//  @param v param
//  @return []byte result
func FromUint16LE(v uint16) []byte {
	return []byte{
		byte(v),
		byte(v >> 8),
	}
}

// ToUint16LE convert []byte of little endian to uint16
func ToUint16LE(b []byte) uint16 {
	switch len(b) {
	case 0:
		return 0
	case 1:
		return uint16(b[0])
	default:
		return uint16(b[0]) | uint16(b[1])<<8
	}
}

// FromUint16BE convert uint16 to []byte in big endian
func FromUint16BE(v uint16) []byte {
	return []byte{
		byte(v >> 8),
		byte(v),
	}
}

// ToUint16BE convert []byte of big endian to uint16
func ToUint16BE(b []byte) uint16 {
	switch len(b) {
	case 0:
		return 0
	case 1:
		return uint16(b[0])
	default:
		return uint16(b[0])<<8 | uint16(b[1])
	}
}
