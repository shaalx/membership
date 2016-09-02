package u

import (
	"github.com/toukii/membership/logu"
	"github.com/toukii/membership/pkg3/httplib"
)

func Fetch(_url string) []byte {
	req := httplib.Get(_url)
	req.Header("Authorization", "")
	req.Header("Host", "")
	bys, err := req.Bytes()
	if logu.CheckErr(err) {
		return nil
	}
	return bys
}

func JsonFetch(_url string) interface{} {
	req := httplib.Get(_url)
	req.Header("Authorization", "")
	req.Header("Host", "")
	var ret interface{}
	err := req.ToJson(&ret)
	if logu.CheckErr(err) {
		return nil
	}
	return ret
}
