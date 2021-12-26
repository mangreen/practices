// https://www.geeksforgeeks.org/pangram-checking/
package main

import (
	"fmt"
	"unicode"
)

func main() {
	s := []string{
		"we promptly judged antique ivory buckles for the next prize",
		"we promptly judged antique ivory buckles for the prizes",
		"the quick brown fox jumps over the lazy dog",
		"the quick brown fox jump over the lazy dog",
	}

	fmt.Println(isPangram(s))
}

func isPangram(ary []string) string {
	ans := ""

	for i := 0; i < len(ary); i++ {
		if helper(ary[i]) {
			ans = ans + "1"
		} else {
			ans = ans + "0"
		}
	}

	return ans
}

func helper(str string) bool {
	m := [26]int{}

	for _, r := range []rune(str) {
		if unicode.IsLetter(r) {
			t := unicode.ToLower(r)

			t = t - 'a'
			m[t]++
		}
	}

	for _, u := range m {
		if u == 0 {
			return false
		}
	}

	return true
}
