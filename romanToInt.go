package main

import "fmt"

func main() {
	s := "III"
	toInt := romanToInt(s)
	fmt.Println(toInt)
}

func romanToInt(s string) int {
	translist := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	result := translist[s[len(s)-1]]
	fmt.Println(s[len(s)-1])
	fmt.Println(result)
	for i := len(s) - 1; i > 0; i-- {
		if translist[s[i]] <= translist[s[i-1]] {
			result += translist[s[i-1]]
		} else {
			result -= translist[s[i-1]]
		}
	}
	return result
}
