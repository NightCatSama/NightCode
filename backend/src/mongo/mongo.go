package mongo

import (
	"fmt"
	"gopkg.in/mgo.v2"
)

var S *mgo.Session
var C *mgo.Collection

func LinkDb() {
	S, err := mgo.Dial("localhost:27017")
	if err != nil {
		panic(err)
	}
	// Optional. Switch the session to a monotonic behavior.
	S.SetMode(mgo.Monotonic, true)

	C = S.DB("nightcode").C("user")
	fmt.Println("[MongoDB is connected]: localhost:27017/nightcode")
}

func CloseDb() {
	S.Close()
}
