package core

import (
	uuid "github.com/satori/go.uuid"
	"gopkg.in/mgo.v2/bson"
	"strconv"
	"strings"
	"time"
)

// UnixMillisId generate timestamp id with millisecond.
//  @return string
func UnixMillisId() string {
	return strconv.FormatInt(time.Now().UnixNano()/1e6, 10)
}

// UnixNanoTimeId generate timestamp id with unixNano.
//  @return string
func UnixNanoTimeId() string {
	return strconv.FormatInt(time.Now().UnixNano(), 10)
}

// UUIDV1 Version 1, based on timestamp and MAC address (RFC 4122).
//  @return string
func UUIDV1() string {
	return uuid.NewV1().String()
}

// UUIDV2 Version 2, based on timestamp, MAC address and POSIX UID/GID (DCE 1.1).
//  @param domain domain
//  @return string uuid string
func UUIDV2(domain byte) string {
	return uuid.NewV2(domain).String()
}

// UUIDV3 Version 3, based on MD5 hashing (RFC 4122).
//  @param ns   uuid namespace
//  @param name domain name
//  @return string uuid string
func UUIDV3(ns uuid.UUID, name string) string {
	return uuid.NewV3(ns, name).String()
}

// UUIDV5 Version 5, based on SHA-1 hashing (RFC 4122).
//  @param ns   uuid namespace
//  @param name domain name
//  @return string uuid string
func UUIDV5(ns uuid.UUID, name string) string {
	return uuid.NewV5(ns, name).String()
}

// RandomUUID Version 4, based on random numbers (RFC 4122).
//  @return string uuid string
func RandomUUID() string {
	return uuid.NewV4().String()
}

// SimpleRandomUUID simple RandomUUI base RandomUUID.
//  @return string return uuid without horizontal line
func SimpleRandomUUID() string {
	return strings.ReplaceAll(RandomUUID(), "-", "")
}

// ObjectId generate mongodb ObjectId, currently, just use mgo.v2 lib, can write self.
//  @return string object id
func ObjectId() string {
	return bson.NewObjectId().Hex()
}

// ObjectIdWithNow generate ObjectId with now, currently, just use mgo.v2 lib, can write self.
//  @return string object id
func ObjectIdWithNow() string {
	return bson.NewObjectIdWithTime(time.Now()).Hex()
}
