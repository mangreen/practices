package main

import (
	"fmt"
)

func main() {
	s := make([]int, 3)
	foo(s, 1)
	fmt.Println(s)
}

func foo(s []int, i int) {
	s = append(s, i)
}
