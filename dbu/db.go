package dbu

import (
	"fmt"
	"github.com/shaalx/merbership/logu"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"os"
	"sync"
)

type MgoDB struct {
	sync.RWMutex
	Session    *mgo.Session
	Collection map[string]*Collection
}

type Collection struct {
	c *mgo.Collection
}

func Conn() string {
	conn := ""
	if len(os.Getenv("MONGODB_USERNAME")) > 0 {
		conn += os.Getenv("MONGODB_USERNAME")

		if len(os.Getenv("MONGODB_PASSWORD")) > 0 {
			conn += ":" + os.Getenv("MONGODB_PASSWORD")
		}

		conn += "@"
	}

	if len(os.Getenv("MONGODB_PORT_27017_TCP_ADDR")) > 0 {
		conn += os.Getenv("MONGODB_PORT_27017_TCP_ADDR")
	} else {
		conn += "localhost"
	}

	if len(os.Getenv("MONGODB_PORT_27017_TCP_PORT")) > 0 {
		conn += ":" + os.Getenv("MONGODB_PORT_27017_TCP_PORT")
	} else {
		conn += ":27017"
	}
	// defaultly using "test" as the db instance
	db := "nation"

	if len(os.Getenv("MONGODB_INSTANCE_NAME")) > 0 {
		db = os.Getenv("MONGODB_INSTANCE_NAME")
	}

	conn += "/" + db
	fmt.Println(conn)
	return conn
}

func NewMgoDB(dailStr string) *MgoDB {
	if len(dailStr) <= 0 {
		dailStr = "127.0.0.1:27017"
	}
	session, err := mgo.Dial(dailStr)
	if logu.CheckErr(err) {
		panic(err)
	}
	return &MgoDB{
		Session:    session,
		Collection: make(map[string]*Collection, 0),
	}
}

func (db *MgoDB) checkSession() bool {
	err := db.Session.Ping()
	return !logu.CheckErr(err)
}

func (db *MgoDB) SetCollection(params ...string) *Collection {
	if !db.checkSession() {
		return nil
	}
	if len(params) < 2 {
		return nil
	}
	c := db.Session.DB(params[0]).C(params[1])
	db.Lock()
	defer db.Unlock()
	myC := &Collection{
		c: c,
	}
	db.Collection[params[1]] = myC
	return myC
}

func (db *MgoDB) Close() {
	db.Session.Close()
}

func (db *MgoDB) GetCollection(params ...string) *Collection {
	if len(params) < 2 {
		return nil
	}
	if c, ok := db.Collection[params[1]]; ok {
		return c
	}
	return db.SetCollection(params...)
}

func (c *Collection) Count(selector bson.M) int {
	if c == nil || c.c == nil {
		return -1
	}
	n, err := c.c.Find(selector).Count()
	if logu.CheckErr(err) {
		return -1
	}
	return n
}

func (c *Collection) Select(selector bson.M) *bson.M {
	if c == nil || c.c == nil {
		return nil
	}
	var result bson.M
	err := c.c.Find(selector).One(&result)
	if logu.CheckErr(err) {
		return nil
	}
	return &result
}

func (c *Collection) Insert(structs ...interface{}) int {
	if c == nil || c.c == nil {
		return -1
	}
	if len(structs) == 0 {
		return 0
	}
	err := c.c.Insert(structs...)
	if logu.CheckErr(err) {
		return -1
	}
	return len(structs)
}
