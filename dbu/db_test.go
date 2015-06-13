package dbu

import (
	"labix.org/v2/mgo/bson"
	// "regexp"
	"testing"
)

var (
	MgoDB_ = NewMgoDB("daocloud")
	// MgoDB_ = NewMgoDB(Conn())
	// upsertC = MgoDB_.GetCollection([]string{"test", "upsert"}...)
	usersC = MgoDB_.GetCollection([]string{"lEyTj8hYrUIKgMfi", "users"}...)
)

func TestConn(t *testing.T) {
	conn := Conn()
	t.Log(conn)
}

type User struct {
	Id   string `bson:"id"`
	Name string `bson:"name"`
}

// func TestUpsert(t *testing.T) {
// 	selector := bson.M{"id": "3"}
// 	change := User{"3", "Two22"}
// 	ret := upsertC.Upsert(selector, change)
// 	t.Log(ret)

// 	change2 := User{"122", "Nil selector "}
// 	ret = upsertC.Upsert(nil, change2)
// 	t.Log(ret)
// }

func TestLike(t *testing.T) {
	// re, err := regexp.Compile(`/S/`)
	// if nil != err {
	// 	t.Error(err)
	// }
	selector := bson.M{"name": bson.RegEx{"a", "s"}}
	t.Log(selector)
	ret := usersC.Like(selector)
	t.Log(ret)
	t.Log(len(ret))

	selector2 := bson.M{"name": bson.RegEx{"蓬", "."}}
	t.Log(selector2)
	ret2 := usersC.Like(selector2)
	t.Log(ret2)
	t.Log(len(ret2))
}
