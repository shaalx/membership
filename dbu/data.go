package dbu

import (
	"encoding/json"
	"github.com/shaalx/merbership/logu"
	"labix.org/v2/mgo/bson"
)

func Bson2Bytes(m *bson.M) []byte {
	b, err := bson.Marshal(m)
	if logu.CheckErr(err) {
		return []byte{}
	}
	return b
}
func Bytes2Bson(b []byte) *bson.M {
	var ret bson.M
	err := bson.Unmarshal(b, &ret)
	if logu.CheckErr(err) {
		return nil
	}
	return &ret
}

func Bson2JBytes(m *bson.M) []byte {
	b, err := json.Marshal(m)
	if logu.CheckErr(err) {
		return []byte{}
	}
	return b
}
func JBytes2Bson(b []byte) *bson.M {
	var ret bson.M
	err := json.Unmarshal(b, &ret)
	if logu.CheckErr(err) {
		return nil
	}
	return &ret
}

func BsonStruct(structs interface{}, m *bson.M) bool {
	data, err := bson.Marshal(m)
	if logu.CheckErr(err) {
		return false
	}
	err = bson.Unmarshal(data, structs)
	if logu.CheckErr(err) {
		return false
	}
	return true
}

func JsonStruct(structs interface{}, m *bson.M) bool {
	data, err := json.Marshal(m)
	if logu.CheckErr(err) {
		return false
	}
	err = json.Unmarshal(data, structs)
	if logu.CheckErr(err) {
		return false
	}
	return true
}

func I2BsonBytes(i interface{}) []byte {
	b, err := bson.Marshal(i)
	if logu.CheckErr(err) {
		return nil
	}
	return b
}

func I2JsonBytes(i interface{}) []byte {
	b, err := json.Marshal(i)
	if logu.CheckErr(err) {
		return nil
	}
	return b
}
