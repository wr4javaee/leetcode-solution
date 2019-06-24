package no0224_basic_calculator

import (
	"fmt"
	"testing"
)

func Test(test *testing.T) {
	res := calculate("2147483647")
	fmt.Println(res)
}
