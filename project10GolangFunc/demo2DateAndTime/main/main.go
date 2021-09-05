package main

import (
	"fmt"
	"time"
)

func main() {
	//当前时间
	now := time.Now()
	fmt.Printf("当前时间为%v,now的类型:%T\n", now, now)
	fmt.Printf("年:%v\n", now.Year())
	fmt.Printf("月:%v\n", now.Month())
	fmt.Printf("月:%v\n", int(now.Month()))
	fmt.Printf("日:%v\n", now.Day())
	fmt.Printf("时:%v\n", now.Hour())
	fmt.Printf("分:%v\n", now.Minute())
	fmt.Printf("秒:%v\n", now.Second())
	fmt.Printf("纳秒:%v\n", now.Nanosecond())

	//格式化日期和时间
	fmt.Printf("当前时间年月日:%d-%d-%d %d:%d:%d\n", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
	dateStr := fmt.Sprintf("当前时间年月日:%d-%d-%d %d:%d:%d\n", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
	fmt.Println(dateStr)
	//使用Time.Format，golang作者比较任性，使用2006-01-02 15:04:05这个时间作为yyyy-MM-dd HH:mm:ss使用
	fmt.Println(now.Format("2006-01-02 15:04:05"))

	//时间常量
	for i := 0; i < 5; i++ {
		time.Sleep(time.Millisecond) //休眠多少时间，使用常量来处理
		fmt.Println(i)
	}

	//获取当前unix时间戳、unixnano时间戳
	fmt.Printf("unix时间戳:%v\n", now.Unix())
	fmt.Printf("unix毫秒时间戳:%v\n", now.UnixMilli())
	fmt.Printf("unix微秒时间戳:%v\n", now.UnixMicro())
	fmt.Printf("unix纳秒时间戳:%v\n", now.UnixNano())
}
