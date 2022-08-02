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

package assert

import (
	"reflect"
	"testing"
)

func TestEqual(t *testing.T) {
	Equal(t, "foo", "bar", "msg!")
}

func TestNotEqual(t *testing.T) {
	type R struct {
		i int
		*R
	}
	r := &R{
		i: 1,
		R: &R{
			i: 2,
			R: &R{
				i: 3,
			},
		},
	}
	r.R.R.R = r

	r2 := &R{
		i: 1,
		R: &R{
			i: 2,
			R: &R{
				i: 3,
			},
		},
	}
	r2.R.R.R = r2
	//r2.R.R = r2 // notEqual
	NotEqual(t, r, r2, "msg!")
}

func TestIsNull(t *testing.T) {
	IsNull(t, "foo", "msg!")
}

func TestIsTrue(t *testing.T) {
	IsTrue(t, reflect.DeepEqual("foo", "foo"), "msg!")
}

func TestIsFalse(t *testing.T) {
	IsFalse(t, reflect.DeepEqual("foo", "bar"), "msg!")
}

func TestPanicPanics(t *testing.T) {
	if got, want := doesPanic(2, func() { panic("foo") }, ""), ""; got != want {
		t.Fatalf("got %q want %q", got, want)
	}
}

func TestPanicPanicsAndMatches(t *testing.T) {
	if got, want := doesPanic(2, func() { panic("foo") }, "foo"), ""; got != want {
		t.Fatalf("got %q want %q", got, want)
	}
}

func TestPanicPanicsAndDoesNotMatch(t *testing.T) {
	if got, want := doesPanic(2, func() { panic("foo") }, "bar"), "\tassert.go:118: got foo which does not match bar\n"; got != want {
		t.Fatalf("got %q want %q", got, want)
	}
}

func TestPanicPanicsAndDoesNotPanic(t *testing.T) {
	if got, want := doesPanic(2, func() {}, "bar"), "\tassert.go:121: did not panic\n"; got != want {
		t.Fatalf("got %q want %q", got, want)
	}
}

func TestMatchesMatches(t *testing.T) {
	if got, want := matches(2, "aaa", "a"), ""; got != want {
		t.Fatalf("got %q want %q", got, want)
	}
}

func TestMatchesDoesNotMatch(t *testing.T) {
	if got, want := matches(2, "aaa", "b"), "\tassert_test.go:102: got aaa which does not match b\n"; got != want {
		t.Fatalf("got %q want %q", got, want)
	}
}
