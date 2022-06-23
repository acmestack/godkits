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

package listx

import "container/list"

type any = interface{}

// Listx extend from list.List
type Listx struct {
	list.List
}

// NotEmpty list
func NotEmpty(lst *list.List) bool {
	return !Empty(lst)
}

// NotEmpty list
func (lst *Listx) NotEmpty() bool {
	return !lst.Empty()
}

// Empty list
func Empty(lst *list.List) bool {
	return lst.Len() == 0
}

// Empty list
func (lst *Listx) Empty() bool {
	return lst.Len() == 0
}

// ForEach iterates over the list, calling fn on each element.
func (lst *Listx) ForEach(fn func(any)) {
	if lst == nil || fn == nil {
		return
	}

	for e := lst.Front(); e != nil; e = e.Next() {
		fn(e.Value)
	}
}
