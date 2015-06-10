package db

import (
	// "fmt"
	"github.com/shaalx/merbership/dbu"
	"testing"
	// "time"
)

var (
	MgoDB = dbu.NewMgoDB(dbu.Conn())
	// usersC  = MgoDB.GetCollection([]string{"lEyTj8hYrUIKgMfi", "users"}...)
	onlineC = MgoDB.GetCollection([]string{"lEyTj8hYrUIKgMfi", "online"}...)
	// usersC = dbu.RawMgoDB()
)

// func TestNow(t *testing.T) {
// 	now := Now()
// 	t.Log(now)
// 	// st := "2006-01-02 08:00"
// 	// fmt.Println(st[5 : len(st)-1-5])
// 	// time3, _ := time.Parse("2006-01-02 15:04", st)
// 	// loc, _ := time.LoadLocation("Local")
// 	// time2, _ := time.ParseInLocation("2006-01-02 15:04", st, loc)
// 	// fmt.Println(time3)
// 	// fmt.Println(time2)
// }

func TestDistinctUids(t *testing.T) {
	uids := DistinctUids(onlineC)
	// t.Log(uids)

	status0, status1, status2 := OnlineUids(uids...)
	// t.Log(status0, status1, status2)
	t.Logf("%d  | %d | %d", len(status0), len(status1), len(status2))

	dcount := len(uids)
	t.Log(dcount)
}
