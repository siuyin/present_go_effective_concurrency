package main

import (
	"fmt"
	"log"
	"time"
)

// 10 OMIT
func worker(i int, ch chan<- int, done chan<- bool) { // HL
	go func(inp int) {
		if inp < 0 {
			log.Printf("can't handle negative numbers:%v", inp)
			done <- true // I give up! I'm done! // HL
			return
		}
		time.Sleep(1 * time.Second)
		//RS OMIT
		ch <- inp * 2
		done <- true // Finished working. I'm done! // HL
		//RE OMIT
	}(i)
}

// 15 OMIT
func main() {
	dat := []int{1, 2, -2, 3}
	done := make(chan bool) // HL
	rCh := make(chan int)   // 1. Buffered channel?
	for _, v := range dat {
		worker(v, rCh, done) // HL
	}
	reportedIn := 0
main: // HL
	for { // loop forever until told to break
		// SS OMIT
		select { //2. Both channels ready? // HL
		case result := <-rCh:
			fmt.Println(result)
		case <-done:
			reportedIn++
			if reportedIn >= len(dat) {
				break main // HL
			}
		}
		// SE OMIT
	}
	time.Sleep(100 * time.Millisecond)
}

// 20 OMIT
