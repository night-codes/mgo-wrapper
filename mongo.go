// Package mongo is wrap for package mgo
package mongo

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"os"
	"time"
)

// Obj is a convenient alias for a map[string]interface{} map, useful for
// dealing with json objects in a native way.  For instance:
//
//     o := mongo.Obj{"a": 1, "b": true}
type Obj map[string]interface{}

// Arr is a convenient alias for a []interface{} slice, useful for
// dealing with json arrays in a native way.  For instance:
//
//     a := mongo.Arr{1, 2, 3, 4}
//     fmt.Println(a)
//     // out: [1,2,3,4]
type Arr []interface{}

var (
	session *mgo.Session
	err     error
)

func init() {
	dbhost := os.Getenv("DBHOST")
	if len(dbhost) == 0 {
		dbhost = "localhost"
	}

	session, err = mgo.Dial(dbhost)
	if err != nil {
		panic(err)
	}

	session.SetMode(mgo.Monotonic, true)

	go func() {
		for range time.Tick(time.Second * 1) {
			ConnectionCheck()
		}
	}()
}

// DB wrap to mgo.Session.DB
func DB(dname string) *mgo.Database {
	ConnectionCheck()
	return session.DB(dname)
}

// ConnectionCheck implements db reconnection
func ConnectionCheck() {
	if err := session.Ping(); err != nil {
		fmt.Println("Lost connection to db!")
		session.Refresh()
		if err := session.Ping(); err == nil {
			fmt.Println("Reconnect to db successful.")
		}
	}
}
