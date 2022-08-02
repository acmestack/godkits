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

package system

import (
	"bytes"
	"os"
	"os/exec"
	"runtime"
)

// IsWindows
//  @return bool
func IsWindows() bool {
	return runtime.GOOS == "windows"
}

// IsLinux
//  @Description:
//  @return bool
func IsLinux() bool {
	return runtime.GOOS == "linux"
}

// IsMac
//  @return bool
func IsMac() bool {
	return runtime.GOOS == "darwin"
}

// GetOsEnv
//  @param key
//  @return string
func GetOsEnv(key string) string {
	return os.Getenv(key)
}

// SetOsEnv
//  @param key
//  @param value
//  @return error
func SetOsEnv(key, value string) error {
	return os.Setenv(key, value)
}

// RemoveOsEnv
//  @param key
//  @return error
func RemoveOsEnv(key string) error {
	return os.Unsetenv(key)
}

// CompareOsEnv gets env and Compare
//  @Description:
//  @param key
//  @param comparedEnv
//  @return bool
func CompareOsEnv(key, comparedEnv string) bool {
	env := GetOsEnv(key)
	if env == "" {
		return false
	}
	return env == comparedEnv
}

// ExecCommand execute command shell /bin/bash -c
//  @Description:
//  @param command
//  @return stdout
//  @return stderr
//  @return err
func ExecCommand(command string) (stdout, stderr string, err error) {
	var out bytes.Buffer
	var errout bytes.Buffer

	cmd := exec.Command("/bin/bash", "-c", command)
	if IsWindows() {
		cmd = exec.Command("cmd")
	}
	cmd.Stdout = &out
	cmd.Stderr = &errout
	err = cmd.Run()

	if err != nil {
		stderr = string(errout.Bytes())
	}
	stdout = string(out.Bytes())

	return
}
