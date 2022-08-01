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

package sha512x

import (
	"crypto/sha512"
	"encoding/hex"
)

// Sha512 Secure Hash Algorithm 2
// @param str string
// return string
func Sha512(str string) string {
	s := sha512.New()
	s.Write([]byte(str))
	return hex.EncodeToString(s.Sum(nil))
}
