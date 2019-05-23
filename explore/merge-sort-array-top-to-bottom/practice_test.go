package merge_sort_array_top_to_bottom

import (
	"fmt"
	"testing"
)

func Test(test *testing.T) {
	sorted := mergeSort([]int{ 2,5,6,7,2,5,6,7,8,1,2,0 })
	fmt.Printf("%v", sorted)
}

