package search

import (
	"encoding/json"
	"github.com/shaalx/membership/logu"
	sjson "github.com/shaalx/membership/pkg3/go-simplejson"
)

// search value
// 查询某个路径path下的key值 string
func SearchI(data []byte, key string, path ...string) interface{} {
	if data == nil {
		return ""
	}
	js, err := sjson.NewJson(data)
	if logu.CheckErr(err) {
		return ""
	}
	return js.GetPath(path...).Get(key).Interface()
}

// search value
// 查询某个路径path下的key值 string
func ISearchI(idata interface{}, key string, path ...string) interface{} {
	if idata == nil {
		return ""
	}
	data, err := json.Marshal(idata)
	if logu.CheckErr(err) {
		return ""
	}
	js, err := sjson.NewJson(data)
	if logu.CheckErr(err) {
		return ""
	}
	return js.GetPath(path...).Get(key).Interface()
}

// search value
// 查询某个路径path下的key值 string
func SearchSValue(data []byte, key string, path ...string) string {
	if data == nil {
		return ""
	}
	js, err := sjson.NewJson(data)
	if logu.CheckErr(err) {
		return ""
	}
	value, err := js.GetPath(path...).Get(key).String()
	if logu.CheckErr(err) {
		return ""
	}
	return value
}

// search value
// 查询某个路径path下的key值 string
func ISearchSValue(idata interface{}, key string, path ...string) string {
	if idata == nil {
		return ""
	}
	data, err := json.Marshal(idata)
	if logu.CheckErr(err) {
		return ""
	}
	js, err := sjson.NewJson(data)
	if logu.CheckErr(err) {
		return ""
	}
	value, err := js.GetPath(path...).Get(key).String()
	if logu.CheckErr(err) {
		return ""
	}
	return value
}

// search value
// 查询某个路径path下的key值 int
func SearchIValue(data []byte, key string, path ...string) int64 {
	if data == nil {
		return -1
	}
	js, err := sjson.NewJson(data)
	if logu.CheckErr(err) {
		return -1
	}
	value, err := js.GetPath(path...).Get(key).Int64()
	if logu.CheckErr(err) {
		return -1
	}
	return value
}

// search value
// 查询某个路径path下的key值 bool
func SearchBValue(data []byte, key string, path ...string) bool {
	if data == nil {
		return false
	}
	js, err := sjson.NewJson(data)
	if logu.CheckErr(err) {
		return false
	}
	value, err := js.GetPath(path...).Get(key).Bool()
	if logu.CheckErr(err) {
		return false
	}
	return value
}

// search value
// 查询某个路径path下的key值 float64 --> int64
func SearchFIValue(data []byte, key string, path ...string) int64 {
	if data == nil {
		return -1
	}
	js, err := sjson.NewJson(data)
	if logu.CheckErr(err) {
		return -1
	}
	value, err := js.GetPath(path...).Get(key).Float64()
	if logu.CheckErr(err) {
		return -1
	}
	return int64(value)
}
