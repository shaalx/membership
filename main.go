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
	http.HandleFunc("/", bookmark)
	http.ListenAndServe(":80", nil)
}

var (
	t *template.Template
	v []interface{}
)

func get(_url string) []byte {
	resp, _ := http.Get(_url)
	b, _ := ioutil.ReadAll(resp.Body)
	return b
}

func init() {
	t, _ = template.New("bookmark.html").ParseFiles("bookmark.html")
	// b := readFile("user.md")
	b := get("http://7xku3c.com1.z0.glb.clouddn.com/bookmark.md")
	v = unmarshal(b)
}

func bookmark(rw http.ResponseWriter, req *http.Request) {
	fmt.Printf("%s  ", req.RemoteAddr)
	t.Execute(rw, v)
}
