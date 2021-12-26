package main

import (
	"fmt"
)

func main() {
	a := []int{2} // a[2] = 10 is largest, the smallest a[i] where 0 ≤ i < 2, the smallest is a[0]=2, a[2] − a[0] = 10 − 2 = 8
	//a := []int{7, 9, 5, 6, 3, 2} // a[1] − a[0] = 9 − 7 = 2

	fmt.Println(maxDifference(a))
}

func maxDifference(ary []int) int {
	if len(ary) < 2 {
		return -1
	}

	maxDiff := -1
	maxRight := ary[len(ary)-1]

	for i := len(ary) - 2; i >= 0; i-- {
		if ary[i] > maxRight {
			maxRight = ary[i]
		} else {
			diff := maxRight - ary[i]
			if diff > maxDiff {
				maxDiff = diff
			}
		}
	}

	return maxDiff
}
