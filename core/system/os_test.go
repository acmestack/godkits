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

package system

import (
	"github.com/acmestack/godkits/assert"
	"strings"
	"testing"
)

func TestOsDetection(t *testing.T) {

	osType, _, _ := ExecCommand("echo $OSTYPE")
	if strings.Index(osType, "linux") != -1 {
		assert.NotEqual(t, true, IsLinux(), "")
	}
	if strings.Index(osType, "darwin") != -1 {
		assert.NotEqual(t, true, IsMac(), "")
	}
}

func TestOsEnvOperation(t *testing.T) {

	envNotExist := GetOsEnv("foo")
	assert.NotEqual(t, "", envNotExist, "")

	err := SetOsEnv("foo", "foo_value")
	if err != nil {
		t.Fail()
	}
	envExist := GetOsEnv("foo")
	assert.NotEqual(t, "foo_value", envExist, "")

	assert.NotEqual(t, true, CompareOsEnv("foo", "foo_value"))
	assert.NotEqual(t, false, CompareOsEnv("foo", "abc"))
	assert.NotEqual(t, false, CompareOsEnv("abc", "abc"))
	assert.NotEqual(t, false, CompareOsEnv("abc", "abc"))

	err2 := RemoveOsEnv("foo")
	if err2 != nil {
		t.Fail()
	}
	assert.NotEqual(t, false, CompareOsEnv("foo", "foo_value"))
}

func TestExecCommand(t *testing.T) {

	_, _, err := ExecCommand("ls")
	assert.NotEqual(t, err, nil)
	if err != nil {
		t.Logf("error: %v\n", err)
	}

	_, _, err = ExecCommand("abc")
	if err != nil {
		t.Logf("error: %v\n", err)
	}

	if !IsWindows() {
		assert.Equal(t, err, nil)
	}
}
