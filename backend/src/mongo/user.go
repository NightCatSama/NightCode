package mongo

import (
	"fmt"
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	Account  string
	Password string
}

// 添加用户
func AddUser(account string, password string) (User, error) {
	user := User{}
	if C == nil {
		return user, fmt.Errorf("找不到数据库")
	}
	err := C.Insert(&User{ account, password })
	user, err = getUserByAccount(account)
	if err != nil {
		return user, err
	}

	return user, err
}

// 得到用户列表
func GetUSers() ([]User, error) {
	users := []User{}
	if C == nil {
		return users, fmt.Errorf("找不到数据库")
	}
	err := C.Find(bson.M{}).All(&users)
	if err != nil {
		fmt.Println("查询用户列表失败", err)
	}
	return users, err
}

// 根据账号搜索用户
func getUserByAccount(account string) (User, error) {
	user := User{}
	if C == nil {
		return user, fmt.Errorf("找不到数据库")
	}
	err := C.Find(bson.M{ "account": account }).One(&user)
	if err != nil {
		fmt.Println("查询用户失败", err)
	}
	return user, err
}