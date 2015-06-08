package db

import (
	// "fmt"
	"testing"
	// "time"
)

func TestNow(t *testing.T) {
	now := Now()
	t.Log(now)
	// st := "2006-01-02 08:00"
	// fmt.Println(st[5 : len(st)-1-5])
	// time3, _ := time.Parse("2006-01-02 15:04", st)
	// loc, _ := time.LoadLocation("Local")
	// time2, _ := time.ParseInLocation("2006-01-02 15:04", st, loc)
	// fmt.Println(time3)
	// fmt.Println(time2)
}
