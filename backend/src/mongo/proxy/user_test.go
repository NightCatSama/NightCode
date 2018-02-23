package proxy

import (
	"math/rand"
	"mongo"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"

// 获取随机字符串
func RandStringBytes(n int) string {
	b := make([]byte, n)
	rand.Seed(42)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

// 测试用户表操作
func TestUserActions(t *testing.T) {
	mongo.LinkDb()

	Convey("链接数据库用户表", t, func() {
		So(mongo.UserCollection, ShouldNotBeNil)

		Convey("获取用户列表", func() {
			_, err := GetUSers()
			So(err, ShouldBeNil)
		})

		account := RandStringBytes(8)
		password := RandStringBytes(8)
		Convey("添加用户", func() {
			_, err := AddUser(account, password)
			So(err, ShouldBeNil)
		})

		Convey("通过用户名查询用户", func() {
			_, err := GetUserByAccount(account)
			So(err, ShouldBeNil)
		})

		Convey("用过用户名删除用户", func() {
			err := RmoveUserByAccount(account)
			So(err, ShouldBeNil)
		})
	})
}
