package main

import (
	"fmt"
	"github.com/Unknwon/macaron"
	"github.com/shaalx/membership/db"
	"github.com/shaalx/membership/dbu"
	"github.com/shaalx/membership/logu"
	"github.com/shaalx/membership/search"
	"github.com/shaalx/membership/u"
	"html/template"
	"labix.org/v2/mgo/bson"
	"log"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"time"
)

var (
	MgoDB   = dbu.NewMgoDB(dbu.Conn())
	usersC  = MgoDB.GetCollection([]string{"lEyTj8hYrUIKgMfi", "users"}...)
	onlineC = MgoDB.GetCollection([]string{"lEyTj8hYrUIKgMfi", "online"}...)
	vcountC = MgoDB.GetCollection([]string{"lEyTj8hYrUIKgMfi", "vcount"}...)
	or      = false
	update  = true
	page    int
)

func init() {
	page = 1
}

func main() {
	go v4()
	m := macaron.Classic()
	m.Use(macaron.Renderer(macaron.RenderOptions{
		Funcs: []template.FuncMap{
			{
				"Degree": func(degree string) string {
					switch degree {
					case "0":
						return "本科"
					case "1":
						return "硕士"
					case "2":
						return "博士"
					}
					return fmt.Sprintf("Unknown degree: %s", degree)
				},
			},
		}}))
	m.Use(macaron.Static("public",
		macaron.StaticOptions{
			// 请求静态资源时的 URL 前缀，默认没有前缀
			Prefix: "public",
			// 禁止记录静态资源路由日志，默认为不禁止记录
			SkipLogging: true,
			// 当请求目录时的默认索引文件，默认为 "index.html"
			IndexFile: "index.html",
			// 用于返回自定义过期响应头，不能为不设置
			// https://developers.google.com/speed/docs/insights/LeverageBrowserCaching
			Expires: func() string { return "max-age=0" },
		}))
	m.Get("/", index)
	m.Get("/db", dbIndex)
	m.Get("/previous", Previous)
	m.Get("/next", Next)
	m.Get("/switch", _switch)
	m.Get("/switchUpdate", switchUpdate)
	m.Get("/all_count", all_count)
	m.Get("/online_count", online_count)
	m.Get("/statistics", statistics)
	m.Get("/online_stat", online_statistics)
	m.Get("/vcount", vcount)
	m.Get("/dropvcount", dropvcount)
	m.Get("/search", searchName)
	m.Get("/search2", searchName2)
	m.Get("/upsert/:uid", upsert)

	m.Run(80)
}
func dbIndex(rw http.ResponseWriter, req *http.Request) {
	uri := req.RequestURI
	URI, err := url.Parse(uri)
	if !logu.CheckErr(err) {
		pageStr := URI.Query().Get("page")
		if len(pageStr) <= 0 {
			page = 1
		} else {
			page64, err := strconv.ParseInt(pageStr, 10, 0)
			if logu.CheckErr(err) {
				page = 1
			} else {
				page = int(page64)
			}
		}
	}
	if page <= 0 {
		page = 1
	}
	page -= 1
	pageSize := 10
	count := all_countInt()
	start := page * pageSize
	if page*pageSize >= count {
		page = count / pageSize
		if pageSize <= count {
			start = 0
		} else {
			start = (page - 1) * pageSize
		}
	}
	if start < 0 {
		start = 0
	}
	page += 1
	var users []interface{}
	err = usersC.C.Find(nil).Skip(start).Limit(pageSize).All(&users)
	if !logu.CheckErr(err) {
		b := dbu.I2JsonBytes(users)
		rw.Write(b)
		// log.Println(users)
	}
}
func index(ctx *macaron.Context) {
	uri := ctx.Req.RequestURI
	URI, err := url.Parse(uri)
	if !logu.CheckErr(err) {
		pageStr := URI.Query().Get("page")
		if len(pageStr) <= 0 {
			page = 1
		} else {
			page64, err := strconv.ParseInt(pageStr, 10, 0)
			if logu.CheckErr(err) {
				page = 1
			} else {
				page = int(page64)
			}
		}
	}
	if page <= 0 {
		page = 1
	}
	page -= 1
	pageSize := 10
	count := all_countInt()
	start := page * pageSize
	if page*pageSize >= count {
		page = count / pageSize
		if pageSize <= count {
			start = 0
		} else {
			start = (page - 1) * pageSize
		}
	}
	if start < 0 {
		start = 0
	}
	page += 1
	var users []interface{}
	err = usersC.C.Find(nil).Skip(start).Limit(pageSize).All(&users)
	if !logu.CheckErr(err) {
		ctx.Data["users"] = users
		ctx.Data["fetch"] = or
		ctx.Data["update"] = update
		ctx.Data["all_count"] = fmt.Sprintf("%v", count)
		ctx.Data["online_count"] = online_count()
		ctx.Data["Previous"] = template.HTMLEscapeString("<h1>Previous</h1>")
		ctx.Data["Next"] = template.HTML(fmt.Sprintf(`<a href="/?page=%d><h1>>>></h1></a>">`, page+1))
		ctx.HTML(200, "index")
	}
}

func upsert(ctn *macaron.Context) interface{} {
	uid := ctn.Params(":uid")
	// return dbu.I2JsonBytes(_upsert(uid))
	iuser := _upsert(uid)
	return search.ISearchSValue(iuser, "avatar_large", []string{}...)
}

func _upsert(uid string) interface{} {
	_url := fmt.Sprintf("=%s", uid)
	bys := u.Fetch(_url)
	user := search.SearchI(bys, "user", []string{}...)
	selector := bson.M{"id": uid}
	n := usersC.Upsert(selector, user)
	log.Println(n)
	return user
}

func Previous(ctn *macaron.Context) {
	page -= 1
	fmt.Println(page)
	ctn.Redirect(fmt.Sprintf("/?page=%d", page))
}

func Next(ctn *macaron.Context) {
	page += 1
	fmt.Println(page)
	ctn.Redirect(fmt.Sprintf("/?page=%d", page))
}

func _switch(ctx *macaron.Context) string {
	or = !or
	if or {
		return "true"
	}
	return "false"
}

func switchUpdate(ctx *macaron.Context) string {
	update = !update
	if update {
		return "true"
	}
	return "false"
}

func all_count() string {
	n := fmt.Sprintf("%v", usersC.Count(nil))
	return n
}

func all_countInt() int {
	return usersC.Count(nil)
}

func online_count() string {
	selector := bson.M{"online_status.status": "2"}
	var ret []interface{}
	onlineC.C.Find(selector).Distinct("online_status.uid", &ret)
	return fmt.Sprintf("%v", len(ret))
}

func online_status(iuids ...interface{}) int {
	_url2 := "son?uids="
	uids := make([]string, 0, len(iuids))
	for _, iuid := range iuids {
		uid := fmt.Sprintf("%v", iuid)
		uids = append(uids, uid)
	}
	juids := strings.Join(uids, ",")
	fmt.Println(juids)
	online_status_url := _url2 + juids
	bys := u.Fetch(online_status_url)
	fmt.Println(string(bys))
	ionline_users := db.SearchIOnlieStatuses(bys)
	return db.OnlineCount(ionline_users)
}

func distinct_uids() []interface{} {
	var ret []interface{}
	err := onlineC.C.Find(nil).Distinct("online_status.uid", &ret)
	if logu.CheckErr(err) {
		return nil
	}
	return ret
}

func v4() {
	_url2 := "45b0b56f9"
	_url := "r_page=24"
	for {
		if or {
			bys := u.Fetch(_url)
			n, uids := db.PersistIUsers(MgoDB.GetCollection([]string{"lEyTj8hYrUIKgMfi", "users"}...), bys, update)
			log.Println(n)

			_ = uids
			juids := strings.Join(uids, ",")
			online_status_url := _url2 + juids
			bys = u.Fetch(online_status_url)
			all, online_count := db.PersistIOnlineStatuses(MgoDB.GetCollection([]string{"lEyTj8hYrUIKgMfi", "online"}...), bys)
			go db.VisitCountStat(MgoDB.GetCollection([]string{"lEyTj8hYrUIKgMfi", "vcount"}...), bys)
			log.Printf("%d / %d", online_count, all)
		}

		heart_bengbengbeng := u.Heart()
		log.Printf(" %d sec later...", heart_bengbengbeng/1000000000)
		time.Sleep(heart_bengbengbeng)
	}
}

func statistics(ctx *macaron.Context) {
	duids := db.DistinctUids(onlineC)
	status0_, status1_, status2_ := db.OnlineUids(duids...)
	status0 := make([]interface{}, 0, len(status0_))
	status1 := make([]interface{}, 0, len(status1_))
	status2 := make([]interface{}, 0, len(status2_))
	var selector bson.M
	for _, iuid := range status0_ {
		uid := strings.TrimSpace(fmt.Sprintf("%v", iuid))
		selector = bson.M{"id": uid}
		user := usersC.ISelectOne(selector)
		if nil != user {
			status0 = append(status0, user)
		}
	}
	for _, iuid := range status1_ {
		uid := strings.TrimSpace(fmt.Sprintf("%v", iuid))
		selector = bson.M{"id": uid}
		user := usersC.ISelectOne(selector)
		if nil != user {
			status1 = append(status1, user)
		}
	}
	for i, iuid := range status2_ {
		if i > 20 {
			break
		}
		uid := strings.TrimSpace(fmt.Sprintf("%v", iuid))
		selector = bson.M{"id": uid}
		user := usersC.ISelectOne(selector)
		if nil != user {
			status2 = append(status2, user)
		}
	}

	ctx.Data["status0"] = status0
	ctx.Data["status1"] = status1
	ctx.Data["status2"] = status2

	all_count_ := all_countInt()
	status_len := len(duids)
	missed := all_count_ - status_len

	ctx.Data["status_len"] = status_len
	ctx.Data["missed_len"] = missed
	ctx.Data["status0_len"] = len(status0)
	ctx.Data["status1_len"] = len(status1)
	ctx.Data["status2_len"] = len(status2_)

	ctx.Data["all_count"] = all_count_
	ctx.Data["fetch"] = or
	ctx.Data["update"] = update

	ctx.HTML(200, "stat")
}

func online_statistics(ctx *macaron.Context) {
	all := onlineC.ISelect(nil)
	status0_m := make(map[string]int64, 0)
	status1_m := make(map[string]int64, 0)
	status2_m := make(map[string]int64, 0)

	for _, iuser := range all {
		bys := dbu.I2JsonBytes(iuser)
		uid := search.SearchSValue(bys, "uid", []string{"online_status"}...)
		status := search.SearchSValue(bys, "status", []string{"online_status"}...)
		switch status {
		case "0":
			status0_m[uid]++
		case "1":
			status1_m[uid]++
		case "2":
			status2_m[uid]++
		}
	}

	var status0, status1, status2 SVisitTimes
	status0 = m2s_sorted(status0_m)
	status1 = m2s_sorted(status1_m)
	status2 = m2s_sorted(status2_m)

	ctx.Data["status0_len"] = len(status0_m)
	ctx.Data["status1_len"] = len(status1_m)
	ctx.Data["status2_len"] = len(status2_m)

	var end0, end1, end2 int
	len0 := len(status0)
	len1 := len(status1)
	len2 := len(status2)
	if 21 > len0 {
		end0 = len0
	} else {
		end0 = 21
	}
	if 21 > len1 {
		end1 = len1
	} else {
		end1 = 21
	}
	if 21 > len2 {
		end2 = len2
	} else {
		end2 = 21
	}
	ctx.Data["status0"] = status0[:end0]
	ctx.Data["status1"] = status1[:end1]
	ctx.Data["status2"] = status2[:end2]

	ctx.Data["all_count"] = len(db.DistinctUids(onlineC))
	ctx.Data["fetch"] = or
	ctx.Data["update"] = update

	ctx.HTML(200, "onlinestat")
}

func m2s_sorted(m map[string]int64) []VisitTimes {
	s := make(SVisitTimes, 0, len(m))
	for k, v := range m {
		s = append(s, NewVisitTimes(k, v))
	}
	sort.Sort(s)
	return s
}

type VisitTimes struct {
	UID   string `json:"uid"`
	Times int64  `json:"times"`
}

func NewVisitTimes(uid string, times int64) VisitTimes {
	return VisitTimes{
		UID:   uid,
		Times: times,
	}
}

type SVisitTimes []VisitTimes

func (s SVisitTimes) Len() int {
	return len(s)
}

func (s SVisitTimes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s SVisitTimes) Less(i, j int) bool {
	return s[i].Times > s[j].Times
}

func vcount(ctn *macaron.Context) {
	var vcounts []db.ViCount
	var mvvcounts []*db.VViCount
	var fvvcounts []*db.VViCount
	err := vcountC.C.Find(nil).Sort("-vcount").Limit(41).All(&vcounts)
	if !logu.CheckErr(err) {
		var selector_u bson.M
		for _, vc := range vcounts {
			selector_u = bson.M{"id": vc.UID}
			var iu interface{}
			err = usersC.C.Find(selector_u).One(&iu)
			if !logu.CheckErr(err) {
				avatar := search.ISearchSValue(iu, "avatar_hd", []string{}...)
				gender := search.ISearchSValue(iu, "gender", []string{}...)
				name := search.ISearchSValue(iu, "name", []string{}...)
				nickname := search.ISearchSValue(iu, "nickname", []string{}...)
				vvcount := db.NewVViCount(vc.UID, vc.VCount, vc.Status, avatar, name, nickname)
				if strings.EqualFold(gender, "0") {
					if len(fvvcounts) >= 10 {
						continue
					}
					fvvcounts = append(fvvcounts, vvcount)
				} else {
					if len(mvvcounts) >= 10 {
						continue
					}
					mvvcounts = append(mvvcounts, vvcount)
				}
			}
		}
		ctn.Data["fvvcounts"] = fvvcounts
		ctn.Data["mvvcounts"] = mvvcounts

		ctn.Data["all_count"] = all_count()
		ctn.Data["fetch"] = or
		ctn.Data["update"] = update
		ctn.HTML(200, "vcount")
	}
}

func dropvcount(ctx *macaron.Context) string {
	or = true
	err := vcountC.C.DropCollection()
	ctx.Redirect("/vcount")
	if logu.CheckErr(err) {
		return "false"
	}
	return "true"
}

func searchName(ctx *macaron.Context) {

	uri := ctx.Req.RequestURI
	URI, err := url.Parse(uri)
	if logu.CheckErr(err) {
		return
	}
	urlQuery := URI.Query()

	searchD := urlQuery.Get("searchD")
	fmt.Println(searchD)
	searchNameStr := urlQuery.Get("searchName")
	searchPage := 1
	query := bson.M{"$or": []bson.M{bson.M{"name": bson.RegEx{searchNameStr, "."}}, bson.M{"nickname": bson.RegEx{searchNameStr, "."}}}}

	pageStr := urlQuery.Get("searchPage")
	if len(pageStr) <= 0 {
		searchPage = 1
	} else {
		page64, err := strconv.ParseInt(pageStr, 10, 0)
		if logu.CheckErr(err) {
			searchPage = 1
		} else {
			if strings.EqualFold(searchD, "next") {
				searchPage = int(page64) + 1
			} else {
				searchPage = int(page64) - 1
			}
		}
	}
	if searchPage <= 0 {
		searchPage = 1
	}
	searchPage -= 1
	pageSize := 10
	count := usersC.Count(query)
	start := searchPage * pageSize
	if searchPage*pageSize >= count {
		searchPage = count / pageSize
		if pageSize <= count {
			start = 0
		} else {
			start = (searchPage - 1) * pageSize
		}
	}
	if start < 0 {
		start = 0
	}
	searchPage += 1
	// fmt.Println(count, searchPage, count/pageSize)
	var users []interface{}
	err = usersC.C.Find(query).Skip(start).Limit(pageSize).All(&users)
	if !logu.CheckErr(err) {
		ctx.Data["users"] = users
	}

	ctx.Data["searchName"] = searchNameStr
	ctx.Data["searchPage"] = searchPage
	ctx.Data["all_count"] = all_count()
	ctx.Data["fetch"] = or
	ctx.Data["update"] = update
	ctx.HTML(200, "search")

}

func searchName2(ctx *macaron.Context) {
	uri := ctx.Req.RequestURI
	URI, err := url.Parse(uri)
	if logu.CheckErr(err) {
		return
	}
	urlQuery := URI.Query()

	searchNameStr := urlQuery.Get("searchName")
	searchPage := 1
	query := bson.M{"$or": []bson.M{bson.M{"name": bson.RegEx{searchNameStr, "."}}, bson.M{"nickname": bson.RegEx{searchNameStr, "."}}}}

	pageStr := urlQuery.Get("page")
	if len(pageStr) <= 0 {
		searchPage = 1
	} else {
		page64, err := strconv.ParseInt(pageStr, 10, 0)
		if logu.CheckErr(err) {
			searchPage = 1
		} else {
			searchPage = int(page64) + 1
		}
	}
	if searchPage <= 0 {
		searchPage = 1
	}
	searchPage -= 1
	pageSize := 10
	count := usersC.Count(query)
	start := searchPage * pageSize
	if searchPage*pageSize >= count {
		searchPage = count / pageSize
		if pageSize <= count {
			start = 0
		} else {
			start = (searchPage - 1) * pageSize
		}
	}
	if start < 0 {
		start = 0
	}
	searchPage += 1

	var users []interface{}
	err = usersC.C.Find(query).Skip(start).Limit(pageSize).All(&users)
	if !logu.CheckErr(err) {
		ctx.Data["users"] = users
	}

	ctx.Data["searchName"] = searchNameStr
	ctx.Data["searchPage"] = searchPage + 1
	ctx.Data["all_count"] = all_count()
	ctx.Data["fetch"] = or
	ctx.Data["update"] = update
	ctx.HTML(200, "search2")

}
