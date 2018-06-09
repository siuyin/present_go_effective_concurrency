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
			log.Println("can't handle negative numbers!")
			done <- true // I give up! I'm done! // HL
			return
		}
		time.Sleep(1 * time.Second)
		ch <- inp * 2
		done <- true // Finished working. I'm done! // HL
	}(i)
}

// 15 OMIT
func main() {
	dat := []int{1, 2, -2, 3}
	done := make(chan bool) // HL
	rCh := make(chan int, len(dat))
	for _, v := range dat {
		worker(v, rCh, done) // HL
	}
	reportedIn := 0
main: // HL
	for { // loop forever until told to break
		select { // HL
		case result := <-rCh:
			fmt.Println(result)
		case <-done:
			reportedIn++
			if reportedIn >= len(dat) {
				break main // HL
			}
		}
	}
}

// 20 OMIT
