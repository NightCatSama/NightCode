package mongo

import (
	"config"
	"fmt"
	"gopkg.in/mgo.v2"
)

var S *mgo.Session
var UserCollection *mgo.Collection
var TestCollection *mgo.Collection

// 链接数据库
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
	TestCollection = S.DB(config.Conf.Database.Db).C("test")
}

func CloseDb() {
	S.Close()
}
