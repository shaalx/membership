package dbu

// import (
// 	// "labix.org/v2/mgo/bson"
// 	"testing"
// )

// // var (
// // 	testBson = &bson.M{
// // 		"AK47": "IS A GUN",
// // 	}
// // )

// // func TestBson2Bytes(t *testing.T) {
// // 	b := Bson2Bytes(testBson)
// // 	t.Log(b)
// // 	t.Log(string(b))

// // 	retBson := Bytes2Bson(b)
// // 	t.Log(retBson)
// // }

// // func TestBson2JBytes(t *testing.T) {
// // 	b := Bson2JBytes(testBson)
// // 	t.Log(b)
// // 	t.Log(string(b))

// // 	retBson := JBytes2Bson(b)
// // 	t.Log(retBson)
// // }

// // // invalid character '\x18' looking for beginning of value
// // func TestBsonBytes2JBytes(t *testing.T) {
// // 	b := Bson2Bytes(testBson)
// // 	t.Log(b)
// // 	t.Log(string(b))

// // 	retBson := JBytes2Bson(b)
// // 	t.Log(retBson)
// // }

// // // Document is corrupted
// // func TestJBytes2BsonBytes(t *testing.T) {
// // 	b := Bson2JBytes(testBson)
// // 	t.Log(b)
// // 	t.Log(string(b))

// // 	retBson := Bytes2Bson(b)
// // 	t.Log(retBson)
// // }

// type Person struct {
// 	Name string `bson:"name"`
// 	Age  int    `bson:"age"`
// }

// /*

// var (
// 	person = Person{"Coco", 21}
// )

// func TestStructIn(t *testing.T) {
// 	db := NewMgoDB("")
// 	defer db.Close()
// 	ok := db.GetCollection([]string{"nation", "person"}...).Insert(person)
// 	t.Log(ok)
// }*/

// func TestBsonStructOut(t *testing.T) {
// 	db := NewMgoDB("")
// 	defer db.Close()
// 	personb := db.GetCollection([]string{"nation", "person"}...).Select(nil)

// 	var ret Person
// 	ok := BsonStruct(&ret, personb)
// 	t.Log(ok)
// 	t.Log(ret)
// }

// func TestJsonStructOut(t *testing.T) {
// 	db := NewMgoDB("")
// 	defer db.Close()
// 	personb := db.GetCollection([]string{"nation", "person"}...).Select(nil)

// 	var ret Person
// 	ok := JsonStruct(&ret, personb)
// 	t.Log(ok)
// 	t.Log(ret)
// }
