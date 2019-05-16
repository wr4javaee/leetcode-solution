package no0075_sort_colors

import (
	"fmt"
	"testing"
)

func Test(test *testing.T) {
	nums := []int { 2,0,2,1,1,0 }
	sortColors(nums)
	fmt.Printf("%v", nums)
}
