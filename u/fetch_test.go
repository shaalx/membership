package u

import (
	"testing"
)

func TestFetch(t *testing.T) {
	_url := "https://api.simplr.cn/0.1/discover/filter.json?identifier=8e65b14e-338b-4191-a5c3-73e45b0b56f9&_per_page=24"
	ret := JsonFetch(_url)
	t.Log(ret)
}
