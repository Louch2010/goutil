package strings

import (
	"strconv"
	"strings"
)

type StringUtil struct{}

//去除字符串空格，当为nil时，返回空串
func (this StringUtil) TrimToEmpty(str string) string {
	return strings.TrimSpace(str)
}

//判断是否为空，当为nil或空串时返回false， 否则返回true
func (this StringUtil) IsEmpty(str string) bool {
	return len(this.TrimToEmpty(str)) == 0
}

//将数字转换成指定长度的字符串，如果小于指定长度，则以0填充，大于指定长度，则截取
func (this StringUtil) IntToStr(num int, length int) string {
	str := strconv.Itoa(num)
	for len(str) < length {
		str = "0" + str
	}
	if len(str) > length {
		str = str[:length]
	}
	return str
}

//将以0填充的字符串转换成数字
func (this StringUtil) StrToInt(str string) (int, error) {
	num, err := strconv.Atoi(strings.TrimLeft(str, "0"))
	return num, err
}
