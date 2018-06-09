package main

import (
	"context"
	"fmt"
	"time"
)

// 10 OMIT
func workA(ctx context.Context) <-chan string {
	ch := make(chan string)
	go func() {
		// 12 OMIT
		ch2 := make(chan string)
		go func() {
			fmt.Printf("in workA: current time: %v\n", time.Now().Format("04:05.000"))
			dl, _ := ctx.Deadline()
			fmt.Printf("in workA: I have till %s to complete my work\n", dl.Format("04:05.000"))
			time.Sleep(50 * time.Millisecond)
			ch2 <- "A done!" // HL
		}()
		select {
		case <-ctx.Done(): // HL
			fmt.Printf("workA exiting. I've been told to stop work! -- %v\n", ctx.Err())
		case r := <-ch2: // HL
			ch <- r
			fmt.Printf("workA sent off results!\n")
		}
		// 13 OMIT
	}()
	return ch
}

// 20 OMIT
func main() {
	bgCtx := context.Background()                                    // get a context
	toCtx, cancel := context.WithTimeout(bgCtx, 60*time.Millisecond) // change me // HL
	defer cancel()
	a := workA(toCtx)
	select {
	case msg := <-a:
		fmt.Printf("main: %v\n", msg)
	case <-toCtx.Done():
		fmt.Printf("main exiting: %v\n", toCtx.Err())
	}
	time.Sleep(100 * time.Millisecond) // so that we can see 'slow' messages
}

// 30 OMIT
