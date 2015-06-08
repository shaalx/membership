package dbu

import (
	"testing"
)

func TestConn(t *testing.T) {
	conn := Conn()
	t.Log(conn)
}
