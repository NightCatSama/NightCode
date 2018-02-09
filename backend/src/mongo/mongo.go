package mongo

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"config"
)

var S *mgo.Session
var UserCollection *mgo.Collection

func LinkDb() {
	uri := config.Conf.Database.Host + ":" + config.Conf.Database.Port
	S, err := mgo.Dial(uri)
	if err != nil {
		panic(err)
	}
	// Optional. Switch the session to a monotonic behavior.
	S.SetMode(mgo.Monotonic, true)

	fmt.Printf("[MongoDB is connected]: %s/%s \n", uri, config.Conf.Database.Db)

	UserCollection = S.DB(config.Conf.Database.Db).C("user")
}

func CloseDb() {
	S.Close()
}
