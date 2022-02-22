package math

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestPartition(t *testing.T) {
	Convey("testing partition", t, func() {
		v := partition("aab")
		So(v, ShouldResemble, [][]string{{"a", "a", "b"}, {"aa", "b"}})
	})
}
