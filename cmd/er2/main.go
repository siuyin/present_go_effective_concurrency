package main

import (
	"fmt"
	"log"
	"time"
)

// 10 OMIT
func worker(i int, ch chan<- int) {
	go func(inp int) {
		if inp < 0 {
			log.Println("can't handle negative numbers!")
			// 2. FIXME: time.Sleep(2 * time.Second)
			close(ch) // HL
			return
		}
		time.Sleep(1 * time.Second)
		ch <- inp * 2
	}(i)
}

// 15 OMIT
func main() {
	dat := []int{1, 2, -2, 3}
	rCh := make(chan int, len(dat))
	for _, v := range dat {
		worker(v, rCh)
	}
	for i := 0; i < len(dat); i++ {
		fmt.Println(<-rCh) // 1. closed channel will not block! // HL
	}
}

// 20 OMIT
