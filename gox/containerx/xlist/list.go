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

type List interface {

	// PushBack push value in tail
	//  @param o value
	PushBack(o interface{})

	// PushFront push value in head
	//  @param o value
	PushFront(o interface{})

	// Remove remove an element
	//  @param o value
	Remove(o interface{})

	// Front get head element
	//  @return interface{}
	Front() interface{}

	// Back get tail element
	//  @return interface{}
	Back() interface{}

	// PopFront pop head element
	//  @return interface{}
	PopFront() interface{}

	// PopBack pop tail element
	//  @return interface{}
	PopBack() interface{}

	// Len calc list length
	//  @return int
	Len() int

	// Foreach The function that accepts polling returns true to continue polling and false to terminate polling
	//  @param f function
	Foreach(f func(interface{}) bool)

	// Find find element in list
	//  @param i element
	//  @return bool exist => true, not exist => false
	Find(i interface{}) bool
}
