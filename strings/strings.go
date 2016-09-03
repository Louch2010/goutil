package strings

import (
	"strings"
)

type StringUtil struct{}

//去除字符串空格，当为nil时，返回空串
func (this StringUtil) TrimToEmpty(str string) string {
	return strings.TrimSpace(str)
}

//判断是否为空，当为nil或空串时返回false，否则返回true
func (this StringUtil) IsEmpty(str string) bool {
	return len(this.TrimToEmpty(str)) == 0
}
