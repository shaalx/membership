package u

import (
	"testing"
)

func TestFetch(t *testing.T) {
	_url := "https://www.baidu.com"
	ret := JsonFetch(_url)
	t.Log(ret)
}
