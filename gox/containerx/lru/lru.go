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

package lru

import (
	"github.com/acmestack/godkits/gox/containerx/xmap"
)

type LRU interface {
	xmap.IMap

	Purge()
}

// SimpleLru type SimpleLru
type SimpleLru struct {
	m     map[interface{}]*QueueElem
	queue *LruQueue

	cap int
}

// NewLruCache new a appoint capacity SimpleLru
//  @param capacity lru cache capacity
//  @return *SimpleLru simple lru
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

// PostDelete delete data
//  @receiver m lru cache
//  @param v param
func (m *SimpleLru) PostDelete(v interface{}) {
	key := v.([2]interface{})[0]
	delete(m.m, key)
}

func (m *SimpleLru) hit(key interface{}, hit bool) {

}

// Purge Purge data
//  @receiver m
func (m *SimpleLru) Purge() {
	m.queue.listeners = nil
	m.queue.list = nil
	m.queue = nil
}

// Put  put value to lru
//  @receiver m lru cache
//  @param key key
//  @param value value
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

// Get get key to value
//  @receiver m SimpleLru
//  @param key key
//  @return value value
//  @return loaded success is true, fail is false
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

// Delete delete by key
//  @receiver m lru cache
//  @param key key
func (m *SimpleLru) Delete(key interface{}) {
	v, ok := m.m[key]
	if ok {
		m.queue.Delete(v)
	}
}

// Size get Lru queue Size
//  @receiver m lru cache
//  @return int lru cache size
func (m *SimpleLru) Size() int {
	return len(m.m)
}
