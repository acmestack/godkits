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

// Map key value mapping
type Map struct {
	size int
	kv   map[any]any
}

// NewMap new map
func NewMap(cap int) *Map {
	m := Map{}
	m.kv = make(map[any]any, cap)
	m.size = 0
	return &m
}

// ContainsKey check the key exist or not
func (m *Map) ContainsKey(key any) bool {
	checkMap(m)
	return m.kv[key] != nil
}

// Put key value to map
func (m *Map) Put(key any, value any) bool {
	checkMap(m)
	if !m.ContainsKey(key) {
		m.size++
	}
	m.kv[key] = value
	return true
}

// Get value by key, may be nil
func (m *Map) Get(key any) any {
	checkMap(m)
	return m.kv[key]
}

// Value by key
func (m *Map) Value(key any) any {
	checkMap(m)
	return m.kv[key]
}

// Remove by key
func (m *Map) Remove(key any) bool {
	checkMap(m)
	delete(m.kv, key)
	m.size--
	return true
}

// Empty the map empty or not
func (m *Map) Empty() bool {
	checkMap(m)
	return m.size == 0
}

// Size map size
func (m *Map) Size() int {
	checkMap(m)
	return m.size
}
