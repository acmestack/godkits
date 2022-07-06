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

package xlist

import (
	"github.com/acmestack/godkits/assert"
	"testing"
)

func TestNewSimpleList(t *testing.T) {
	simpleList := NewSimpleList()
	type R struct {
		i int
	}
	r := &R{
		i: 1,
	}
	r2 := &R{
		i: 2,
	}
	simpleList.PushBack(r)
	simpleList.PushBack(r2)
	simpleList.PushFront(r2)
	simpleList.PushFront(r)

	var index = 0

	simpleList.Foreach(func(key interface{}) bool {
		return false
	})
	simpleList.Foreach(func(key interface{}) bool {
		index++
		return true
	})
	assert.NotEqual(t, index, simpleList.Len(), "Foreach error")

	assert.NotEqual(t, simpleList.Find(r2), true, "r2 does not exist in the simpleList")

	assert.NotEqual(t, simpleList.PopBack(), r2, "PopBack error")
	assert.NotEqual(t, simpleList.PopFront(), r, "PopFront error")

	simpleList.Remove(r)
	assert.NotEqual(t, simpleList.Find(r), false, "Remove error")

	assert.NotEqual(t, simpleList.Back(), r2, "Back error")
	assert.NotEqual(t, simpleList.Front(), r2, "Front error")

	assert.NotEqual(t, simpleList.PopBack(), r2, "PopBack error")
	assert.NotEqual(t, simpleList.PopBack(), nil, "PopBack error")
	assert.NotEqual(t, simpleList.PopFront(), nil, "PopFront error")

	assert.NotEqual(t, simpleList.Len(), 0, "Len error")
	assert.NotEqual(t, simpleList.Back(), nil, "Back error")
	assert.NotEqual(t, simpleList.Front(), nil, "Front error")

}
