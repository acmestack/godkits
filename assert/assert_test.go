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
	"reflect"
	"testing"
)

func TestEqual(t *testing.T) {
	Equal(t, "foo", "bar", "msg!")
	//Equal(t, "foo", "bar", "this should blow up")
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
	//r2.R.R = r2
	NotEqual(t, r, r2, "msg!")
	//NotEqual(t, "foo", "foo", "this should blow up")
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
