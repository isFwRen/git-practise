package main

import "fmt"

func main() {

	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	sliceArr := make([]int, 0, len(slice)+1)
	sliceArr[2] = 2
	fmt.Println("slice_length11:", len(sliceArr))
	fmt.Println("slice11:", cap(sliceArr))
	for i := 0; i < len(slice); i++ {
		sliceArr = append(sliceArr, slice[i])
	}
	fmt.Println("slice_length22:", len(sliceArr))
	fmt.Println("slice22:", cap(sliceArr))
	fmt.Println(sliceArr)
}
