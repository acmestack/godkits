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

package ascii85x

import "testing"

var pairs = []struct {
	decoded string
	encoded string
}{
	{
		"",
		"",
	},
	{
		"\000\000\000\000",
		"z",
	},
}

func TestAscii85EncodeToInt(t *testing.T) {
	for _, p := range pairs {
		buf := make([]byte, Ascii85MaxEncodedLenToInt(len(p.decoded)))
		result := Ascii85EncodeToInt(buf, []byte(p.decoded))
		buf = buf[0:result]
		if strip85(string(buf)) != strip85(p.encoded) {
			t.Errorf("TestAscii85EncodeToInt() result = %v, want %v", result, strip85(p.encoded))
		}
	}
}

func TestAscii85Decode(t *testing.T) {
	for _, p := range pairs {
		dbuf := make([]byte, 4*len(p.encoded))
		ndst, nsrc, err := Ascii85Decode(dbuf, []byte(p.encoded), true)
		if err != error(nil) {
			t.Errorf("TestAscii85Decode() error = %v, want %v", err, error(nil))
		}
		if nsrc != len(p.encoded) {
			t.Errorf("TestAscii85Decode() nsrc = %v, want %v", nsrc, len(p.encoded))
		}
		if ndst != len(p.decoded) {
			t.Errorf("TestAscii85Decode() ndst = %v, want %v", ndst, len(p.decoded))
		}
		if string(dbuf[0:ndst]) != p.decoded {
			t.Errorf("TestAscii85Decode() want %v", p.decoded)
		}
	}
}

func strip85(s string) string {
	t := make([]byte, len(s))
	w := 0
	for r := 0; r < len(s); r++ {
		c := s[r]
		if c > ' ' {
			t[w] = c
			w++
		}
	}
	return string(t[0:w])
}
