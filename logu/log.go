package logu

import (
	"fmt"
	"time"
)

func CheckErr(err error) bool {
	if nil != err {
		fmt.Printf("%v \t %v\n", time.Now().String(), err)
		return true
	}
	return false
}
