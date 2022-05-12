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

package listx

import (
	"reflect"

	"github.com/openingo/godkits/array"
)

// defaultArraylistSize the default array size 10
const defaultArraylistSize = 10

// Arraylist array lst
type Arraylist struct {
	elements []any
	size     int
}

// New an array list
func New(element ...any) *Arraylist {
	lst := &Arraylist{}
	lst.elements = make([]any, defaultArraylistSize)
	if len(element) > 0 {
		lst.Add(element...)
	}
	return lst
}

// Add element to array list
func (lst *Arraylist) Add(element ...any) {
	if lst.size+len(element) >= len(lst.elements)-1 {
		newElements := make([]any, lst.size+len(element)+1)
		copy(newElements, lst.elements)
		lst.elements = newElements
	}

	for _, value := range element {
		lst.elements[lst.size] = value
		lst.size++
	}
}

// RemoveAtIndex from array list at index
func (lst *Arraylist) RemoveAtIndex(index int) any {
	if index < 0 || index >= lst.size {
		return nil
	}

	current := lst.elements[index]
	lst.elements[index] = nil
	copy(lst.elements[index:], lst.elements[index+1:lst.size])
	lst.size--
	return current
}

// Remove from array list
func (lst *Arraylist) Remove(element any) bool {
	return reflect.DeepEqual(lst.RemoveAtIndex(lst.IndexOf(element)), element)
}

// Get element from array list at index
func (lst *Arraylist) Get(index int) any {
	if index < 0 || index >= lst.size {
		return nil
	}
	return lst.elements[index]
}

// IndexOf element from array list
func (lst *Arraylist) IndexOf(element any) int {
	_, index := lst.find(element)
	return index
}

// Empty array list
func (lst *Arraylist) Empty() bool {
	return lst == nil || lst.Size() == 0 || array.Empty(lst.elements)
}

// Size of array list
func (lst *Arraylist) Size() int {
	return lst.size
}

// Contains element at array list or not
func (lst *Arraylist) Contains(element any) (bool, int) {
	return lst.find(element)
}

func (lst *Arraylist) find(element any) (bool, int) {
	for idx, current := range lst.elements {
		if current == element {
			return true, idx
		}
	}
	return false, -1
}
