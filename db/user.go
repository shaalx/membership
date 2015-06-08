package db

import (
	"github.com/shaalx/merbership/dbu"
	"github.com/shaalx/merbership/search"
	"labix.org/v2/mgo/bson"
	"log"
)

func searchIUsers(data []byte) []interface{} {
	iusers := search.SearchArray(data, "users", []string{}...)
	return iusers
}

func Persist(DB *dbu.Collection, idata ...interface{}) int {
	return DB.Insert(idata...)
}

func PersistIUsers(DB *dbu.Collection, data []byte) int {
	iuser := searchIUsers(data)
	users := make([]interface{}, 0, len(iuser))
	for _, iu := range iuser {
		bys := dbu.I2JsonBytes(iu)
		if nil != bys {
			uid := search.SearchSValue(bys, "id")
			selector := bson.M{
				"id": uid,
			}
			if 0 == DB.Count(selector) {
				users = append(users, iu)
				log.Printf("%s >>>>>>>>>>>>>>> {DB}", uid)
			} else {
				log.Printf("%s has existed.", uid)
			}
		}
	}
	return Persist(DB, users...)
}
