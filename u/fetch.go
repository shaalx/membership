package u

import (
	"github.com/shaalx/membership/logu"
	"github.com/shaalx/membership/pkg3/httplib"
)

func Fetch(_url string) []byte {
	req := httplib.Get(_url)
	req.Header("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJub25jZSI6IjA2NWYyODEzMmYwY2NkN2E0NTNiIiwidWlkIjoiNTUyNWY1YmRhMzQxMTQzYTRlNmE4OTk2IiwiZXhwIjoxNDM4NDE3OTgwfQ.5Uqb1OAi1c0_WUSl1DcdlEQ4-9ReLuaHTv1EQkIxEHE")
	req.Header("Host", "api.simplr.cn")
	bys, err := req.Bytes()
	if logu.CheckErr(err) {
		return nil
	}
	return bys
}

func JsonFetch(_url string) interface{} {
	req := httplib.Get(_url)
	req.Header("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJub25jZSI6IjA2NWYyODEzMmYwY2NkN2E0NTNiIiwidWlkIjoiNTUyNWY1YmRhMzQxMTQzYTRlNmE4OTk2IiwiZXhwIjoxNDM4NDE3OTgwfQ.5Uqb1OAi1c0_WUSl1DcdlEQ4-9ReLuaHTv1EQkIxEHE")
	req.Header("Host", "api.simplr.cn")
	var ret interface{}
	err := req.ToJson(&ret)
	if logu.CheckErr(err) {
		return nil
	}
	return ret
}
