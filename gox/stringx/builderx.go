/*
 * Copyright (c) 2022, OpeningO
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

package stringx

import (
	"strings"

	"github.com/openingo/godkits/array"
)

// BuilderX extend from strings.Builder
type BuilderX struct {
	strings.Builder
}

// JoinString to current builder
func (builder *BuilderX) JoinString(strArray ...string) (int, error) {
	if array.Empty(strArray) {
		return 0, nil
	}
	total := 0
	for _, str := range strArray {
		write, err := builder.WriteString(str)
		if err != nil {
			return total, err
		}
		total += write
	}
	return total, nil
}

// JoinByte to current builder
func (builder *BuilderX) JoinByte(bytes ...byte) (int, error) {
	if array.Empty(bytes) {
		return 0, nil
	}
	total := 0
	for _, c := range bytes {
		if err := builder.WriteByte(c); err != nil {
			return total, err
		}
		total += 1
	}
	return total, nil
}
