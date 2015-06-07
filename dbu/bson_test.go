package dbu

import (
	"labix.org/v2/mgo/bson"
	"testing"
)

var (
	testBson = &bson.M{
		"AK47": "IS A GUN",
	}
)

func TestBson2Bytes(t *testing.T) {
	b := Bson2Bytes(testBson)
	t.Log(b)
	t.Log(string(b))

	retBson := Bytes2Bson(b)
	t.Log(retBson)
}

func TestBson2JBytes(t *testing.T) {
	b := Bson2JBytes(testBson)
	t.Log(b)
	t.Log(string(b))

	retBson := JBytes2Bson(b)
	t.Log(retBson)
}

// invalid character '\x18' looking for beginning of value
func TestBsonBytes2JBytes(t *testing.T) {
	b := Bson2Bytes(testBson)
	t.Log(b)
	t.Log(string(b))

	retBson := JBytes2Bson(b)
	t.Log(retBson)
}

// Document is corrupted
func TestJBytes2BsonBytes(t *testing.T) {
	b := Bson2JBytes(testBson)
	t.Log(b)
	t.Log(string(b))

	retBson := Bytes2Bson(b)
	t.Log(retBson)
}
