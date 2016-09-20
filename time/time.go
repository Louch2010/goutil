package time

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

type DateUtil struct{}

/*获取当前时间戳*/
func (this DateUtil) TimestampNow() int64 {
	return time.Now().Unix()
}

/*获取当前日期*/
func (this DateUtil) Date6Now() string {
	return time.Now().Format(DateTime_Pattern_yyMMdd)
}
func (this DateUtil) Date8Now() string {
	return time.Now().Format(DateTime_Pattern_yyyyMMdd)
}
func (this DateUtil) Date10Now() string {
	return time.Now().Format(DateTime_Pattern_yyyy_MM_dd)
}

/*获取当前时间*/
func (this DateUtil) Time14Now() string {
	return time.Now().Format(DateTime_Pattern_yyyyMMddHHmmss)
}
func (this DateUtil) Time17Now() string {
	return time.Now().Format(DateTime_Pattern_yyyyMMddHHmmssSSS)
}
func (this DateUtil) TimeFullNow() string {
	return time.Now().Format(DateTime_Pattern_yyyy_MM_dd_HH_mm_ss)
}

/*日期时间格式转换*/
func (this DateUtil) DateTimeFormat(t time.Time, pattern string) string {
	return t.Format(pattern)
}
func (this DateUtil) Date6Format(t time.Time) string {
	return t.Format(DateTime_Pattern_yyMMdd)
}
func (this DateUtil) Date8Format(t time.Time) string {
	return t.Format(DateTime_Pattern_yyyyMMdd)
}
func (this DateUtil) Date10Format(t time.Time) string {
	return t.Format(DateTime_Pattern_yyyy_MM_dd)
}
func (this DateUtil) Time14Format(t time.Time) string {
	return t.Format(DateTime_Pattern_yyyyMMddHHmmss)
}
func (this DateUtil) Time17Format(t time.Time) string {
	return t.Format(DateTime_Pattern_yyyyMMddHHmmssSSS)
}
func (this DateUtil) TimeFullFormat(t time.Time) string {
	return t.Format(DateTime_Pattern_yyyy_MM_dd_HH_mm_ss)
}

/*日期时间解析*/
func (this DateUtil) ParseTime14(str string) (time.Time, error) {
	location, _ := time.LoadLocation("Local")
	t, err := time.ParseInLocation(DateTime_Pattern_yyyyMMddHHmmss, str, location)
	return t, err
}
