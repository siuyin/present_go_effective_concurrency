package main

import (
	"fmt"
	"time"
)

// 10 OMIT
func worker(i int, ch chan<- int) {
	go func(inp int) { // HL
		time.Sleep(1 * time.Second)
		ch <- inp * 2
	}(i)
}
func main() {
	dat := []int{1, 2, 3}
	rCh := make(chan int, len(dat))
	for _, v := range dat {
		worker(v, rCh)
	}
	for i := 0; i < len(dat); i++ {
		fmt.Println(<-rCh)
	}
}

// 20 OMIT
