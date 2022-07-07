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

// IsNull Judge whether the value is nil
//  @param t    test
//  @param got  want type
//  @param args error prompt content
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

// IsTrue Judge whether the value is true
//  @param t      test
//  @param result true
//  @param args   error prompt content
func IsTrue(t *testing.T, result bool, args ...interface{}) {
	tt(t, result, 1, args...)
}

// IsFalse Judge whether the value is false
//  @param t      test
//  @param result false
//  @param args   error prompt content
func IsFalse(t *testing.T, result bool, args ...interface{}) {
	tt(t, !result, 1, args...)
}

// Equal Judge whether the values are not equal
//  @param t   test
//  @param exp expect result
//  @param got compared value
//  @param args error prompt content
func Equal(t *testing.T, exp, got interface{}, args ...interface{}) {
	equal(t, exp, got, 1, args...)
}

// NotEqual Judge whether the values are equal
//  @param t   test
//  @param exp expect result
//  @param got compared value
//  @param args error prompt content
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

// assert  assertion
//  @param t      test
//  @param result result
//  @param f      function
//  @param cd	  hierarchy
func assert(t *testing.T, result bool, f func(), cd int) {
	if !result {
		_, file, line, _ := runtime.Caller(cd + 1)
		t.Errorf("%s:%d", file, line)
		f()
		t.FailNow()
	}
}

// equal judge equal
//  @param t   test
//  @param exp expect result
//  @param got compared value
//  @param cd  hierarchy
//  @param args error prompt content
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

// tt test equal
//  @param t      test
//  @param result wanted result
//  @param cd     hierarchy
//  @param args   error prompt content
func tt(t *testing.T, result bool, cd int, args ...interface{}) {
	fn := func() {
		t.Errorf("!  Failure")
		if len(args) > 0 {
			t.Error("!", " -", fmt.Sprint(args...))
		}
	}
	assert(t, result, fn, cd+1)
}
