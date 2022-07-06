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
	// PushBack
	//  @Description:
	//  @param o
	PushBack(o interface{})

	// PushFront
	//  @Description:
	//  @param o
	PushFront(o interface{})

	// Remove
	//  @Description:
	//  @param o
	Remove(o interface{})

	// Front
	//  @Description:
	//  @return interface{}
	Front() interface{}

	// Back
	//  @Description:
	//  @return interface{}
	Back() interface{}

	// PopFront
	//  @Description:
	//  @return interface{}
	PopFront() interface{}

	// PopBack
	//  @Description:
	//  @return interface{}
	PopBack() interface{}

	// Len
	//  @Description:
	//  @return int
	Len() int

	// Foreach
	//  @Description: The function that accepts polling returns true to continue polling and false to terminate polling
	//  @param f
	Foreach(f func(interface{}) bool)

	// Find
	//  @Description:
	//  @param i
	//  @return bool
	Find(i interface{}) bool
}
