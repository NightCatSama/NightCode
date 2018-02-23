package mongo

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestLinkDb(t *testing.T) {
	LinkDb()

	Convey("连接数据库", t, func() {
		So(TestCollection, ShouldNotBeNil)
	})
}
