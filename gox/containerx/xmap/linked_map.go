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

import "container/list"

type node struct {
	prev *node
	next *node
	v    interface{}
}

func (n *node) init() {
	n.prev = nil
	n.next = nil
	n.v = nil
}

type LinkedMap struct {
	head node
	m    map[interface{}]*node
}

func NewLinkedMap() *LinkedMap {
	ret := &LinkedMap{
		m: make(map[interface{}]*node),
	}
	ret.head.init()
	return ret
}

func (m *LinkedMap) init() {
	if m.head.next == nil {
		m.head.next = &m.head
		m.head.prev = &m.head
	}
}

// insert
//  @Description: insert
//  @receiver m
//  @param v
//  @param at
//  @return *node
func (m *LinkedMap) insert(v interface{}, at *node) *node {
	e := &node{v: v}
	n := at.next
	at.next = e
	e.prev = at
	e.next = n
	n.prev = e

	return e
}

// Put
//  @Description: Put key,value
//  @receiver m
//  @param key
//  @param value
func (m *LinkedMap) Put(key, value interface{}) {
	if e, ok := m.m[key]; ok {
		e.v = [2]interface{}{key, value}
		return
	}

	m.init()
	n := m.insert([2]interface{}{key, value}, m.head.prev)
	m.m[key] = n
}

// GetOrPut
//  @Description: Try to add an element to the map. If the element already exists, return the existing element directly without adding it
//  @receiver m
//  @param key
//  @param value
//  @return actual Key already exists return old value, not exists return new value
//  @return loaded loaded success is true, fail is false
func (m *LinkedMap) GetOrPut(key, value interface{}) (actual interface{}, loaded bool) {
	o, ok := m.m[key]
	if ok {
		return o.v.([2]interface{})[1], true
	}

	m.init()
	n := m.insert([2]interface{}{key, value}, m.head.prev)
	m.m[key] = n
	return value, false
}

// Get
//  @Description:
//  @receiver m
//  @param key
//  @return value
//  @return loaded loaded success is true, fail is false
func (m *LinkedMap) Get(key interface{}) (value interface{}, loaded bool) {
	o, ok := m.m[key]
	if ok {
		return o.v.([2]interface{})[1], ok
	}
	return nil, false
}

// Delete
//  @Description: delete by key
//  @receiver m
//  @param key
func (m *LinkedMap) Delete(key interface{}) {
	if n, ok := m.m[key]; ok {
		n.prev.next = n.next
		n.next.prev = n.prev
		n.init()
		delete(m.m, key)
	}
}

// Size
//  @Description:
//  @receiver m
//  @return int
func (m *LinkedMap) Size() int {
	return len(m.m)
}

// Foreach
//  @Description: Foreach O(N)
//  @receiver m
//  @param f The function that accepts polling returns true to continue polling and false to terminate polling
func (m *LinkedMap) Foreach(f func(key interface{}, value interface{}) bool) {
	for e := m.head.next; e != nil && e != &m.head; e = e.next {
		kv := e.v.([2]interface{})
		if !f(kv[0], kv[1]) {
			break
		}
	}
}

// Find
//  @Description: Find key to Map
//  @receiver m
//  @param key
//  @return bool loaded loaded success is true, fail is false
func (m *LinkedMap) Find(key interface{}) bool {
	_, ok := m.m[key]
	return ok
}

type SimpleLinkedMap struct {
	l *list.List
	m map[interface{}]*list.Element
}

func NewSimpleLinkedMap() *SimpleLinkedMap {
	return &SimpleLinkedMap{list.New(), make(map[interface{}]*list.Element)}
}

// Put
//  @Description: Put key,value
//  @receiver m
//  @param key
//  @param value
func (m *SimpleLinkedMap) Put(key, value interface{}) {
	if e, ok := m.m[key]; ok {
		e.Value = [2]interface{}{key, value}
	}
	m.m[key] = m.l.PushBack([2]interface{}{key, value})
}

// GetOrPut
//  @Description: Try to add an element to the map. If the element already exists, return the existing element directly without adding it
//  @receiver m
//  @param key
//  @param value
//  @return actual Key already exists return old value, not exists return new value
//  @return loaded loaded success is true, fail is false
func (m *SimpleLinkedMap) GetOrPut(key, value interface{}) (actual interface{}, loaded bool) {
	o, ok := m.m[key]
	if ok {
		return o.Value.([2]interface{})[1], true
	}
	m.m[key] = m.l.PushBack([2]interface{}{key, value})
	return value, false
}

// Get
//  @Description:
//  @receiver m
//  @param key
//  @return value
//  @return loaded loaded success is true, fail is false
func (m *SimpleLinkedMap) Get(key interface{}) (value interface{}, loaded bool) {
	o, ok := m.m[key]
	if ok {
		return o.Value.([2]interface{})[1], ok
	}
	return nil, false
}

// Delete
//  @Description:
//  @receiver m
//  @param key
func (m *SimpleLinkedMap) Delete(key interface{}) {
	if e, ok := m.m[key]; ok {
		m.l.Remove(e)
		delete(m.m, key)
	}
}

// Size
//  @Description:
//  @receiver m
//  @return int
func (m *SimpleLinkedMap) Size() int {
	return len(m.m)
}

// Foreach
//  @Description: Foreach O(N)
//  @receiver m
//  @param f The function that accepts polling returns true to continue polling and false to terminate polling
func (m *SimpleLinkedMap) Foreach(f func(key interface{}, value interface{}) bool) {
	for e := m.l.Front(); e != nil; e = e.Next() {
		kv := e.Value.([2]interface{})
		if !f(kv[0], kv[1]) {
			break
		}
	}
}

// Find
//  @Description: Find key to Map
//  @receiver m
//  @param key
//  @return bool loaded loaded success is true, fail is false
func (m *SimpleLinkedMap) Find(key interface{}) bool {
	_, ok := m.m[key]
	return ok
}
