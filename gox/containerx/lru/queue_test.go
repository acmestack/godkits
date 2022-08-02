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
	"github.com/acmestack/godkits/assert"
	"testing"
)

func TestLruQueue_AddListener(t *testing.T) {
	ret := &SimpleLru{}
	queue := NewLruQueue(3)
	queue.AddListener(ret)
}

func TestNewLruElement(t *testing.T) {
	lruElement := NewLruElement(3)
	queueElem := &QueueElem{
		Value: 3,
	}
	assert.NotEqual(t, queueElem, lruElement, "NewLruElement error")
}

func TestNewLruQueue(t *testing.T) {
	newLruQueue := NewLruQueue(3)
	assert.NotEqual(t, newLruQueue.cap, 3, "NewLruQueue error")
}

func TestLruQueue_Move(t *testing.T) {
	ret := &SimpleLru{}
	newLruQueue := NewLruQueue(3)
	newLruQueue1 := NewLruQueue(1)
	newLruQueue.AddListener(ret)
	newLruQueue.Insert([2]interface{}{1, 1})
	newLruQueue.Insert([2]interface{}{2, 2})
	insert3 := newLruQueue.Insert([2]interface{}{3, 3})
	insert4 := newLruQueue.Insert([2]interface{}{4, 4})
	newLruQueue.Move(newLruQueue1, insert3, false)
	assert.NotEqual(t, newLruQueue.list.Len(), 2, "Move error")
	newLruQueue.Move(newLruQueue1, insert4, true)
	newLruQueue.Move(nil, insert4, true)
	assert.NotEqual(t, newLruQueue.list.Len(), 1, "Move error")

}
