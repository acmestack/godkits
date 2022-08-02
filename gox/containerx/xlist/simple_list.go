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

// PushBack add an element at the end of the linked list
//  @receiver l list
//  @param o element
func (l *SimpleList) PushBack(o interface{}) {
	l.l.PushBack(o)
}

// PushFront Add an element at the head of the linked list
//  @receiver l list
//  @param o element
func (l *SimpleList) PushFront(o interface{}) {
	l.l.PushFront(o)
}

// Remove remove element in list
//  @receiver l list
//  @param o element
func (l *SimpleList) Remove(o interface{}) {
	for e := l.l.Front(); e != nil; e = e.Next() {
		if reflect.DeepEqual(o, e.Value) {
			l.l.Remove(e)
			return
		}
	}
}

// Front First element (first element), if nil is not returned
//  @receiver l list
//  @return interface{}
func (l *SimpleList) Front() interface{} {
	e := l.l.Front()
	if e != nil {
		return e.Value
	}
	return nil
}

// Back Tail element (the last element), if nil is not returned
//  @receiver l
//  @return interface{}
func (l *SimpleList) Back() interface{} {
	e := l.l.Back()
	if e != nil {
		return e.Value
	}
	return nil
}

// PopFront Get the first element and remove it, if nil is not returned
//  @receiver l list
//  @return interface{}
func (l *SimpleList) PopFront() interface{} {
	e := l.l.Front()
	if e != nil {
		l.l.Remove(e)
		return e.Value
	}
	return nil
}

// PopBack Get the tail element and remove it, if nil is not returned
//  @receiver l list
//  @return interface{}
func (l *SimpleList) PopBack() interface{} {
	e := l.l.Back()
	if e != nil {
		l.l.Remove(e)
		return e.Value
	}
	return nil
}

// Len get list length
//  @receiver l list
//  @return int list length
func (l *SimpleList) Len() int {
	return l.l.Len()
}

// Foreach
//  @receiver l
//  @param f The function that accepts polling returns true to continue polling and false to terminate polling
func (l *SimpleList) Foreach(f func(interface{}) bool) {
	for e := l.l.Front(); e != nil; e = e.Next() {
		if !f(e.Value) {
			return
		}
	}
}

// Find Query whether there is a parameter object in the linked list
//  @receiver l list
//  @param i elemnt
//  @return bool success is true, fail is false
func (l *SimpleList) Find(i interface{}) bool {
	for e := l.l.Front(); e != nil; e = e.Next() {
		if reflect.DeepEqual(i, e.Value) {
			return true
		}
	}
	return false
}
