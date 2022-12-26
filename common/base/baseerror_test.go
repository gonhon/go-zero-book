package base

import (
	"fmt"
	"testing"
)

func TestNewCode(t *testing.T) {
	err := NewCodeError(500, "error")
	fmt.Printf("err:%v", err.Error())

	str := fmt.Sprintf("`%%%s%%`", "哈哈")
	fmt.Printf("%v", str)
}
