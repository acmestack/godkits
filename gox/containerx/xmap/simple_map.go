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

package xmap

type SimpleMap map[interface{}]interface{}

func NewSimpleMap() *SimpleMap {
	return &SimpleMap{}
}

// Put
//  @Description: Put key,value
//  @receiver m
//  @param key
//  @param value
//
func (m SimpleMap) Put(key, value interface{}) {
	m[key] = value
}

// GetOrPut
//  @Description: Try to add an element to the map. If the element already exists, return the existing element directly without adding it
//  @receiver m
//  @param key
//  @param value
//  @return actual Key already exists return old value, not exists return new value
//  @return loaded loaded success is true, fail is false
//
func (m SimpleMap) GetOrPut(key, value interface{}) (actual interface{}, loaded bool) {
	o, ok := m[key]
	if ok {
		return o, true
	}
	m[key] = value
	return value, false
}

// Get
//  @Description:
//  @receiver m
//  @param key
//  @return value
//  @return loaded loaded success is true, fail is false
//
func (m SimpleMap) Get(key interface{}) (value interface{}, loaded bool) {
	o, ok := m[key]
	if ok {
		return o, true
	}
	return nil, false
}

// Delete
//  @Description:
//  @receiver m
//  @param key
//
func (m SimpleMap) Delete(key interface{}) {
	delete(m, key)
}

// Size
//  @Description:
//  @receiver m
//  @return int
//
func (m SimpleMap) Size() int {
	return len(m)
}

// Foreach
//  @Description: Foreach O(N)
//  @receiver m
//  @param f The function that accepts polling returns true to continue polling and false to terminate polling
//
func (m SimpleMap) Foreach(f func(interface{}, interface{}) bool) {
	for k, v := range m {
		if !f(k, v) {
			break
		}
	}
}

// Find
//  @Description: Find key to Map
//  @receiver m
//  @param key
//  @return bool loaded loaded success is true, fail is false
//
func (m SimpleMap) Find(key interface{}) bool {
	_, ok := m[key]
	return ok
}
