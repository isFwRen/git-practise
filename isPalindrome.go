package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 121
	palindrome := isPalindrome(x)
	fmt.Println(palindrome)
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	xStr := strconv.Itoa(x)
	for i, j := 0, len(xStr)-1; i < j; i, j = i+1, j-1 {
		fmt.Println(i, j)
		if xStr[i] != xStr[j] {
			return false
		}
	}
	return true
}
