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

import (
	"strconv"

	"github.com/openingo/godkits/gox/defaultx"
)

// NotEmpty string
func NotEmpty(str string) bool {
	return !Empty(str)
}

// Empty string
func Empty(str string) bool {
	return str == ""
}

// NoneEmpty strings
func NoneEmpty(str ...string) bool {
	return !AnyEmpty(str...)
}

// AnyEmpty strings
func AnyEmpty(str ...string) bool {
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

// DefaultIfEmpty if the string return the defaultStr
func DefaultIfEmpty(str string, defaultStr string) string {
	if Empty(str) {
		return defaultStr
	}
	return str
}

// ToInt convert string int
// wrap strconv.Atoi
func ToInt(str string) int {
	if Empty(str) {
		return 0
	}
	value, err := strconv.Atoi(str)
	return defaultx.DefaultIntIfError(err, value, 0)
}

// ToFloat64 convert string to float64
// wrap strconv.ParseFloat
func ToFloat64(str string) float64 {
	if Empty(str) {
		return 0.0
	}
	value, err := strconv.ParseFloat(str, 64)
	return defaultx.DefaultFloat64IfError(err, value, 0.0)
}

// ToBool convert string to bool
// wrap strconv.ParseBool
func ToBool(str string) bool {
	if Empty(str) {
		return false
	}
	return defaultx.DefaultIfError(strconv.ParseBool(str))
}

// ToComplex convert string to complex
// wrap strconv.ParseComplex
func ToComplex(str string) complex128 {
	if Empty(str) {
		return 0
	}
	value, err := strconv.ParseComplex(str, 128)
	return defaultx.DefaultComplexIfError(err, value, 0)
}

// ToUnit convert string to unit
// If the base argument is 0, the true base is implied by the string's
// prefix following the sign (if present): 2 for "0b", 8 for "0" or "0o",
// 16 for "0x", and 10 otherwise.
// wrap strconv.ParseUint
func ToUnit(str string, base int) uint64 {
	if Empty(str) {
		return 0
	}
	value, err := strconv.ParseUint(str, base, 64)
	return defaultx.DefaultUint64IfError(err, value, 0)
}

// ToBytes concert string to bytes
func ToBytes(str string) []byte {
	if Empty(str) {
		return nil
	}
	return []byte(str)
}
