package goutil

import (
	"strings"
)

type StringUtil struct{}

func (this *StringUtil) TrimToEmpty(str string) string {
	if str == nil {
		return ""
	}
	return strings.TrimSpace(str)
}
