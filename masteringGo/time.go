package main

import (
	"fmt"
	"time"
)

func main() {
	// 这里是关于time包的常见使用场景
	// 获取当前时间
	fmt.Println(
		time.Now(),
		time.Now().Unix(), // 获取当前秒级时间戳
		time.Now().UnixNano() / 1000000, // 获取当前毫秒级时间戳
		time.Now().UnixNano(), // 获取当前纳秒级时间戳
	)
	// 睡眠一定时间 2s，让出资源 等待2s后调度
	time.Sleep(time.Second * 2)
	fmt.Println(time.Now(), time.Now().Unix(), time.Now().UnixNano() / 1000000, time.Now().UnixNano())

	// 时间单位，go通过Duration都表示时间单位 如下分别是 时、分、秒、毫秒、微妙、纳秒
	fmt.Println(time.Hour, time.Minute, time.Second, time.Millisecond, time.Microsecond, time.Nanosecond)

	now := time.Now()
	// 时间格式化
	// go的时间格式化非常有意思。。。不是使用常见的YYY表示年 mm表示月之类的 而是用一个字符串例子表示格式化的格式
	// /usr/local/go/src/time/format.go:73 在这可以找到支持的所有类型
	// 不过并没有我们常用的 只好自己拼一个了
	timeStr := now.Format("2006-01-02 15:04:05")
	fmt.Println(timeStr)

	// 从字符串解析时间，注意time.Parse会按照UTC去转换，实际我们需要的是本地时间
	t1, err := time.ParseInLocation("2006-01-02 15:04:05", timeStr, time.Local)
	if err != nil {
		panic(err)
	}
	// 对比转换前后的秒级时间戳
	fmt.Println(t1.Unix() == now.Unix())

	// 获取时间的年月日时分秒等信息
	fmt.Println(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())

	// 操作时间 设置 一个月前 一天后 一小时后等等，并比较
	beforeMonth := now.AddDate(0, -1, 0)
	nextDay := now.AddDate(0, 0, 1)
	nextHour := now.Add(time.Hour * 1)
	nextMinute := now.Add(time.Minute * 1)
	fmt.Println(beforeMonth.Format("2006-01-02 15:04:05"),
		nextDay.Format("2006-01-02 15:04:05"),
		nextHour.Format("2006-01-02 15:04:05"),
		nextMinute.Format("2006-01-02 15:04:05"))
	fmt.Println(nextDay.After(now))
}
