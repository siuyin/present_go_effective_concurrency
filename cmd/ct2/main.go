package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

// 10 OMIT
func work(ctx context.Context, id string) <-chan string {
	ch := make(chan string)
	go func() {
		rand.Seed(time.Now().UnixNano())
		ch2 := make(chan string)
		go func() {
			sleepTime := time.Duration(rand.Intn(90)+10) * time.Millisecond
			fmt.Printf("%v: I need %v\n", id, sleepTime)
			time.Sleep(sleepTime)
			ch2 <- fmt.Sprintf("%v done!", id) // HL
		}()
		select {
		case <-ctx.Done(): // HL
			fmt.Printf("worker %v exiting. I've been told to stop work! -- %v\n", id, ctx.Err())
		case r := <-ch2: // HL
			ch <- r
			fmt.Printf("worker %v sent off results!\n", id)
		}
	}()
	return ch
}

// 20 OMIT
func main() {
	bgCtx := context.Background()              // get a context
	ccCtx, cancel := context.WithCancel(bgCtx) // HL
	defer cancel()
	a := work(ccCtx, "A")
	b := work(ccCtx, "B")
	select {
	case msg := <-a:
		fmt.Printf("main from A: %v\n", msg)
	case msg := <-b:
		fmt.Printf("main from B: %v\n", msg)
	}
	fmt.Println("Telling workers to stop work.")
	cancel()
	time.Sleep(100 * time.Millisecond) // so that we can see 'slow' messages
}

// 30 OMIT
