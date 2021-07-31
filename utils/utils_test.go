package utils

import (
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

func TestConvertInfToList(t *testing.T) {
	convey.Convey("Convert To Interface To list of struct", t, func() {
		type ABC struct {
			A string `json:"a"`
			B string `json:"b"`
		}
		listObjJSON := `[
			{
				"a": "123",
				"b": "456"
			},
			{
				"a": "abc",
				"b": "def"
			}
		]`
		listObj := make([]ABC, 0)
		err := ParseInterfaceToList(listObjJSON, &listObj)
		convey.So(err, convey.ShouldEqual, nil)
		convey.So(len(listObj), convey.ShouldEqual, 2)
		convey.So(listObj[0].A, convey.ShouldEqual, "123")
		convey.So(listObj[0].B, convey.ShouldEqual, "456")
		convey.So(listObj[1].A, convey.ShouldEqual, "abc")
		convey.So(listObj[1].B, convey.ShouldEqual, "def")
	})
}
