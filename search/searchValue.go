package search

import (
	sjson "github.com/shaalx/sstruct/pkg3/go-simplejson"
)

// search value
// 查询某个路径path下的key值 string
func SearchSValue(data []byte, key string, path ...string) string {
	if data == nil {
		return ""
	}
	js, err := sjson.NewJson(data)
	if checkError(err) {
		return ""
	}
	value, err := js.GetPath(path...).Get(key).String()
	if checkError(err) {
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
	if checkError(err) {
		return -1
	}
	value, err := js.GetPath(path...).Get(key).Int64()
	if checkError(err) {
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
	if checkError(err) {
		return false
	}
	value, err := js.GetPath(path...).Get(key).Bool()
	if checkError(err) {
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
	if checkError(err) {
		return -1
	}
	value, err := js.GetPath(path...).Get(key).Float64()
	if checkError(err) {
		return -1
	}
	return int64(value)
}
