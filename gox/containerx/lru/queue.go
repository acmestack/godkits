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

import "container/list"

type QueueElem list.Element

// QueueListener
//  @Description: QueueListener
//
type QueueListener interface {

	// PostTouch Element hit, which is called after the element is moved to the queue head
	//  @param v any type
	PostTouch(v interface{})

	// PostInsert Called after the element is newly added to the queue header
	//  @param v any type
	PostInsert(v interface{})

	// PostDelete Called after the element is deleted from the queue
	//  @param v any type
	PostDelete(v interface{})
}

type Queue interface {
	AddListener(listener QueueListener)

	Touch(elem *QueueElem)

	Move(other *LruQueue, elem *QueueElem, notify bool) *QueueElem

	Insert(v interface{}) *QueueElem

	Delete(elem *QueueElem)
}

type LruQueue struct {
	list      *list.List
	cap       int
	listeners []QueueListener
}

// NewLruElement NewLruElement by value
//  @param v value
//  @return *QueueElem
func NewLruElement(v interface{}) *QueueElem {
	return &QueueElem{
		Value: v,
	}
}

// NewLruQueue new a appoint capacity NewLruQueue
//  @param cap capacity
//  @return *LruQueue lruQueue
func NewLruQueue(cap int) *LruQueue {
	return &LruQueue{
		list:      list.New(),
		cap:       cap,
		listeners: []QueueListener{},
	}
}

// AddListener add queue listener
//  @receiver q queue
//  @param listener listener
func (q *LruQueue) AddListener(listener QueueListener) {
	q.listeners = append(q.listeners, listener)
}

// Touch touch an element
//  @receiver q queue
//  @param elem element
func (q *LruQueue) Touch(elem *QueueElem) {
	if elem != nil {
		q.list.MoveToFront((*list.Element)(elem))
		for _, l := range q.listeners {
			l.PostTouch(elem.Value)
		}
	}
}

// Move more element to other queue
//  @receiver q queue
//  @param other other queue
//  @param elem element
//  @param notify
//  @return *QueueElem
func (q *LruQueue) Move(other *LruQueue, elem *QueueElem, notify bool) *QueueElem {
	if other == nil || elem == nil {
		return nil
	}
	v := q.list.Remove((*list.Element)(elem))
	if notify {
		for _, l := range q.listeners {
			l.PostDelete(elem.Value)
		}
	}

	elem = (*QueueElem)(other.list.PushFront(v))
	if notify {
		for _, l := range q.listeners {
			l.PostInsert(elem.Value)
		}
	}
	return elem
}

// Insert insert element into queue
//  @receiver q queue
//  @param v value
//  @return *QueueElem
func (q *LruQueue) Insert(v interface{}) *QueueElem {
	if q.list.Len() == q.cap {
		e := q.list.Back()
		if e != nil {
			q.Delete((*QueueElem)(e))
		}
	}

	e := q.list.PushFront(v)
	for _, l := range q.listeners {
		l.PostInsert(v)
	}
	return (*QueueElem)(e)
}

// Delete delete element in queue
//  @receiver q queue
//  @param elem
func (q *LruQueue) Delete(elem *QueueElem) {
	v := q.list.Remove((*list.Element)(elem))
	for _, l := range q.listeners {
		l.PostDelete(v)
	}
}
