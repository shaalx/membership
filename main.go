package main

import (
	"net/http"

	"github.com/shaalx/merbership/db"
	"github.com/shaalx/merbership/dbu"
	"github.com/shaalx/merbership/logu"
	"github.com/shaalx/merbership/pkg3/httplib"
	"github.com/shaalx/merbership/u"
	"log"
	// "strings"
	"time"
)

func search(rw http.ResponseWriter, req *http.Request) {

}

func main() {
	v4()
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
	// MgoDB = dbu.NewMgoDB(dbu.Conn())
	peopleC = dbu.RawMgoDB()
)

func v4() {
	// _url2 := "https://api.simplr.cn/0.1/user/online_status.json?uids=" //557471f7a341140630d4d319%2C551473eda34114331d3bfaf5%2C55560f1da3411422e127ca91%2C5557411da341140b6f82663a%2C552cb566a34114109b2925e2%2C55138c1ca3411440863bfbda%2C5513c8f2a3411478a13bf3bf%2C55140832a34114196f3bf27b%2C5563d800bd4b873a164155fd%2C55142b6ca3411428603bf5a2%2C550db261a341143b0ae91507%2C5549f272a34114481e27cda9%2C5566bc63a3411429a9f6da87%2C555bca9ba3411411bd826212%2C555fdbe0bd4b8706478a1c67%2C555dc48dbd4b87204faa60ff%2C5552eeeca341143d6927c76c%2C550eea42a34114649df2a9cd%2C55195b21a341145a8415aa91%2C555d59dfa341143873826ee1%2C555332f8a341145c0e27c89d%2C5539daf9a3411434a96ab8a6%2C5552ceb5a3411431fa27bcc5%2C5518da59a341142d171598ab&identifier=8e65b14e-338b-4191-a5c3-73e45b0b56f9"
	_url := "https://api.simplr.cn/0.1/discover/filter.json?identifier=8e65b14e-338b-4191-a5c3-73e45b0b56f9&_per_page=24"
	for {
		bys := fetch(_url)
		n, uids := db.RawPersistIUsers(peopleC, bys)
		log.Println(n)

		_ = uids
		// juids := strings.Join(uids, ",")
		// online_status_url := _url2 + juids
		// bys = fetch(online_status_url)
		// all, online_count := db.PersistIOnlineStatuses(MgoDB.GetCollection([]string{"test", "online"}...), bys)
		// log.Printf("%d / %d", online_count, all)
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
