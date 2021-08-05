package util

import (
	"github.com/Lofanmi/chinese-calendar-golang/calendar"

	"fmt"
	"time"
)

func Lunar() {
	t := time.Now()
	// 1. ByTimestamp
	// 时间戳
	c := calendar.ByTimestamp(t.Unix())
	// 2. BySolar
	// 公历
	c2 := calendar.BySolar(1991, 10, 16, 20, 11, 9)
	// 3. ByLunar
	// 农历(最后一个参数表示是否闰月)
	c3 := calendar.ByLunar(1991, 9, 9, 20, 11, 9, false)
	s1, _ := c.ToJSON()
	s2, _ := c2.ToJSON()
	s3, _ := c3.ToJSON()
	fmt.Println(string(s1))
	fmt.Println(string(s2))
	fmt.Println(string(s3))
}
