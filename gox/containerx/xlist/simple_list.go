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
	"container/list"
	"reflect"
)

type SimpleList struct {
	l *list.List
}

func NewSimpleList() *SimpleList {
	return &SimpleList{
		l: list.New(),
	}
}

// PushBack
//  @Description: Add an element at the end of the linked list
//  @receiver l
//  @param o
func (l *SimpleList) PushBack(o interface{}) {
	l.l.PushBack(o)
}

// PushFront
//  @Description: Add an element at the head of the linked list
//  @receiver l
//  @param o
func (l *SimpleList) PushFront(o interface{}) {
	l.l.PushFront(o)
}

// Remove
//  @Description: remove O(N)
//  @receiver l
//  @param o
func (l *SimpleList) Remove(o interface{}) {
	for e := l.l.Front(); e != nil; e = e.Next() {
		if reflect.DeepEqual(o, e.Value) {
			l.l.Remove(e)
			return
		}
	}
}

// Front
//  @Description: First element (first element), if nil is not returned
//  @receiver l
//  @return interface{}
func (l *SimpleList) Front() interface{} {
	e := l.l.Front()
	if e != nil {
		return e.Value
	}
	return nil
}

// Back
//  @Description: Tail element (the last element), if nil is not returned
//  @receiver l
//  @return interface{}
func (l *SimpleList) Back() interface{} {
	e := l.l.Back()
	if e != nil {
		return e.Value
	}
	return nil
}

// PopFront
//  @Description: Get the first element and remove it, if nil is not returned
//  @receiver l
//  @return interface{}
func (l *SimpleList) PopFront() interface{} {
	e := l.l.Front()
	if e != nil {
		l.l.Remove(e)
		return e.Value
	}
	return nil
}

// PopBack
//  @Description: Get the tail element and remove it, if nil is not returned
//  @receiver l
//  @return interface{}
func (l *SimpleList) PopBack() interface{} {
	e := l.l.Back()
	if e != nil {
		l.l.Remove(e)
		return e.Value
	}
	return nil
}

// Len
//  @Description:
//  @receiver l
//  @return int
func (l *SimpleList) Len() int {
	return l.l.Len()
}

// Foreach
//  @Description:
//  @receiver l
//  @param f The function that accepts polling returns true to continue polling and false to terminate polling
func (l *SimpleList) Foreach(f func(interface{}) bool) {
	for e := l.l.Front(); e != nil; e = e.Next() {
		if !f(e.Value) {
			return
		}
	}
}

// Find
//  @Description: Query whether there is a parameter object in the linked list
//  @receiver l
//  @param i
//  @return bool bool success is true, fail is false
func (l *SimpleList) Find(i interface{}) bool {
	for e := l.l.Front(); e != nil; e = e.Next() {
		if reflect.DeepEqual(i, e.Value) {
			return true
		}
	}
	return false
}
