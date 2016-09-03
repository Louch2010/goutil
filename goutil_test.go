package goutil

import (
	"fmt"
	"testing"
)

func TestDate(t *testing.T) {
	fmt.Println(DateUtil().Date6Now())
	fmt.Println(DateUtil().Date8Now())
	fmt.Println(DateUtil().Date10Now())
	fmt.Println(DateUtil().Time14Now())
	fmt.Println(DateUtil().Time17Now())
	fmt.Println(DateUtil().TimeFullNow())
}
