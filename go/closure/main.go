package main

import (
	"fmt"
)

func main() {
	// https://www.bangbangde.com/post/golang_goroutines_and_closures.html

	done := make(chan bool)

	s := []int{1, 2, 3}

	for _, a := range s {
		go func() {
			fmt.Println(a)
		}()
	}

	for _ = range s {
		<-done
	}
}
