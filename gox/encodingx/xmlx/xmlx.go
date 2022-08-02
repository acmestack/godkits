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

package xmlx

import (
	"encoding/xml"
)

type any = interface{}

// ToXmlBytes convert any to xml bytes
// wrap xml.Marshal
func ToXmlBytes(v any) ([]byte, error) {
	return xml.Marshal(v)
}

// ToXmlString convert any to xml string
// wrap xml.Marshal
func ToXmlString(v any) (string, error) {
	marshal, err := xml.Marshal(v)
	return string(marshal), err
}

// XmlBytesToAny convert xml bytes to any
// wrap xml.Unmarshal
func XmlBytesToAny(bytes []byte, v any) error {
	return xml.Unmarshal(bytes, v)
}

// XmlStringToAny convert xml bytes to any
// wrap xml.Unmarshal
func XmlStringToAny(str string, v any) error {
	return xml.Unmarshal([]byte(str), v)
}
