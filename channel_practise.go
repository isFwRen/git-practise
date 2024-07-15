package main

import (
	"fmt"
	"time"
)

func main() {
	traverse_channel()
}

func traverse_channel() {
	ch := make(chan int, 10)
	start := time.Now()
	go func() {
		ch <- 1
		ch <- 2
		ch <- 3
		close(ch)
	}()
	//ch <- 1
	//ch <- 2
	//ch <- 3
	//close(ch)
	for resp := range ch {
		fmt.Println(resp)
	}
	time.Sleep(time.Second * 10)
	fmt.Printf("%vms elapsed\n", time.Since(start).Milliseconds())
}
