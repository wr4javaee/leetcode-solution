package main

import "fmt"

func main() {
	s := "01234"
	for i := 0; i < len(s); i++ {
		fmt.Println(s[i])
	}

	for inx, val := range s {
		fmt.Println(inx)
		fmt.Println(val)
	}
}
