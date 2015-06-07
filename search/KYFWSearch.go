package search

import (
	"github.com/shaalx/sstruct/utils"
)

func KYFWStatus(data []byte) bool {
	key := "status"
	path := []string{}
	return SearchBValue(data, key, path...)
}

// 更新内容 数组
func KYFWData(data []byte) []interface{} {
	return SearchArray(data, "data", []string{}...)
}

func Data(data []byte) []map[string]interface{} {
	return_map := make([]map[string]interface{}, 0)
	ary := SearchArray(data, "data", []string{}...)
	for _, val := range ary {
		ary_i := utils.I2Bytes(val)
		map_i := KYFWStem(ary_i)
		return_map = append(return_map, map_i)
	}
	return return_map
}

// 标题、摘要等题干
func KYFWStem(data []byte) map[string]interface{} {
	stem := make(map[string]interface{}, 1)
	stem["station_train_code"] = SearchSValue(data, "station_train_code", []string{"queryLeftNewDTO"}...)
	stem["from_station_name"] = SearchSValue(data, "from_station_name", []string{"queryLeftNewDTO"}...)
	stem["to_station_name"] = SearchSValue(data, "to_station_name", []string{"queryLeftNewDTO"}...)
	stem["start_time"] = SearchSValue(data, "start_time", []string{"queryLeftNewDTO"}...)
	stem["arrive_time"] = SearchSValue(data, "arrive_time", []string{"queryLeftNewDTO"}...)
	stem["lishi"] = SearchSValue(data, "lishi", []string{"queryLeftNewDTO"}...)
	stem["yz_num"] = SearchSValue(data, "yz_num", []string{"queryLeftNewDTO"}...)
	stem["yw_num"] = SearchSValue(data, "yw_num", []string{"queryLeftNewDTO"}...)
	stem["rz_num"] = SearchSValue(data, "rz_num", []string{"queryLeftNewDTO"}...)
	// has_image, ok := stem["has_image"].(bool)
	// if has_image && ok {
	// 	stem["middle_image"] = SearchSValue(data, "url", []string{"middle_image"}...)
	// }
	return stem
}
