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

package sha1x

import "testing"

func TestSha1(t *testing.T) {
	testItems := []struct {
		in  string
		out string
	}{
		{
			"aaa",
			"7e240de74fb1ed08fa08d38063f6a6a91462a815",
		},
		{
			"hello",
			"aaf4c61ddcc5e8a2dabede0f3b482cd9aea9434d",
		},
		{
			"",
			"da39a3ee5e6b4b0d3255bfef95601890afd80709",
		},
		{
			"                                            ",
			"e0ebcc1adc06efa5740dc2a67a9e7dee824a7810",
		},
	}

	for _, item := range testItems {
		if result := Sha1(item.in); result != item.out {
			t.Errorf("Sha1() Error, want=%v, got=%v", item.out, result)
		}
	}
}
