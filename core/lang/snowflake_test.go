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

package lang

import (
	"github.com/acmestack/godkits/assert"
	"sync"
	"testing"
	"time"
)

func TestNewSnowflake(t *testing.T) {
	var i, j int64
	for i = 0; i < 32; i++ {
		for j = 0; j < 32; j++ {
			_, err := NewSnowflake(i, j)
			assert.IsTrue(t, err == nil, err)
		}
	}
	_, err := NewSnowflake(0, -1)
	assert.IsTrue(t, err != nil, err)
	_, err2 := NewSnowflake(-1, 0)
	assert.IsTrue(t, err2 != nil, err)
}

func TestNextVal(t *testing.T) {
	s, err := NewSnowflake(0, 0)
	assert.IsTrue(t, err == nil, err)
	var i int64
	for i = 0; i < sequenceMask*10; i++ {
		val := s.NextVal()
		assert.IsFalse(t, val == 0, err)
	}
}

func TestUnique(t *testing.T) {
	var wg sync.WaitGroup
	var check sync.Map
	s, err := NewSnowflake(0, 0)
	assert.IsTrue(t, err == nil, err)
	for i := 0; i < 1000000; i++ {
		wg.Add(1)
		// Simulate multithreading to generate data
		go func() {
			defer wg.Add(-1)
			val := s.NextVal()
			_, ok := check.Load(val)
			assert.IsTrue(t, !ok, "Data already exists in map")
			check.Store(val, 0)
			assert.IsTrue(t, val != 0, "Unique NextVal Error")
		}()
	}
	wg.Wait()
}

func TestGetTime(t *testing.T) {
	s, err := NewSnowflake(0, 1)
	assert.IsTrue(t, err == nil, err)
	val := s.NextVal()
	formatDate := time.Now().Format("2006-01-02 15:04:05")
	assert.IsTrue(t, formatDate == GetGenTime(val), err)
}

func TestGetDeviceID(t *testing.T) {
	s, err := NewSnowflake(28, 11)
	assert.IsTrue(t, err == nil, err)
	val := s.NextVal()
	dataCenterId, workerId := GetDeviceID(val)
	if dataCenterId != 28 || workerId != 11 {
		t.Fail()
	}
}

func TestGetTimestampStatus(t *testing.T) {
	status := GetTimestampStatus()
	assert.IsTrue(t, status < 100, "epoch exceeded current time")
}
