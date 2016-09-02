package search

import (
	"github.com/toukii/membership/logu"
	sjson "github.com/toukii/membership/pkg3/go-simplejson"
)

/*
* 获取某路径的数组
 */
func SearchArray(data []byte, key string, path ...string) []interface{} {
	if data == nil {
		return nil
	}
	js, err := sjson.NewJson(data)
	if logu.CheckErr(err) {
		return nil
	}
	js = js.GetPath(path...).Get(key)
	ary, err := js.Array()
	if logu.CheckErr(err) {
		return nil
	}
	return ary
}

/*
* 获取路径下的数组，只有数组
 */
func SearchArrays(data []byte, path ...string) []interface{} {
	if data == nil {
		return nil
	}
	js, err := sjson.NewJson(data)
	if logu.CheckErr(err) {
		return nil
	}
	ary, err := js.GetPath(path...).Array()
	if logu.CheckErr(err) {
		return nil
	}
	return ary
}
