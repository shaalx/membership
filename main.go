package main

import (
	"fmt"
	"github.com/Unknwon/macaron"
	"github.com/shaalx/merbership/db"
	"github.com/shaalx/merbership/dbu"
	"github.com/shaalx/merbership/logu"
	"github.com/shaalx/merbership/u"
	"html/template"
	"labix.org/v2/mgo/bson"
	"log"
	"net/url"
	"strconv"
	"strings"
	"time"
)

var (
	MgoDB   = dbu.NewMgoDB(dbu.Conn())
	usersC  = MgoDB.GetCollection([]string{"lEyTj8hYrUIKgMfi", "users"}...)
	onlineC = MgoDB.GetCollection([]string{"lEyTj8hYrUIKgMfi", "online"}...)
	// usersC = dbu.RawMgoDB()
	or     = false
	update = true
	page   int
)

func init() {
	page = 1
}

func main() {
	go v4()
	m := macaron.Classic()
	m.Use(macaron.Renderer())
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
	m.Get("/previous", Previous)
	m.Get("/next", Next)
	m.Get("/switch", _switch)
	m.Get("/switchUpdate", switchUpdate)
	m.Get("/all_count", all_count)
	m.Get("/online_count", online_count)
	m.Get("/statistics", statistics)

	m.Run(80)
}

func index(ctx *macaron.Context) {
	// var users []interface{}
	// err := usersC.C.Find(nil).Limit(10).All(&users)
	// if !logu.CheckErr(err) {
	// 	ctx.Data["users"] = users
	// 	ctx.Data["fetch"] = or
	// 	ctx.Data["update"] = update
	// 	ctx.Data["all_count"] = all_count()
	// 	ctx.Data["online_count"] = online_count()
	// 	ctx.HTML(200, "index")
	// }

	// page = 1
	uri := ctx.Req.RequestURI
	URI, err := url.Parse(uri)
	if !logu.CheckErr(err) {
		pageStr := URI.Query().Get("page")
		page64, err := strconv.ParseInt(pageStr, 10, 0)
		if logu.CheckErr(err) {
			page = 1
		} else {
			page = int(page64)
		}
	}
	if page <= 0 {
		page = 1
	}
	page -= 1
	pageSize := 10
	count := all_countInt()
	end := (page + 1) * pageSize
	if page*pageSize >= count {
		page -= 1
		end = count
	}
	start := page * pageSize
	page += 1
	var users []interface{}
	err = usersC.C.Find(nil).Limit(end).All(&users)
	if !logu.CheckErr(err) {
		ctx.Data["users"] = users[start:]
		ctx.Data["fetch"] = or
		ctx.Data["update"] = update
		ctx.Data["all_count"] = fmt.Sprintf("%v", count)
		ctx.Data["online_count"] = online_count()
		// ctx.Data["Previous"] = template.HTML(fmt.Sprintf(`<a href="/?page=%d><h1><<<</h1></a>">`, page-1))
		ctx.Data["Previous"] = template.HTMLEscapeString("<h1>Previous</h1>")
		ctx.Data["Next"] = template.HTML(fmt.Sprintf(`<a href="/?page=%d><h1>>>></h1></a>">`, page+1))
		ctx.HTML(200, "index")
	}
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
	// uids := distinct_uids()
	// on := online_status(uids)
	// return fmt.Sprintf("%d/%d", on, len(uids))
}

func online_status(iuids ...interface{}) int {
	_url2 := "https://api.simplr.cn/0.1/user/online_status.json?uids="
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
	_url2 := "https://api.simplr.cn/0.1/user/online_status.json?uids=" //557471f7a341140630d4d319%2C551473eda34114331d3bfaf5%2C55560f1da3411422e127ca91%2C5557411da341140b6f82663a%2C552cb566a34114109b2925e2%2C55138c1ca3411440863bfbda%2C5513c8f2a3411478a13bf3bf%2C55140832a34114196f3bf27b%2C5563d800bd4b873a164155fd%2C55142b6ca3411428603bf5a2%2C550db261a341143b0ae91507%2C5549f272a34114481e27cda9%2C5566bc63a3411429a9f6da87%2C555bca9ba3411411bd826212%2C555fdbe0bd4b8706478a1c67%2C555dc48dbd4b87204faa60ff%2C5552eeeca341143d6927c76c%2C550eea42a34114649df2a9cd%2C55195b21a341145a8415aa91%2C555d59dfa341143873826ee1%2C555332f8a341145c0e27c89d%2C5539daf9a3411434a96ab8a6%2C5552ceb5a3411431fa27bcc5%2C5518da59a341142d171598ab&identifier=8e65b14e-338b-4191-a5c3-73e45b0b56f9"
	_url := "https://api.simplr.cn/0.1/discover/filter.json?identifier=8e65b14e-338b-4191-a5c3-73e45b0b56f9&_per_page=24"
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
			log.Printf("%d / %d", online_count, all)
		}

		heart_bengbengbeng := u.Heart()
		log.Printf(" %d sec later...", heart_bengbengbeng/1000000000)
		time.Sleep(heart_bengbengbeng)
	}
}

func statistics(ctx *macaron.Context) {
	duids := db.DistinctUids(onlineC)
	status0, status1, status2 := db.OnlineUids(duids...)
	ctx.Data["status0"] = status0
	ctx.Data["status1"] = status1
	ctx.Data["status2"] = status2

	ctx.Data["all_count"] = len(duids)
	ctx.Data["status0_len"] = len(status0)
	ctx.Data["status1_len"] = len(status1)
	ctx.Data["status2_len"] = len(status2)

	ctx.Data["fetch"] = or
	ctx.Data["update"] = update

	ctx.HTML(200, "stat")
}
