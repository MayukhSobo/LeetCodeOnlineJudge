package main

import (
	"fmt"
)

// digitSums -- Returns the sum of all the digits in a number
func digitSums(num int) int {
	count := 0
	for num > 0 {
		count += num % 10
		num /= 10
	}
	return count
}

func main() {
	N := 13
	curDiv := 0
	for i := 1; i <= N/2+1; i++ {
		rem := N % i
		if rem == 0 && i > curDiv {
			if digitSums(i) > digitSums(curDiv) {
				curDiv = i
			}
		}
	}
	if digitSums(curDiv) < digitSums(N) {
		curDiv = N
	}
	fmt.Println(curDiv)
}
