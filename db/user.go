package db

import (
	"fmt"
	"github.com/shaalx/merbership/dbu"
	// "github.com/shaalx/merbership/logu"
	"github.com/shaalx/merbership/search"
	"labix.org/v2/mgo/bson"
	"log"
	"strings"
	"time"
)

type OnlineStatus struct {
	Time          string      `bson:"time"`
	IOnlineStatus interface{} `bson:"online_status"`
}

func Now() string {
	_now := time.Now()
	return _now.String()
	// loc, err := time.LoadLocation("Europe/Paris")
	// if logu.CheckErr(err) {
	// 	return _now.String()
	// }
	// st := "2006-01-02 08:00"
	// time3, _ := time.Parse("2006-01-02 15:04", st)
	// fmt.Println("time3 is :", time3)
	// nowf := time.Now().Format(st)
	// nowp, err := time.ParseInLocation(nowf, st, loc)
	// if logu.CheckErr(err) {
	// 	return _now.String()
	// }
	// fmt.Printf("parse time is %v", nowp)
	// return nowp.String()
}

func searchIUsers(data []byte) []interface{} {
	iusers := search.SearchArray(data, "users", []string{}...)
	return iusers
}

func searchIOnlieStatuses(data []byte) []interface{} {
	ionline_status := search.SearchArray(data, "online_status", []string{}...)
	return ionline_status
}

func statOnline(online_status []interface{}) int {
	online_count := 0
	for _, ion := range online_status {
		if on, ok := ion.(map[string]interface{}); ok {
			status := fmt.Sprintf("%v", on["status"])
			if strings.EqualFold(status, "2") {
				online_count++
			}
		}
	}
	return online_count
}

func Persist(DB *dbu.Collection, idata ...interface{}) int {
	return DB.Insert(idata...)
}

func PersistIUsers(DB *dbu.Collection, data []byte) (int, []string) {
	iuser := searchIUsers(data)
	users := make([]interface{}, 0, len(iuser))
	uids := make([]string, 0, len(iuser))
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
			uids = append(uids, uid)
		}
	}
	return Persist(DB, users...), uids
}

func PersistIOnlineStatuses(DB *dbu.Collection, data []byte) (int, int) {
	ionlines := searchIOnlieStatuses(data)

	online_statuses := make([]interface{}, 0, len(ionlines))
	now := Now()
	for _, ion := range ionlines {
		online_status := OnlineStatus{Time: now, IOnlineStatus: ion}
		online_statuses = append(online_statuses, online_status)
	}
	return Persist(DB, online_statuses...), statOnline(ionlines)
}
