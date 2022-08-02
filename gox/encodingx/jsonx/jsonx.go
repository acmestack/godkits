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

package jsonx

import "encoding/json"

type any = interface{}

// ToJsonBytes convert any to json bytes
// wrap json.Marshal
func ToJsonBytes(v any) ([]byte, error) {
	return json.Marshal(v)
}

// ToJsonString convert any to json string
// wrap json.Marshal
func ToJsonString(v any) (string, error) {
	marshal, err := json.Marshal(v)
	return string(marshal), err
}

// JsonBytesToAny convert json bytes to any
// wrap json.Unmarshal
func JsonBytesToAny(bytes []byte, v any) error {
	return json.Unmarshal(bytes, v)
}

// JsonStringToAny convert json bytes to any
// wrap json.Unmarshal
func JsonStringToAny(str string, v any) error {
	return json.Unmarshal([]byte(str), v)
}
