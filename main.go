package main

import (
	"github.com/Unknwon/macaron"

	"fmt"
	// "bytes"
	"labix.org/v2/mgo/bson"
	"net/http"

	"github.com/shaalx/merbership/db"
	"github.com/shaalx/merbership/dbu"
	"github.com/shaalx/merbership/logu"
	"github.com/shaalx/merbership/pkg3/httplib"
	"github.com/shaalx/merbership/u"
	"log"
	"strings"
	"time"
)

var (
	or = false
)

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
	// m.Get("/detail/:uid", detail)
	m.Get("/switch", _switch)
	m.Get("/all_count", all_count)
	m.Get("/online_count", online_count)

	m.Run(80)
}

func index(ctx *macaron.Context) {
	var users []interface{}
	err := usersC.C.Find(nil).Limit(10).All(&users)
	if !logu.CheckErr(err) {
		ctx.Data["users"] = users
		ctx.Data["fetch"] = or
		ctx.Data["all_count"] = all_count()
		ctx.Data["online_count"] = online_count()
		ctx.HTML(200, "index")
	}
}

func _switch(ctx *macaron.Context) {
	or = !or
	ctx.Redirect("/")
}

func all_count() string {
	n := fmt.Sprintf("%v", usersC.Count(nil))
	return n
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
	bys := fetch(online_status_url)
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

func count(rw http.ResponseWriter, req *http.Request) {
	n := usersC.Count(nil)
	rw.Write([]byte(fmt.Sprintf("Count : %d", n)))
}

func all(rw http.ResponseWriter, req *http.Request) {
	ret := usersC.Select(nil)
	rw.Write([]byte(fmt.Sprintf("all\n : %v", ret)))
}

func online_all(rw http.ResponseWriter, req *http.Request) {
	ret := onlineC.Select(nil)
	rw.Write([]byte(fmt.Sprintf("all\n : %v", ret)))
}

func search(rw http.ResponseWriter, req *http.Request) {
	uid := req.URL.Query().Get("uid")
	selector := bson.M{
		"id": uid,
	}
	ret := usersC.Select(selector)
	rw.Write([]byte(fmt.Sprintf("search : \n%v", ret)))
}

func dorun(rw http.ResponseWriter, req *http.Request) {
	or = !or
	fetch := ""
	if or {
		fetch = "will be fetching ..."
	} else {
		fetch = "will be having a rest ."
	}
	rw.Write([]byte(fetch))
}

func old_main() {
	go v4()
	http.HandleFunc("/dorun", dorun)
	http.HandleFunc("/", count)
	http.HandleFunc("/users", all)
	http.HandleFunc("/online", online_all)
	http.HandleFunc("/search", search)
	http.ListenAndServe(":80", nil)
}

func vv1() interface{} {
	// req := httplib.Get("https://api.simplr.cn/0.1/leanchat/signature.json?peerId=5525f5bee4b03381b313a552&identifier=8e65b14e-338b-4191-a5c3-73e45b0b56f9&watchPeerIds=")
	// req := httplib.Get("https://api.simplr.cn/0.1/timeline/notifications.json?identifier=8e65b14e-338b-4191-a5c3-73e45b0b56f9&lastTimestampMs=1433231285170&maxCount=20")
	// req := httplib.Get("https://api.simplr.cn/0.1/auth/login.json?areaId=54fa0e13a341141ad9071274&username=15921911727&password=acec59c80712e9df1428e78fcf04f74458b0f10cb9bb85c981055b79664312e2")

	// req := httplib.Get("https://api.simplr.cn/0.1/album/photos.json?identifier=8e65b14e-338b-4191-a5c3-73e45b0b56f9&uid=5525f5bda341143a4e6a8996&_per_page=8")
	// req := httplib.Get("https://api.simplr.cn/0.1/profile/get.json?identifier=8e65b14e-338b-4191-a5c3-73e45b0b56f9&uid=555530e8a34114438c27bbe7")
	// req := httplib.Get("https://api.simplr.cn/0.1/user/get.json?identifier=8e65b14e-338b-4191-a5c3-73e45b0b56f9&uid=555530e8a34114438c27bbe7")
	// req := httplib.Get("https://api.simplr.cn/0.1/interview/answers.json?identifier=8e65b14e-338b-4191-a5c3-73e45b0b56f9&uid=55149a65a341145f113bee58")
	// req := httplib.Get("https://api.simplr.cn/0.1/user/counts.json?uids=55471927a34114089627c504&identifier=8e65b14e-338b-4191-a5c3-73e45b0b56f9")
	// req := httplib.Get("https://api.simplr.cn/0.1/friendship/show.json?identifier=8e65b14e-338b-4191-a5c3-73e45b0b56f9&uids=5552f388a341143d5427ccdb")

	// req := httplib.Get("https://api.simplr.cn/0.1/user/visitors.json?identifier=8e65b14e-338b-4191-a5c3-73e45b0b56f9&_per_page=18&uid=554af806a34114407627bf4b")
	// req := httplib.Get("https://api.simplr.cn/0.1/timeline/get_discover_timeline.json?identifier=8e65b14e-338b-4191-a5c3-73e45b0b56f9&user_ref=1&_per_page=1")
	// req := httplib.Get("https://api.simplr.cn/0.1/friendship/followers.json?identifier=8e65b14e-338b-4191-a5c3-73e45b0b56f9&_per_page=1&uid=5552f388a341143d5427ccdb")
	// req := httplib.Get("https://api.simplr.cn/0.1/friendship/friends.json?identifier=8e65b14e-338b-4191-a5c3-73e45b0b56f9&_per_page=1")
	// req := httplib.Get("https://api.simplr.cn/0.1/timeline/get_discover_timeline.json?identifier=8e65b14e-338b-4191-a5c3-73e45b0b56f9&user_ref=1&_per_page=10")
	// req := httplib.Get("https://api.simplr.cn/0.1/interview/answers.json?identifier=8e65b14e-338b-4191-a5c3-73e45b0b56f9")
	// req := httplib.Get("https://api.simplr.cn/0.1/timeline/get_discover_timeline.json?identifier=8e65b14e-338b-4191-a5c3-73e45b0b56f9&user_ref=1&_per_page=1")

	// req := httplib.Get("https://api.simplr.cn/0.1/user/online_status.json?uids=5563d800bd4b873a164155fd&identifier=8e65b14e-338b-4191-a5c3-73e45b0b56f9")
	req := httplib.Get("https://api.simplr.cn/0.1/discover/filter.json?identifier=8e65b14e-338b-4191-a5c3-73e45b0b56f9&_per_page=2")
	// req := httplib.Get("https://api.simplr.cn/0.1/discover/filter.json?identifier=8e65b14e-338b-4191-a5c3-73e45b0b56f9")
	// req := httplib.Get("https://api.simplr.cn/0.1/discover/filter.json?identifier=8e65b14e-338b-4191-a5c3-73e45b0b56f9&_per_page=240")
	// req := httplib.Get("https://api.simplr.cn/0.1/discover/filter.json?departmentId=54fa0e13a341141ad9071261&identifier=8e65b14e-338b-4191-a5c3-73e45b0b56f9&gender=0&schoolId=54fa0e13a341141ad9071254&degree=1&_per_page=240&grade=2014")

	// req := httplib.Get("https://api.simplr.cn/0.1/auth/refresh_token.json?identifier=8e65b14e-338b-4191-a5c3-73e45b0b56f9&refresh_token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJoYXNoIjoiOTE5ZjEzZWQ5YjRhOWQ1NmIxN2Y4MDRhY2U0ODViOTQzMTA1ODQyOSIsImlkIjoiNTU2ZWMzYzRhMzQxMTQzNjExZjZkNDQwIn0.SF0vsWxH_h9jK0RNfkV51yK2jz4XP68zs9wRVu5nhqg")
	// req := httplib.Get("https://api.simplr.cn/0.1/public/top_schools.json")
	// req := httplib.Get("https://api.simplr.cn/0.1/public/school_departments.json?schoolId=55325319a341147b16db72b3")

	req.Header("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJub25jZSI6IjgzOTlmMjJmNjg5MzUzNGQ3OGE5IiwidWlkIjoiNTUyNWY1YmRhMzQxMTQzYTRlNmE4OTk2IiwiZXhwIjoxNDM1OTAyMjExfQ.5LNYiRI6SiGbLsDmLbsPc4X6JrhyDh1X2_5kVFV4VMg")
	req.Header("Host", "api.simplr.cn")
	// log.Println(req.String())
	var v interface{}
	log.Println(req.ToJson(&v))
	// log.Println(v)
	return v
}

var (
	MgoDB   = dbu.NewMgoDB(dbu.Conn())
	usersC  = MgoDB.GetCollection([]string{"lEyTj8hYrUIKgMfi", "users"}...)
	onlineC = MgoDB.GetCollection([]string{"lEyTj8hYrUIKgMfi", "online"}...)
	// usersC = dbu.RawMgoDB()
)

func v4() {
	_url2 := "https://api.simplr.cn/0.1/user/online_status.json?uids=" //557471f7a341140630d4d319%2C551473eda34114331d3bfaf5%2C55560f1da3411422e127ca91%2C5557411da341140b6f82663a%2C552cb566a34114109b2925e2%2C55138c1ca3411440863bfbda%2C5513c8f2a3411478a13bf3bf%2C55140832a34114196f3bf27b%2C5563d800bd4b873a164155fd%2C55142b6ca3411428603bf5a2%2C550db261a341143b0ae91507%2C5549f272a34114481e27cda9%2C5566bc63a3411429a9f6da87%2C555bca9ba3411411bd826212%2C555fdbe0bd4b8706478a1c67%2C555dc48dbd4b87204faa60ff%2C5552eeeca341143d6927c76c%2C550eea42a34114649df2a9cd%2C55195b21a341145a8415aa91%2C555d59dfa341143873826ee1%2C555332f8a341145c0e27c89d%2C5539daf9a3411434a96ab8a6%2C5552ceb5a3411431fa27bcc5%2C5518da59a341142d171598ab&identifier=8e65b14e-338b-4191-a5c3-73e45b0b56f9"
	_url := "https://api.simplr.cn/0.1/discover/filter.json?identifier=8e65b14e-338b-4191-a5c3-73e45b0b56f9&_per_page=24"
	for {
		if or {
			bys := fetch(_url)
			n, uids := db.PersistIUsers(MgoDB.GetCollection([]string{"lEyTj8hYrUIKgMfi", "users"}...), bys)
			log.Println(n)

			_ = uids
			juids := strings.Join(uids, ",")
			online_status_url := _url2 + juids
			bys = fetch(online_status_url)
			all, online_count := db.PersistIOnlineStatuses(MgoDB.GetCollection([]string{"lEyTj8hYrUIKgMfi", "online"}...), bys)
			log.Printf("%d / %d", online_count, all)
		}

		heart_bengbengbeng := u.Heart()
		log.Printf(" %d sec later...", heart_bengbengbeng/1000000000)
		time.Sleep(heart_bengbengbeng)
	}
}

func fetch(_url string) []byte {
	req := httplib.Get(_url)
	req.Header("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJub25jZSI6IjgzOTlmMjJmNjg5MzUzNGQ3OGE5IiwidWlkIjoiNTUyNWY1YmRhMzQxMTQzYTRlNmE4OTk2IiwiZXhwIjoxNDM1OTAyMjExfQ.5LNYiRI6SiGbLsDmLbsPc4X6JrhyDh1X2_5kVFV4VMg")
	req.Header("Host", "api.simplr.cn")
	bys, err := req.Bytes()
	if logu.CheckErr(err) {
		return nil
	}
	return bys
}
