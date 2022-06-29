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
	"github.com/acmestack/godkits/assert"
	"testing"
)

func TestNewLruCache(t *testing.T) {

	simpleLru := NewLruCache(2)

	simpleLru.Put("1", "1")
	simpleLru.Put("2", "2")
	simpleLru.Put("4", "4")
	// After exceeding capacity, Check whether LRU can delete header elements
	_value1, _ := simpleLru.Get("1")
	assert.NotEqual(t, _value1, nil, "get key 1 value Should be 1")

	// check value is normal
	_value4, _ := simpleLru.Get("4")
	assert.NotEqual(t, _value4, "4", "get key 4 value Should be 4")

}

func TestSimpleLru_Get(t *testing.T) {
	simpleLru := NewLruCache(2)

	simpleLru.Put("1", "1")
	simpleLru.Put("2", "2")
	simpleLru.Put("2", "3")
	// check put repeat key
	_value1, _ := simpleLru.Get("2")
	assert.NotEqual(t, _value1, "3", "get key 2 value Should be 3")

}

func TestSimpleLru_Delete(t *testing.T) {
	simpleLru := NewLruCache(2)

	simpleLru.Put("1", "1")
	simpleLru.Delete("1")
	// check delete key
	_value1, _ := simpleLru.Get("1")
	assert.NotEqual(t, _value1, nil, "get key 1 value Should be nil")
}

func TestSimpleLru_Size(t *testing.T) {
	simpleLru := NewLruCache(2)

	simpleLru.Put("1", "1")
	simpleLru.Put("2", "2")

	// check size
	_size := simpleLru.Size()
	assert.NotEqual(t, _size, 2, "get size Should be 2")
}

func TestSimpleLru_Purge(t *testing.T) {
	simpleLru := NewLruCache(2)
	simpleLru.Put("1", "1")
	simpleLru.Put("2", "2")

	simpleLru.Purge()
	// check Purge
	assert.NotEqual(t, simpleLru.queue, (*LruQueue)(nil), "get simpleLru.queue Should be nil")
}
