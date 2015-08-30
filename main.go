package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/shaalx/membership/u"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

type Bookmark struct {
	Title    string `json:"title"`
	Official string `json:"official"`
	Bgpic    string `json:"bgpic"`
	Site     string `json:"site"`
	Remark   string `json:"remark"`
}

var cache *u.LFUCache

func readFile(name string) []byte {
	f, _ := os.OpenFile(name, os.O_RDONLY, 0064)
	defer f.Close()
	rd := bufio.NewReader(f)
	b, _ := ioutil.ReadAll(rd)
	fmt.Println(string(b))
	return b
}

func unmarshal(b []byte) []*Bookmark {
	var v []*Bookmark
	err := json.Unmarshal(b, &v)
	if nil != err {
		fmt.Println(err)
		return nil
	}
	return v
}

func main() {
	go updateBookmarks(time.Second)
	http.HandleFunc("/", bookmark)
	http.HandleFunc("/lfu", lfu)
	http.ListenAndServe(":80", nil)
}

var (
	t      *template.Template
	v      []*Bookmark
	update chan bool
)

func get(_url string) []byte {
	resp, _ := http.Get(_url)
	b, _ := ioutil.ReadAll(resp.Body)
	return b
}

func init() {
	update = make(chan bool, 10)
	t, _ = template.New("bookmark.html").ParseFiles("bookmark.html")
	// b := readFile("user.md")
	b := get("http://7xku3c.com1.z0.glb.clouddn.com/bookmark.md")
	v = unmarshal(b)
	cache = u.NewLFUCache(4)
	for i := len(v) - 1; i >= 0; i-- {
		cache.Set(v[i].Title, v[i])
	}
}

func bookmarkInCache() []*Bookmark {
	vals := cache.Vals()
	ret := make([]*Bookmark, len(vals))
	for i, it := range vals {
		bok := it.V.(*Bookmark)
		ret[i] = bok
	}
	return ret
}

func updateBookmarks(d time.Duration) {
	ticker := time.NewTicker(d)
	for {
		<-ticker.C
		v = bookmarkInCache()
		<-update
	}
}

func bookmark(rw http.ResponseWriter, req *http.Request) {
	fmt.Printf("%s  ", req.RemoteAddr)
	t.Execute(rw, v)
}

func lfu(rw http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		return
	}
	req.ParseForm()
	title := req.Form.Get("title")
	fmt.Printf("[ %s %s ]", req.RemoteAddr, title)
	cache.Get(title)
	update <- true
}
