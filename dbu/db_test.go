package dbu

import (
	"labix.org/v2/mgo/bson"
	"testing"
)

var (
	MgoDB_  = NewMgoDB(Conn())
	upsertC = MgoDB_.GetCollection([]string{"test", "upsert"}...)
)

func TestConn(t *testing.T) {
	conn := Conn()
	t.Log(conn)
}

type User struct {
	Id   string `bson:"id"`
	Name string `bson:"name"`
}

func TestUpsert(t *testing.T) {
	selector := bson.M{"id": "3"}
	change := User{"3", "Two22"}
	ret := upsertC.Upsert(selector, change)
	t.Log(ret)

	change2 := User{"122", "Nil selector "}
	ret = upsertC.Upsert(nil, change2)
	t.Log(ret)
}
