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

type IMap interface {
	// Put
	//  @Description: put to IMap
	//  @param key
	//  @param value
	//
	Put(key, value interface{})

	// Get
	//  @Description: get by IMap
	//  @param key
	//  @return value
	//  @return loaded success is true, fail is false
	//
	Get(key interface{}) (value interface{}, loaded bool)

	// Delete
	//  @Description: delete by IMap
	//  @param key
	//
	Delete(key interface{})

	// Size
	//  @Description: get IMap Size
	//  @return int
	//
	Size() int
}

type Map interface {
	IMap

	// GetOrPut
	//  @Description: get or put to Map
	//  @param key
	//  @param value
	//  @return actual Key already exists return old value, not exists return new value
	//  @return loaded loaded success is true, fail is false
	//
	GetOrPut(key, value interface{}) (actual interface{}, loaded bool)

	// Foreach
	//  @Description: Foreach Map O(N)
	//  @param f The function that accepts polling returns true to continue polling and false to terminate polling
	//
	Foreach(f func(interface{}, interface{}) bool)
}
