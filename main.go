package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
)

func readFile(name string) []byte {
	f, _ := os.OpenFile(name, os.O_RDONLY, 0064)
	defer f.Close()
	rd := bufio.NewReader(f)
	b, _ := ioutil.ReadAll(rd)
	return b
}

func unmarshal(b []byte) []interface{} {
	var v []interface{}
	err := json.Unmarshal(b, &v)
	if nil != err {
		fmt.Println(err)
		return nil
	}
	return v
}

func main() {
	http.HandleFunc("/", info)
	http.ListenAndServe(":80", nil)
}

var (
	t *template.Template
	v []interface{}
)

func init() {
	t, _ = template.New("info.html").ParseFiles("info.html")
	b := readFile("user.md")
	v = unmarshal(b)
}

func info(rw http.ResponseWriter, req *http.Request) {
	fmt.Printf("%s  ", req.RemoteAddr)
	t.Execute(rw, v)
}
