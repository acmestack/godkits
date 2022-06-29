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

package lru

import (
	"github.com/acmestack/godkits/gox/containerx/xmap"
)

// LRU
//  @Description: interface LRU
//
type LRU interface {
	xmap.IMap

	Purge()
}

// SimpleLru
//  @Description: type SimpleLru
//
type SimpleLru struct {
	m     map[interface{}]*QueueElem
	queue *LruQueue

	cap int
}

// NewLruCache
//  @Description: new a appoint capacity SimpleLru
//  @param capacity
//  @return *SimpleLru
//
func NewLruCache(capacity int) *SimpleLru {
	ret := &SimpleLru{
		m:   map[interface{}]*QueueElem{},
		cap: capacity,
	}
	ret.queue = NewLruQueue(capacity)
	ret.queue.AddListener(ret)
	return ret
}

func (m *SimpleLru) PostTouch(v interface{}) {
}

func (m *SimpleLru) PostInsert(v interface{}) {
}

// PostDelete
//  import PostDelete
//  @Description: Purge data
//  @receiver m
//
func (m *SimpleLru) PostDelete(v interface{}) {
	key := v.([2]interface{})[0]
	delete(m.m, key)
}

func (m *SimpleLru) hit(key interface{}, hit bool) {

}

// Purge
//  @Description: Purge data
//  @receiver m
//
func (m *SimpleLru) Purge() {
	m.queue.listeners = nil
	m.queue.list = nil
	m.queue = nil
}

// Put
//  @Description: put value to lru
//  @receiver m
//  @param key key
//  @param value value
//
func (m *SimpleLru) Put(key, value interface{}) {
	e, ok := m.m[key]
	if ok {
		e.Value = [2]interface{}{key, value}
		m.queue.Touch(e)
	} else {
		elem := m.queue.Insert([2]interface{}{key, value})
		m.m[key] = elem
	}
}

// Get
//  @Description: Get key to value
//  @receiver m SimpleLru
//  @param key key
//  @return value value
//  @return loaded success is true, fail is false
//
func (m *SimpleLru) Get(key interface{}) (value interface{}, loaded bool) {
	v, ok := m.m[key]
	if ok {
		m.queue.Touch(v)
		// 命中
		m.hit(key, true)
		return v.Value.([2]interface{})[1], true
	} else {
		// 未命中
		m.hit(key, false)
		return nil, false
	}
}

// Delete
//  @Description: delete by key
//  @receiver m
//  @param key
//
func (m *SimpleLru) Delete(key interface{}) {
	v, ok := m.m[key]
	if ok {
		m.queue.Delete(v)
	}
}

// Size
//  @Description: get Lru queue Size
//  @receiver m
//  @return int
//
func (m *SimpleLru) Size() int {
	return len(m.m)
}
