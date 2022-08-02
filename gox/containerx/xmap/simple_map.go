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

package xmap

type SimpleMap map[interface{}]interface{}

func NewSimpleMap() *SimpleMap {
	return &SimpleMap{}
}

// Put put key,value into map
//  @receiver m map
//  @param key   key
//  @param value value
func (m SimpleMap) Put(key, value interface{}) {
	m[key] = value
}

// GetOrPut Try to add an element to the map. If the element already exists, return the existing element directly without adding it
//  @receiver m map
//  @param key   key
//  @param value value
//  @return actual Key already exists return old value, not exists return new value
//  @return loaded success is true, fail is false
func (m SimpleMap) GetOrPut(key, value interface{}) (actual interface{}, loaded bool) {
	o, ok := m[key]
	if ok {
		return o, true
	}
	m[key] = value
	return value, false
}

// Get get value by key, if exist return true, not exist return false
//  @receiver m map
//  @param key    key
//  @return value value
//  @return loaded success is true, fail is false
func (m SimpleMap) Get(key interface{}) (value interface{}, loaded bool) {
	o, ok := m[key]
	if ok {
		return o, true
	}
	return nil, false
}

// Delete delete element by key
//  @receiver m map
//  @param key key
func (m SimpleMap) Delete(key interface{}) {
	delete(m, key)
}

// Size calc map size
//  @receiver m map
//  @return int map size
func (m SimpleMap) Size() int {
	return len(m)
}

// Foreach O(N)
//  @receiver m map
//  @param f The function that accepts polling returns true to continue polling and false to terminate polling
func (m SimpleMap) Foreach(f func(interface{}, interface{}) bool) {
	for k, v := range m {
		if !f(k, v) {
			break
		}
	}
}

// Find find key in Map
//  @receiver m map
//  @param key key
//  @return bool loaded success is true, fail is false
func (m SimpleMap) Find(key interface{}) bool {
	_, ok := m[key]
	return ok
}
