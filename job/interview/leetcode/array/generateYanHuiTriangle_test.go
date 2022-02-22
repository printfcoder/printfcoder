package array

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestGenerateYanHuiTriangle(t *testing.T) {
	Convey("Testing Generate YanHui Triangle", t, func() {
		output := generateYanHuiTriangle(1)
		So(output, ShouldResemble, [][]int{{1}})
		output = generateYanHuiTriangle(2)
		So(output, ShouldResemble, [][]int{{1}, {1, 1}})
		output = generateYanHuiTriangle(5)
		So(output, ShouldResemble, [][]int{{1}, {1, 1}, {1, 2, 1}, {1, 3, 3, 1}, {1, 4, 6, 4, 1}})
	})
}
