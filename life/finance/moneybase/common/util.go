package common

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"time"

	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

const Today = "today"

// GbkToUtf8 GBK 转 UTF-8
func GbkToUtf8(s []byte) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewDecoder())
	d, e := ioutil.ReadAll(reader)
	if e != nil {
		return nil, e
	}
	return d, nil
}

func ParseDayStartAndEnd(date string) (start int64, end int64) {
	if date == Today {
		date = TodayStr()
	}

	// 20230101235929 -> 20230101000001 20230101235959，转成一天的开始和结束
	shijianPrefix := date[0:8]
	start, _ = strconv.ParseInt(shijianPrefix+"000001", 10, 64)
	end, _ = strconv.ParseInt(shijianPrefix+"235959", 10, 64)

	return
}

// AddDate 增加日期，date格式：20060102
func AddDate(date string, dayNum int) (start, end int64, err error) {
	da, err := time.Parse("20060102", date)
	if err != nil {
		err = fmt.Errorf("[AddDate] invalid date format: %s", err)
		return
	}

	da = da.AddDate(0, 0, dayNum)

	// format s := t.Format("2006-01-02 15:04:05.999999999 -0700 MST")
	dateNewStr := strings.Replace(strings.Split(da.String(), " ")[0], "-", "", -1)
	dateNewStr += "010101"

	start, end = ParseDayStartAndEnd(dateNewStr)
	return
}

func TodayStr() string {
	now := time.Now()
	year := now.Year()
	mon := now.Month()
	day := now.Day()

	monS := fmt.Sprintf("%d", mon)
	if mon < 10 {
		monS = "0" + monS
	}

	dayS := fmt.Sprintf("%d", day)
	if day < 10 {
		dayS = "0" + dayS
	}
	ret := fmt.Sprintf("%d%s%s", year, monS, dayS)

	return ret
}
