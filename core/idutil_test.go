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

package core

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
	"testing"
)

func TestGenUUIDRandom(t *testing.T) {
	fmt.Println(UnixMillisId())
	fmt.Println(UnixMillisId())
	fmt.Println(UnixNanoTimeId())
	fmt.Println(UUIDV1())
	fmt.Println(UUIDV2('A'))
	fmt.Println(UUIDV3(uuid.NamespaceOID, "test"))
	fmt.Println(UUIDV5(uuid.NamespaceOID, "test"))
	fmt.Println(RandomUUID())
	fmt.Println(SimpleRandomUUID())
	fmt.Println(ObjectId())
	fmt.Println(ObjectIdWithNow())
}
