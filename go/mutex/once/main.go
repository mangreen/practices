// https://draveness.me/golang/docs/part3-runtime/ch06-concurrency/golang-sync-primitives/
package main

import (
	"fmt"
	"sync"
)

func main() {
	o := &sync.Once{}
	for i := 0; i < 10; i++ {
		o.Do(func() {
			fmt.Println("only once")
		})
	}
}
