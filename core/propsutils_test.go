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

package core

import (
	"github.com/acmestack/godkits/assert"
	"testing"
)

func TestGetConfigFromFile(t *testing.T) {
	p := NewProperties()
	_ = p.GetConfigFromFile("props.properties")
	val, ok := p.Get("id")
	assert.IsTrue(t, ok)
	assert.Equal(t, 1, val)
	val2, ok2 := p.Get("name")
	assert.IsTrue(t, ok2)
	assert.Equal(t, "xiaoming", val2)
	val3, ok3 := p.Get("mm")
	assert.IsFalse(t, ok3)
	assert.IsTrue(t, val3 == "")
}
