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
	"path/filepath"
	"reflect"
	"regexp"
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

// Panic asserts that function fn() panics.
// It assumes that recover() either returns a string or
// an error and fails if the message does not match
// the regular expression in 'matches'.
//  @param t tet
//  @param fn function
//  @param matches matcher
func Panic(t *testing.T, fn func(), matches string) {
	if x := doesPanic(2, fn, matches); x != "" {
		fmt.Println(x)
		t.Fail()
	}
}

// doesPanic do panic
//  @param skip judge skip
//  @param fn func
//  @param expr got
//  @return err error
func doesPanic(skip int, fn func(), expr string) (err string) {
	defer func() {
		r := recover()
		if r == nil {
			err = fail(skip, "did not panic")
			return
		}
		var v string
		switch r.(type) {
		case error:
			v = r.(error).Error()
		case string:
			v = r.(string)
		}
		err = matches(skip, v, expr)
	}()
	fn()
	return ""
}

// Matches asserts that a value matches a given regular expression.
//  @param t test
//  @param value value
//  @param expr got
func Matches(t *testing.T, value, expr string) {
	if x := matches(2, value, expr); x != "" {
		fmt.Println(x)
		t.Fail()
	}
}

// matches matches got and expr value
//  @param skip judge skip
//  @param value value
//  @param expr got
//  @return string result
func matches(skip int, value, expr string) string {
	ok, err := regexp.MatchString(expr, value)
	if err != nil {
		return fail(skip, "invalid pattern %q. %s", expr, err)
	}
	if !ok {
		return fail(skip, "got %s which does not match %s", value, expr)
	}
	return ""
}

func fail(skip int, format string, args ...interface{}) string {
	_, file, line, _ := runtime.Caller(skip)
	return fmt.Sprintf("\t%s:%d: %s\n", filepath.Base(file), line, fmt.Sprintf(format, args...))
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
