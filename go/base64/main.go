package main

import (
	"base64/pkg"
	"fmt"
)

func main() {
	s := "ABCD"
	b64 := pkg.Encode(s)

	fmt.Println(b64)

	str, err := pkg.Decode(b64)
	if err != nil {
		fmt.Println("[ERROR]", err.Error())
		return
	}

	fmt.Println(str)
}
