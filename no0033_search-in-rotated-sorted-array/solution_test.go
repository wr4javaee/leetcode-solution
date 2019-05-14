package no0033

import (
	"fmt"
	"testing"
)

func TestSearch(test *testing.T) {

	searchRes := search([]int{5,1,3}, 5)
	fmt.Println(searchRes)
}
