package base

import (
	"fmt"
	"testing"
)

func TestNewCode(t *testing.T) {
	err := NewCodeError(500, "error")
	fmt.Printf("err:%v", err.Error())

}
