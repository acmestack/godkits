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

package listx

import "container/list"

type Any = interface{}

// Listx extend from list.List
type Listx struct {
	list.List
}

// NotEmpty judge list is not empty
//  @param lst list
//  @return bool true => not empty, false => empty
func NotEmpty(lst *list.List) bool {
	return !Empty(lst)
}

// NotEmpty judge list is not empty
//  @receiver lst list
//  @return bool true => not empty, false => empty
func (lst *Listx) NotEmpty() bool {
	return !lst.Empty()
}

// Empty judge list is empty
//  @param lst list
//  @return bool true => empty, false => not empty
func Empty(lst *list.List) bool {
	return lst.Len() == 0
}

// Empty judge list is empty
//  @receiver lst list
//  @return bool true => empty, false => not empty
func (lst *Listx) Empty() bool {
	return lst.Len() == 0
}

// ForEach iterates over the list, calling fn on each element.
//  @receiver lst list
//  @param fn function
func (lst *Listx) ForEach(fn func(Any)) {
	if lst == nil || fn == nil {
		return
	}

	for e := lst.Front(); e != nil; e = e.Next() {
		fn(e.Value)
	}
}
