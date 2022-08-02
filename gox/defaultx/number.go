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

package defaultx

import "log"

// DefaultComplexIfError if error use defaultValue
func DefaultComplexIfError(err error, value complex128, defaultValue complex128) complex128 {
	if err != nil {
		log.Println(err)
		return defaultValue
	}
	return value
}

// DefaultFloat64IfError if error use defaultValue
func DefaultFloat64IfError(err error, value float64, defaultValue float64) float64 {
	if err != nil {
		log.Println(err)
		return defaultValue
	}
	return value
}

// DefaultIntIfError if error use defaultValue
func DefaultIntIfError(err error, value int, defaultValue int) int {
	if err != nil {
		log.Println(err)
		return defaultValue
	}
	return value
}

// DefaultUint64IfError if error use defaultValue
func DefaultUint64IfError(err error, value uint64, defaultValue uint64) uint64 {
	if err != nil {
		log.Println(err)
		return defaultValue
	}
	return value
}
