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

package stringsx

import (
	"strconv"
	"strings"

	"github.com/acmestack/godkits/gox/defaultx"
)

// NotEmpty string
func NotEmpty(str string) bool {
	return !Empty(str)
}

// Empty string
func Empty(str string) bool {
	return str == "\n" || str == "\t" || str == "" || len(Trim(str)) == 0
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
		if Empty(one) {
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
	return ToIntOrDefault(str, 0)
}

// ToIntOrDefault convert string int, if str empty or err you can get default value tha you want.
// wrap strconv.Atoi
func ToIntOrDefault(str string, defVal int) int {
	if Empty(str) {
		return defVal
	}
	value, err := strconv.Atoi(str)
	return defaultx.DefaultIntIfError(err, value, defVal)
}

// ToFloat64 convert string to float64
// wrap strconv.ParseFloat
func ToFloat64(str string) float64 {
	return ToFloat64OrDefault(str, 0.0)
}

// ToFloat64OrDefault convert string to float64, if str empty or err you can get default value tha you want.
// wrap strconv.ParseFloat
func ToFloat64OrDefault(str string, defVal float64) float64 {
	if Empty(str) {
		return defVal
	}
	value, err := strconv.ParseFloat(str, 64)
	return defaultx.DefaultFloat64IfError(err, value, defVal)
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

// Trim the " " cutset
func Trim(str string) string {
	return strings.Trim(str, " ")
}

// ReplaceAll Replace all old in str with new
func ReplaceAll(str string, old string, new string) string {
	if AnyEmpty(str, old, new) {
		return str
	}
	return strings.ReplaceAll(str, old, new)
}

// StripBlank Deletes leading and trailing whitespace characters
/*
 * <p>Strips whitespace from the start and end of a String.</p>
 *
 * <p>This is similar to {@link Trim(String)} but removes whitespace.
 * Whitespace is defined by {@link Empty(char)}.</p>
 *
 * <pre>
 * StripBlank("")       = ""
 * StripBlank("   ")    = ""
 * StripBlank("abc")    = "abc"
 * StripBlank("  abc")  = "abc"
 * StripBlank("abc  ")  = "abc"
 * StripBlank(" abc ")  = "abc"
 * StripBlank(" ab c ") = "ab c"
 * </pre>
 */
func StripBlank(str string) string {
	return Strip(str, "")
}

// Strip Delete the characters that start and end with stripChars
/*
 * <p>Strips any of a set of characters from the start and end of a String.
 * This is similar to {@link String#trim()} but allows the characters
 * to be stripped to be controlled.</p>
 *
 * <p>A {@code  EmptyStr} input String returns {@code  EmptyStr}.
 * An empty string ("") input returns the empty string.</p>
 *
 * <p>If the stripChars String is {@code  blank}, whitespace is
 * stripped as defined by {@link Empty(char)}
 *
 * <pre>
 * Strip("", *)            = ""
 * Strip("abc", "")      = "abc"
 * Strip("  abc", "")    = "abc"
 * Strip("abc  ", "")    = "abc"
 * Strip(" abc ", "")    = "abc"
 * Strip("  abcyx", "xyz") = "  abc"
 * </pre>
 */
func Strip(str string, stripChars string) string {
	if Empty(str) {
		return str
	}
	str = StripStart(str, stripChars)
	return StripEnd(str, stripChars)
}

// StripStart Delete the characters that start with stripChars
/*
 * <p>Strips any of a set of characters from the start of a String.</p>
 *
 * <p>A {@code EmptyStr} input String returns {@code EmptyStr}.
 * An empty string ("") input returns the empty string.</p>
 *
 * <p>If the stripChars String is {@code EmptyStr}, whitespace is
 * stripped as defined by {@link Empty(char)}.</p>
 *
 * <pre>
 * StripStart("", *)            = ""
 * StripStart("abc", "")        = "abc"
 * StripStart("abc", "")      = "abc"
 * StripStart("  abc", "")    = "abc"
 * StripStart("abc  ", "")    = "abc  "
 * StripStart(" abc ", "")    = "abc "
 * StripStart("yxabc  ", "xyz") = "abc  "
 * </pre>
 */
func StripStart(str string, stripChars string) string {
	if Empty(str) {
		return str
	}
	start := 0
	strlen := len(str)
	if Empty(stripChars) {
		for start != strlen && Empty(string(str[start])) {
			start++
		}
	} else {
		for start != strlen && strings.Contains(stripChars, string(str[start])) {
			start++
		}
	}
	return str[start:]
}

// StripEnd Delete the characters that end with stripChars
/*
 * <p>Strips any of a set of characters from the end of a String.</p>
 *
 * <p>A {@code EmptyStr} input String returns {@code EmptyStr}.
 * An empty string ("") input returns the empty string.</p>
 *
 * <p>If the stripChars String is {@code EmptyStr}, whitespace is
 * stripped as defined by {@link Empty(char)}.</p>
 *
 * <pre>
 * StripEnd("", *)            = ""
 * StripEnd("abc", "")        = "abc"
 * StripEnd("abc", null)      = "abc"
 * StripEnd("  abc", null)    = "  abc"
 * StripEnd("abc  ", null)    = "abc"
 * StripEnd(" abc ", null)    = " abc"
 * StripEnd("  abcyx", "xyz") = "  abc"
 * StripEnd("120.00", ".0")   = "12"
 * </pre>
 */
func StripEnd(str string, stripChars string) string {
	if Empty(str) {
		return str
	}
	end := len(str)
	if Empty(stripChars) {
		for end != 0 && Empty(string(str[end-1])) {
			end--
		}
	} else {
		for end != 0 && strings.Contains(stripChars, string(str[end-1])) {
			end--
		}
	}
	return str[:end]
}
