package u

import (
	"github.com/shaalx/merbership/logu"
	"github.com/shaalx/merbership/pkg3/httplib"
)

func Fetch(_url string) []byte {
	req := httplib.Get(_url)
	req.Header("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJub25jZSI6IjgzOTlmMjJmNjg5MzUzNGQ3OGE5IiwidWlkIjoiNTUyNWY1YmRhMzQxMTQzYTRlNmE4OTk2IiwiZXhwIjoxNDM1OTAyMjExfQ.5LNYiRI6SiGbLsDmLbsPc4X6JrhyDh1X2_5kVFV4VMg")
	req.Header("Host", "api.simplr.cn")
	bys, err := req.Bytes()
	if logu.CheckErr(err) {
		return nil
	}
	return bys
}
