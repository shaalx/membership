package db

import (
	"fmt"
	"github.com/shaalx/membership/dbu"
	"github.com/shaalx/membership/logu"
	"github.com/shaalx/membership/u"
	// "github.com/shaalx/membership/logu"
	"github.com/shaalx/membership/search"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"log"
	"strings"
	"time"
)

type ViCount struct {
	UID    string `bson:"uid"`
	Status string `bson:"status"`
	VCount int64  `bson:"vcount"`
}

type VViCount struct {
	UID      string `bson:"uid"`
	VCount   int64  `bson:"vcount"`
	Status   string `bson:"status"`
	Avatar   string `bson:"avatar"`
	Name     string `bson:"name"`
	NickName string `bson:"nickname"`
}

func NewVViCount(uid string, vcount int64, status, avatar, name, nickname string) *VViCount {
	return &VViCount{
		UID:      uid,
		VCount:   vcount,
		Status:   status,
		Avatar:   avatar,
		Name:     name,
		NickName: nickname,
	}
}

type OnlineStatus struct {
	Time          string      `bson:"time"`
	IOnlineStatus interface{} `bson:"online_status"`
}
type DOnlineStatus struct {
	Time          int64       `bson:"time"`
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

func NowUnix() int64 {
	_now := time.Now()
	return _now.Unix()
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

func SearchIUsers(data []byte) []interface{} {
	iusers := search.SearchArray(data, "users", []string{}...)
	return iusers
}

func SearchIOnlieStatuses(data []byte) []interface{} {
	ionline_status := search.SearchArray(data, "online_status", []string{}...)
	return ionline_status
}

func OnlineCount(online_status []interface{}) int {
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

func SearchOnlineUids(online_status []interface{}) ([]interface{}, []interface{}, []interface{}) {
	istatus0 := make([]interface{}, 0, len(online_status))
	istatus1 := make([]interface{}, 0, len(online_status))
	istatus2 := make([]interface{}, 0, len(online_status))
	for _, ion := range online_status {
		if on, ok := ion.(map[string]interface{}); ok {
			status := fmt.Sprintf("%v", on["status"])
			iuid := on["uid"]
			if strings.EqualFold(status, "2") {
				istatus2 = append(istatus2, iuid)
			} else if strings.EqualFold(status, "1") {
				istatus1 = append(istatus1, iuid)
			} else if strings.EqualFold(status, "0") {
				istatus0 = append(istatus0, iuid)
			}
		}
	}
	return istatus0, istatus1, istatus2
}

func DistinctUids(onlineC *dbu.Collection) []string {
	var ret []interface{}
	err := onlineC.C.Find(nil).Distinct("online_status.uid", &ret)
	if logu.CheckErr(err) {
		return nil
	}
	uids := make([]string, 0, len(ret))
	for _, iui := range ret {
		ui := fmt.Sprintf("%v", iui)
		ui = strings.TrimSpace(ui)
		uids = append(uids, ui)
	}
	return uids
}

func OnlineUids(iuids ...string) ([]interface{}, []interface{}, []interface{}) {
	_url2 := "https://api.simplr.cn/0.1/user/online_status.json?uids="
	juids := strings.Join(iuids, ",")
	url_ := _url2 + juids
	bys := u.Fetch(url_)
	ionline_users := SearchIOnlieStatuses(bys)
	return SearchOnlineUids(ionline_users)
}

func Persist(DB *dbu.Collection, idata ...interface{}) int {
	return DB.Insert(idata...)
}

func UpdatePersist(DB *dbu.Collection, idata ...interface{}) int {
	var selector bson.M
	var count int
	for _, it := range idata {
		data, ok := it.(map[string]interface{})
		if ok {
			selector = bson.M{"id": data["id"]}
			DB.C.Update(selector, it)
			count++
		}
	}
	return count
}

func PersistIUsers(DB *dbu.Collection, data []byte, updateInsert bool) (int, []string) {
	iuser := SearchIUsers(data)
	users := make([]interface{}, 0, len(iuser))
	updateInsertUsers := make([]interface{}, 0, len(iuser))
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
			} else if updateInsert {
				updateInsertUsers = append(updateInsertUsers, iu)
				log.Printf("%s ->->->->->->->->->-> {DB}", uid)
			} else {
				log.Printf("%s has existed.", uid)
			}
			uids = append(uids, uid)
		}
	}
	var updateCount int
	if updateInsert {
		updateCount = UpdatePersist(DB, updateInsertUsers...)
	}
	return Persist(DB, users...) + updateCount, uids
}

func RawPersistIUsers(DB *mgo.Collection, data []byte) (int, []string) {
	iuser := SearchIUsers(data)
	users := make([]interface{}, 0, len(iuser))
	uids := make([]string, 0, len(iuser))
	for _, iu := range iuser {
		bys := dbu.I2JsonBytes(iu)
		if nil != bys {
			uid := search.SearchSValue(bys, "id")
			selector := bson.M{
				"id": uid,
			}
			if n, err := DB.Find(selector).Count(); err == nil && n == 0 {
				users = append(users, iu)
				log.Printf("%s >>>>>>>>>>>>>>> {DB}", uid)
			} else {
				log.Printf("%s has existed.", uid)
			}
			uids = append(uids, uid)
		}
	}
	err := DB.Insert(users...)
	if logu.CheckErr(err) {
		return -1, uids
	}
	return len(users), uids
}

func PersistIOnlineStatuses(DB *dbu.Collection, data []byte) (int, int) {
	ionlines := SearchIOnlieStatuses(data)

	online_statuses := make([]interface{}, 0, len(ionlines))
	now := Now()
	for _, ion := range ionlines {
		online_status := OnlineStatus{Time: now, IOnlineStatus: ion}
		online_statuses = append(online_statuses, online_status)
	}
	return Persist(DB, online_statuses...), OnlineCount(ionlines)
}

func UpdatePersistIOnlineStatuses(DB *dbu.Collection, data []byte) (int, int) {
	ionlines := SearchIOnlieStatuses(data)
	ret_count := 0
	// online_statuses := make([]interface{}, 0, len(ionlines))
	now := NowUnix()
	for _, ion := range ionlines {
		online_status := DOnlineStatus{Time: now, IOnlineStatus: ion}
		if iuser, ok := ion.(map[string]interface{}); ok {
			if uid, ok := iuser["uid"]; ok {
				selector := bson.M{"online_status.uid": uid}
				n := DB.Upsert(selector, online_status)
				if n > 0 {
					ret_count++
				}
			}
		}
		// online_statuses = append(online_statuses, online_status)
	}
	return ret_count, len(ionlines)
}

func VisitCountStat(DB *dbu.Collection, data []byte) (int, int) {
	ionlines := SearchIOnlieStatuses(data)
	ret_count := 0

	for _, ion := range ionlines {
		uid := search.ISearchSValue(ion, "uid", []string{}...)
		status := fmt.Sprintf("%v", search.ISearchI(ion, "status", []string{}...))
		if len(status) <= 0 {
			status = "0"
		}
		if len(uid) <= 0 {
			continue
		}
		selector := bson.M{"uid": uid, "status": status}
		var old ViCount
		err := DB.C.Find(selector).One(&old)
		if logu.CheckErr(err) {
			old.UID = fmt.Sprintf("%v", uid)
			old.VCount = 1
		} else {
			old.VCount++
		}
		old.Status = status
		n := DB.Upsert(selector, old)
		if n > 0 {
			ret_count++
		}
		// online_statuses = append(online_statuses, online_status)
	}
	return ret_count, len(ionlines)
}
