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

import (
	"reflect"

	"github.com/acmestack/godkits/array"
)

// defaultArraylistSize the default array size 10
const defaultArraylistSize = 10

// Arraylist array lst
type Arraylist struct {
	elements []Any
	size     int
}

// New create an array list
//  @param element array element
//  @return *Arraylist
func New(element ...Any) *Arraylist {
	lst := &Arraylist{}
	lst.elements = make([]Any, defaultArraylistSize)
	if len(element) > 0 {
		lst.Add(element...)
	}
	return lst
}

// Add save element to array list
//  @receiver lst lst
//  @param element array element
func (lst *Arraylist) Add(element ...Any) {
	if lst.size+len(element) >= len(lst.elements)-1 {
		newElements := make([]Any, lst.size+len(element)+1)
		copy(newElements, lst.elements)
		lst.elements = newElements
	}

	for _, value := range element {
		lst.elements[lst.size] = value
		lst.size++
	}
}

// RemoveAtIndex from array list at index
// RemoveAtIndex remove from array list at index
//  @receiver lst lst
//  @param index array index
//  @return any
func (lst *Arraylist) RemoveAtIndex(index int) Any {
	if index < 0 || index >= lst.size {
		return nil
	}

	current := lst.elements[index]
	lst.elements[index] = nil
	copy(lst.elements[index:], lst.elements[index+1:lst.size])
	lst.size--
	return current
}

// Remove delete from array list
//  @receiver lst
//  @param element element
//  @return bool true or false
func (lst *Arraylist) Remove(element Any) bool {
	return reflect.DeepEqual(lst.RemoveAtIndex(lst.IndexOf(element)), element)
}

// Get get element from array list at index
//  @receiver lst
//  @param index element index
//  @return Any current result
func (lst *Arraylist) Get(index int) Any {
	if index < 0 || index >= lst.size {
		return nil
	}
	return lst.elements[index]
}

// IndexOf index element from array list
//  @receiver lst
//  @param element element
//  @return int current element index
func (lst *Arraylist) IndexOf(element Any) int {
	_, index := lst.find(element)
	return index
}

// Empty judge array is empty
//  @receiver lst
//  @return bool true => empty, false => not empty
func (lst *Arraylist) Empty() bool {
	return lst == nil || lst.Size() == 0 || array.Empty(lst.elements)
}

// Size calc array length
//  @receiver lst
//  @return int array length
func (lst *Arraylist) Size() int {
	return lst.size
}

// Contains judge array exist element, and return exist result and index
//  @receiver lst
//  @param element element
//  @return bool   exist result
//  @return int index
func (lst *Arraylist) Contains(element Any) (bool, int) {
	return lst.find(element)
}

// find query element in array, and return exist and index
//  @receiver lst
//  @param element element
//  @return bool true => exist, false => not exist
//  @return int  current element index
func (lst *Arraylist) find(element Any) (bool, int) {
	for idx, current := range lst.elements {
		if current == element {
			return true, idx
		}
	}
	return false, -1
}
