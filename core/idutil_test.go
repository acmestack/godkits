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
