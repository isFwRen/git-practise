package main

import "fmt"

func main() {
	strs := []string{"flower", "flower", "flower", "flower", "flower", "flowe"}
	prefix := longestCommonPrefix(strs)
	fmt.Println(prefix)
}

func longestCommonPrefix(strs []string) string {
	for i := range strs[0] {
		for _, v := range strs {
			if i == len(v) || v[i] != strs[0][i] {
				return strs[0][:i]
			}
		}
	}
	return strs[0]
}
