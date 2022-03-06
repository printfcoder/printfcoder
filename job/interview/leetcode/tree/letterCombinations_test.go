package tree

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestLetterCombinations(t *testing.T) {
	Convey("testing letterCombinations", t, func() {
		v := letterCombinations("23")
		So(v, ShouldResemble, []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"})
		v = letterCombinations("2")
		So(v, ShouldResemble, []string{"a", "b", "c"})
		v = letterCombinations("22")
		So(v, ShouldResemble, []string{"aa", "ab", "ac", "ba", "bb", "bc", "ca", "cb", "cc"})
		v = letterCombinations("234")
		So(v, ShouldResemble, []string{"adg", "adh", "adi", "aeg", "aeh", "aei", "afg", "afh", "afi", "bdg", "bdh", "bdi", "beg", "beh", "bei", "bfg", "bfh", "bfi", "cdg", "cdh", "cdi", "ceg", "ceh", "cei", "cfg", "cfh", "cfi"})
		v = letterCombinations("2345")
		So(v, ShouldResemble, []string{"adgj", "adgk", "adgl", "adhj", "adhk", "adhl", "adij", "adik", "adil", "aegj", "aegk", "aegl", "aehj", "aehk", "aehl", "aeij", "aeik", "aeil", "afgj", "afgk", "afgl", "afhj", "afhk", "afhl", "afij", "afik", "afil", "bdgj", "bdgk", "bdgl", "bdhj", "bdhk", "bdhl", "bdij", "bdik", "bdil", "begj", "begk", "begl", "behj", "behk", "behl", "beij", "beik", "beil", "bfgj", "bfgk", "bfgl", "bfhj", "bfhk", "bfhl", "bfij", "bfik", "bfil", "cdgj", "cdgk", "cdgl", "cdhj", "cdhk", "cdhl", "cdij", "cdik", "cdil", "cegj", "cegk", "cegl", "cehj", "cehk", "cehl", "ceij", "ceik", "ceil", "cfgj", "cfgk", "cfgl", "cfhj", "cfhk", "cfhl", "cfij", "cfik", "cfil"})
		v = letterCombinations("")
		So(v, ShouldResemble, []string{})
	})
}
