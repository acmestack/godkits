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

package xmap

type IMap interface {
	// Put put (key, value) into map
	//  @param key   key
	//  @param value value
	Put(key, value interface{})

	// Get get value element by key
	//  @param key key
	//  @return value value
	//  @return loaded success is true, fail is false
	Get(key interface{}) (value interface{}, loaded bool)

	// Delete delete element by key in map
	//  @param key key
	Delete(key interface{})

	// Size calc size
	//  @return int map size
	Size() int
}

type Map interface {
	IMap

	// GetOrPut get value when key exist, not exists return new value, value into map, allow nil
	//  @param key   key
	//  @param value value
	//  @return actual Key already exists return old value, not exists return new value
	//  @return loaded uccess is true, fail is false
	GetOrPut(key, value interface{}) (actual interface{}, loaded bool)

	// Foreach O(N)
	//  @param f The function that accepts polling returns true to continue polling and false to terminate polling
	Foreach(f func(interface{}, interface{}) bool)
}
