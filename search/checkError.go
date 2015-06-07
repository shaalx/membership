package search

import (
	"github.com/shaalx/sstruct/service/log"
)

func checkError(err error) bool {
	if log.IsError("{go-simplejson analyse json}", err) {
		return true
	}
	return false
}
