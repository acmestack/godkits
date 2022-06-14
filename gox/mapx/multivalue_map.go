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

package mapx

import (
	"github.com/acmestack/godkits/gox/containerx/listx"
)

// MultiValueMap one key mapping multi values
type MultiValueMap struct {
	size int
	kvs  map[any]*listx.Arraylist
}

// NewMultiValueMap multi values map
func NewMultiValueMap(cap int) *MultiValueMap {
	m := &MultiValueMap{}
	m.kvs = make(map[any]*listx.Arraylist, cap)
	m.size = 0
	return m
}

// ContainsKey check the key exist or not
func (m *MultiValueMap) ContainsKey(key any) bool {
	checkMultiValuesMap(m)
	return m.kvs[key] != nil
}

// Put value to mutil value map
func (m *MultiValueMap) Put(key any, value ...any) bool {
	checkMultiValuesMap(m)
	values := m.Values(key)
	if values.Empty() {
		values = listx.New(value...)
		m.kvs[key] = values
		m.size++
	} else {
		values.Add(value...)
	}
	return true
}

// Get values by key
func (m *MultiValueMap) Get(key any) *listx.Arraylist {
	checkMultiValuesMap(m)
	return m.Values(key)
}

// Values by key
func (m *MultiValueMap) Values(key any) *listx.Arraylist {
	checkMultiValuesMap(m)
	return m.kvs[key]
}

// Remove key from mutil value map
func (m *MultiValueMap) Remove(key any) bool {
	checkMultiValuesMap(m)
	delete(m.kvs, key)
	m.size--
	return true
}

// Empty the map empty or not
func (m *MultiValueMap) Empty() bool {
	checkMultiValuesMap(m)
	return m.size == 0
}

// Size mutil value map size
func (m *MultiValueMap) Size() int {
	checkMultiValuesMap(m)
	return m.size
}
