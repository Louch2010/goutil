package strings

import (
	"strings"
)

type StringUtil struct{}

func (this StringUtil) TrimToEmpty(str string) string {
	return strings.TrimSpace(str)
}

func (this StringUtil) GetHello() string {
	return "hello~"
}
