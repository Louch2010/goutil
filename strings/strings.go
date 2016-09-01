package strings

import (
	"strings"
)

func TrimToEmpty(str string) string {
	return strings.TrimSpace(str)
}

func GetHello() string {
	return "hello~"
}
