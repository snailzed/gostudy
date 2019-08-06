package base

import (
	"fmt"
	"time"
)

func main() {
	//now
	now := time.Now()
	fmt.Println(now)
	format(now)
	//当前时间戳
	fmt.Println(now.Unix())

	//时间戳转 time类型
	var timeSeon int64 = 1564400072
	timeObj := time.Unix(timeSeon, 0)
	format(timeObj)

	//一些常量 最小单位为纳秒，如果不写明单位，则默认为纳秒
	fmt.Printf("nano second:%d\n", time.Nanosecond)
	fmt.Printf("micro second:%d\n", time.Microsecond)
	fmt.Printf("milli second:%d\n", time.Millisecond)
	fmt.Printf("second:%d\n", time.Second)

	//格式化输出 2006-01-02 15:04:05 go诞生时间
	fmt.Println(now.Format("2006-01-02 15:04:05"))

	//ticker定时器
	c := time.Tick(time.Second)
	for i := range c {
		fmt.Println(i)
	}

}

func format(now time.Time) {
	//当前时间
	year := now.Year()
	month := now.Month()
	day := now.Day()
	hour := now.Hour()
	min := now.Minute()
	second := now.Second()
	fmt.Printf("%4d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, min, second)
}
