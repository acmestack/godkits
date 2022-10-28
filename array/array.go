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

package array

type Any = interface{}

// NotEmpty judge array is not empty
//  @param arr  array
//  @return bool not empty => true, empty => false
func NotEmpty(arr ...Any) bool {
	return !Empty(arr)
}

// Empty judge array is empty
//  @param arr   array
//  @return bool empty => true, not empty => false
func Empty(arr ...Any) bool {
	return arr == nil || len(arr) == 0
}
