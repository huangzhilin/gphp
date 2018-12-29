package gphp

import (
	"time"
)

const default_format = "2006-01-02 15:04:05"

//获取 当前时间戳
func Time() int {
	return int(time.Now().Unix())
}

/*
时间戳 转 字符串
format	    格式    2006-01-02 15:04:05
timestamp	时间戳
*/
func Date(format string, timestamp int) string {
	return time.Unix(int64(timestamp), 0).Format(format)
}

/*
字符串 转 时间戳
format	    格式    2006-01-02 15:04:05
*/
func StrToTime(format string) int {
	loc, _ := time.LoadLocation("Local") //重要：获取时区
	theTime, _ := time.ParseInLocation(default_format, format, loc)
	return int(theTime.Unix())
}

//睡眠时间：睡Second秒
func Sleep(Second int) {
	time.Sleep(time.Duration(Second) * time.Second)
}
