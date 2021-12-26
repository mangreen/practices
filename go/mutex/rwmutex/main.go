//https://juejin.im/post/6844903925364031496
package main

import (
	"sync"
	"time"
)

var mu sync.RWMutex
var data map[string]string

func main() {
	// data = map[string]string{"hoge": "fuga"}
	// mu = sync.RWMutex{}
	// go read()
	// go write()
	// go read()
	// go read()
	// time.Sleep(5 * time.Second)

	rUnlockMoreThanRLock()
}

// 读方法
func read() {
	println("read_start")
	mu.RLock()
	defer mu.RUnlock()
	time.Sleep(1 * time.Second)
	println("read_complete", data["hoge"])
}

// 写方法
func write() {
	println("write_start")
	//仔细看下这两行代码,此处是注释掉的
	mu.Lock()
	defer mu.Unlock()
	time.Sleep(2 * time.Second)
	data["hoge"] = "piyo"
	println("write_complete")
}

// https://blog.csdn.net/chenbaoke/article/details/41957725
// 當RUnlock多於RLock多個時，便會報錯，進入死鎖
func rUnlockMoreThanRLock() {
	l := new(sync.RWMutex)
	l.RUnlock()
	l.RUnlock() // 此處出現死鎖
	println("1")
	l.RLock()
}
