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
)

// HexEncodeToString convert bytes to string
func HexEncodeToString(src []byte) string {
	return hex.EncodeToString(src)
}

// HexDecodeToBytes build string to hex bytes
func HexDecodeToBytes(src string) ([]byte, error) {
	return hex.DecodeString(src)
}

// HexDump build bytes dump to hex string
// actually, hex dump can do anything
func HexDump(src []byte) string {
	return hex.Dump(src)
}
