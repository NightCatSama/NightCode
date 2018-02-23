package proxy

import (
	"fmt"
	"gopkg.in/mgo.v2/bson"

	"mongo"
)

type User struct {
	Account  string `json:"account"`
	Password string `json:"password"`
}

// 添加用户
func AddUser(account string, password string) (User, error) {
	user := User{}
	if mongo.UserCollection == nil {
		return user, fmt.Errorf("未连接数据库")
	}
	user, err := GetUserByAccount(account)
	if err == nil {
		return user, fmt.Errorf("用户已存在")
	}

	err = mongo.UserCollection.Insert(&User{account, password})
	if err != nil {
		return user, err
	}

	return user, err
}

// 获取用户列表
func GetUSers() ([]User, error) {
	users := []User{}
	if mongo.UserCollection == nil {
		return users, fmt.Errorf("未连接数据库")
	}
	err := mongo.UserCollection.Find(bson.M{}).All(&users)
	if err != nil {
		fmt.Println("查询用户列表失败", err)
	}
	return users, err
}

// 根据账号搜索用户
func GetUserByAccount(account string) (User, error) {
	user := User{}
	if mongo.UserCollection == nil {
		return user, fmt.Errorf("未连接数据库")
	}
	err := mongo.UserCollection.Find(bson.M{"account": account}).One(&user)
	if err != nil {
		fmt.Println("查询用户失败", err)
	}
	return user, err
}

// 删除用户
func RmoveUserByAccount(account string) error {
	err := mongo.UserCollection.Remove(bson.M{"account": account})
	return err
}
