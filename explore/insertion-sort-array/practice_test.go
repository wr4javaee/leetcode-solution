package insertion_sort_array

import (
	"fmt"
	"testing"
)

func Test(test *testing.T) {

	sorted := insertionSort([]int{ 2,5,6,7,2,5,6,7,8,1,2,0 })
	fmt.Printf("%v", sorted)
}
