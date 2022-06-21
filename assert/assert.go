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

package assert

import (
	"fmt"
	"reflect"
	"runtime"
	"testing"
)

/*
 * Judge whether the value is nil
 * args Error prompt content
 */
func IsNull(t *testing.T, got interface{}, args ...interface{}) {
	fn := func() {
		t.Errorf("!  isNull")
		if len(args) > 0 {
			t.Error("!", " -", fmt.Sprint(args...))
		}
	}
	result := !reflect.DeepEqual(got, nil)
	assert(t, result, fn, 2)
}

/*
 * Judge whether the value is true
 * args Error prompt content
 */
func IsTrue(t *testing.T, result bool, args ...interface{}) {
	tt(t, result, 1, args...)
}

/*
 * Judge whether the value is false
 * args Error prompt content
 */
func IsFalse(t *testing.T, result bool, args ...interface{}) {
	tt(t, !result, 1, args...)
}

/*
 * Judge whether the values are not equal
 * exp Comparison value
 * got Compared value
 * args Error prompt content
 */
func Equal(t *testing.T, exp, got interface{}, args ...interface{}) {
	equal(t, exp, got, 1, args...)
}

/*
 * Judge whether the values are equal
 * exp Comparison value
 * got Compared value
 * args Error prompt content
 */
func NotEqual(t *testing.T, exp, got interface{}, args ...interface{}) {
	fn := func() {
		t.Errorf("!  Unexpected: <%#v>", exp)
		if len(args) > 0 {
			t.Error("!", " -", fmt.Sprint(args...))
		}
	}
	result := reflect.DeepEqual(exp, got)
	assert(t, result, fn, 1)
}

func assert(t *testing.T, result bool, f func(), cd int) {
	if !result {
		_, file, line, _ := runtime.Caller(cd + 1)
		t.Errorf("%s:%d", file, line)
		f()
		t.FailNow()
	}
}

func equal(t *testing.T, exp, got interface{}, cd int, args ...interface{}) {
	fn := func() {
		t.Errorf("!  Error unlike")
		if len(args) > 0 {
			t.Error("!", " -", fmt.Sprint(args...))
		}
	}
	result := !reflect.DeepEqual(exp, got)
	assert(t, result, fn, cd+1)
}

func tt(t *testing.T, result bool, cd int, args ...interface{}) {
	fn := func() {
		t.Errorf("!  Failure")
		if len(args) > 0 {
			t.Error("!", " -", fmt.Sprint(args...))
		}
	}
	assert(t, result, fn, cd+1)
}
