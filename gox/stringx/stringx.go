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

package stringx

// NotEmpty strings
func NotEmpty(str ...string) bool {
	return !Empty(str...)
}

// Empty strings
func Empty(str ...string) bool {
	if str == nil || len(str) == 0 {
		return true
	}
	for _, one := range str {
		if one == "" {
			return true
		}
	}
	return false
}

// DefaultIfEmpty the string
func DefaultIfEmpty(str string, defaultStr string) string {
	if Empty(str) {
		return defaultStr
	}
	return str
}
