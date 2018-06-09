package main

import (
	"fmt"
	"log"
	"time"
)

// 10 OMIT
func worker(i int, ch chan<- int) {
	go func(inp int) {
		if inp < 0 { // HL
			log.Println("can't handle negative numbers!")
			return
		}
		time.Sleep(1 * time.Second)
		ch <- inp * 2
	}(i)
}
func main() {
	dat := []int{1, 2, -2, 3} // HL
	rCh := make(chan int, len(dat))
	for _, v := range dat {
		worker(v, rCh)
	}
	for i := 0; i < len(dat); i++ {
		fmt.Println(<-rCh)
	}
}

// 20 OMIT
