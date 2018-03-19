package test

import (
	"testing"
	. "github.com/smartystreets/goconver/convey"
	"fmt"
)

func TestTmpFunc1(t *testing.T) {
	Convey(
		"testGoConver", t, func() {
			So(1, ShouldEqual, 1)
		})
	Convey(
		"random int", t, func() {
			So(RandomInt(), ShouldEqual, 1)
		})
	Convey(
		"is same", t, func() {
			fmt.Println(Same())
			So(Same(),ShouldBeFalse )
		})


}
