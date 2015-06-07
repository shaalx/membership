package search

func TTDisplayInfo(data []byte) string {
	key := "display_info"
	path := []string{"tips"}
	return SearchSValue(data, key, path...)
}

// 本次更新个数
func TTUpdataCount(data []byte) int64 {
	return SearchIValue(data, "total_number", []string{}...)
}

// 更新内容 数组
func TTContent(data []byte) []interface{} {
	return SearchArray(data, "data", []string{}...)
}

// 标题、摘要等题干
func TTStem(data []byte) map[string]interface{} {
	stem := make(map[string]interface{}, 15)
	stem["tag"] = SearchSValue(data, "tag", []string{}...)
	stem["title"] = SearchSValue(data, "title", []string{}...)
	stem["keywords"] = SearchSValue(data, "keywords", []string{}...)
	stem["abstract"] = SearchSValue(data, "abstract", []string{}...)
	stem["has_image"] = SearchBValue(data, "has_image", []string{}...)
	stem["article_url"] = SearchSValue(data, "article_url", []string{}...)
	stem["publish_time"] = SearchFIValue(data, "publish_time", []string{}...)
	has_image, ok := stem["has_image"].(bool)
	if has_image && ok {
		stem["middle_image"] = SearchSValue(data, "url", []string{"middle_image"}...)
	}
	return stem
}

func TTImage(data []byte) []string {
	image := make([]string, 10, 30)

	return image
}
