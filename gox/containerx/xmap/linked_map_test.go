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

import (
	"github.com/acmestack/godkits/assert"
	"testing"
)

func TestNewLinkedMap(t *testing.T) {
	linkedMap := NewLinkedMap()

	linkedMap.Put("key", "value1")
	value1, _ := linkedMap.Get("key")
	assert.NotEqual(t, value1, "value1", "value1 error")
	getOrPut, _ := linkedMap.GetOrPut("key2", "value2")
	linkedMap.GetOrPut("key3", "value3")
	assert.NotEqual(t, getOrPut, "value2", "value1 error")
	var index = 0
	linkedMap.Foreach(func(key interface{}, value interface{}) bool {
		index++
		return true
	})
	linkedMap.Foreach(func(key interface{}, value interface{}) bool {
		return false
	})
	assert.NotEqual(t, index, 3, "Foreach error")

	assert.NotEqual(t, linkedMap.Size(), 3, "Size error")

	assert.NotEqual(t, linkedMap.Find("key"), true, "Find error")

	linkedMap.Delete("key")
	assert.NotEqual(t, linkedMap.Size(), 2, "Delete error")

	get2, _ := linkedMap.Get("key4")
	assert.NotEqual(t, get2, nil, "Get error")

	_, getOrPut = linkedMap.GetOrPut("key3", "key4")
	assert.NotEqual(t, getOrPut, true, "Get error")

}
func TestNewSimpleLinkedMap(t *testing.T) {
	linkedMap := NewSimpleLinkedMap()

	linkedMap.Put("key", "value1")
	value1, _ := linkedMap.Get("key")
	assert.NotEqual(t, value1, "value1", "value1 error")
	getOrPut, _ := linkedMap.GetOrPut("key2", "value2")
	linkedMap.GetOrPut("key3", "value3")
	assert.NotEqual(t, getOrPut, "value2", "value1 error")
	var index = 0
	linkedMap.Foreach(func(key interface{}, value interface{}) bool {
		index++
		return true
	})
	linkedMap.Foreach(func(key interface{}, value interface{}) bool {
		return false
	})
	assert.NotEqual(t, index, 3, "Foreach error")

	assert.NotEqual(t, linkedMap.Size(), 3, "Size error")

	assert.NotEqual(t, linkedMap.Find("key"), true, "Find error")

	linkedMap.Delete("key")
	assert.NotEqual(t, linkedMap.Size(), 2, "Delete error")

	get2, _ := linkedMap.Get("key4")
	assert.NotEqual(t, get2, nil, "Get error")

	_, getOrPut = linkedMap.GetOrPut("key3", "key4")
	assert.NotEqual(t, getOrPut, true, "Get error")

}
