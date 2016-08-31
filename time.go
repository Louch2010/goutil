package goutil

import (
	"time"
)

const (
	DateTime_Pattern_yyMMdd     = "060102"
	DateTime_Pattern_yyyyMMdd   = "20060102"
	DateTime_Pattern_yyyy_MM_dd = "2006-01-02"

	DateTime_Pattern_yyyyMMddHHmmss          = "20060102150405"
	DateTime_Pattern_yyyy_MM_dd_HH_mm_ss     = "2006-01-02 15:04:05"
	DateTime_Pattern_yyyyMMddHHmmssSSS       = "20060102150405Z07"
	DateTime_Pattern_yyyy_MM_dd_HH_mm_ss_SSS = "2006-01-02 15:04:05.Z07"
)

/*获取当前时间戳*/
func TimestampNow() int64 {
	return time.Now().Unix()
}

/*获取当前日期*/
func Date6Now() string {
	return time.Now().Format(DateTime_Pattern_yyMMdd)
}
func Date8Now() string {
	return time.Now().Format(DateTime_Pattern_yyyyMMdd)
}
func Date10Now() string {
	return time.Now().Format(DateTime_Pattern_yyyy_MM_dd)
}

/*获取当前时间*/
func Time14Now() string {
	return time.Now().Format(DateTime_Pattern_yyyyMMddHHmmss)
}
func Time17Now() string {
	return time.Now().Format(DateTime_Pattern_yyyyMMddHHmmssSSS)
}
func TimeFullNow() string {
	return time.Now().Format(DateTime_Pattern_yyyy_MM_dd_HH_mm_ss)
}

/*日期时间格式转换*/
func DateTimeFormat(t time.Time, pattern string) string {
	return t.Format(pattern)
}
func Date6Format(t time.Time) string {
	return t.Format(DateTime_Pattern_yyMMdd)
}
func Date8Format(t time.Time) string {
	return t.Format(DateTime_Pattern_yyyyMMdd)
}
func Date10Format(t time.Time) string {
	return t.Format(DateTime_Pattern_yyyy_MM_dd)
}
func Time14Format(t time.Time) string {
	return t.Format(DateTime_Pattern_yyyyMMddHHmmss)
}
func Time17Format(t time.Time) string {
	return t.Format(DateTime_Pattern_yyyyMMddHHmmssSSS)
}
func TimeFullFormat(t time.Time) string {
	return t.Format(DateTime_Pattern_yyyy_MM_dd_HH_mm_ss)
}

/*func main() {
	fmt.Println(Date6Now())
	fmt.Println(Date8Now())
	fmt.Println(Date10Now())
	fmt.Println(Time14Now())
	fmt.Println(Time17Now())
	fmt.Println(TimeFullNow())
}*/
