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
	"github.com/acmestack/godkits/log"
	uuid "github.com/satori/go.uuid"
	"testing"
)

func TestGenUUIDRandom(t *testing.T) {
	log.Info(UnixMillisId())
	log.Info(UnixNanoTimeId())
	log.Info(UUIDV1())
	log.Info(UUIDV2('A'))
	log.Info(UUIDV3(uuid.NamespaceOID, "test"))
	log.Info(UUIDV5(uuid.NamespaceOID, "test"))
	log.Info(RandomUUID())
	log.Info(SimpleRandomUUID())
	log.Info(ObjectId())
	log.Info(ObjectIdWithNow())
}
