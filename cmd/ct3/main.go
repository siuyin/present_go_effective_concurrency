package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

// 10 OMIT
func getPrice(ctx context.Context, prodID string) <-chan float32 {
	ch := make(chan float32)
	go func() {
		price := float32(1.23)
		userID := ctx.Value(userIDKey{}) // HL
		ch <- price
		log.Printf("user %v requested price for %v, replied %v\n", userID, prodID, price)
	}()
	return ch
}
func getAddr(ctx context.Context, postcode string) <-chan string {
	ch := make(chan string)
	go func() {
		addr := "13 Elm Street"
		userID := ctx.Value(userIDKey{}) // HL
		ch <- addr
		log.Printf("user %v requested nearest store to %v, replied %v\n", userID, postcode, addr)
	}()
	return ch
}

// 20 OMIT
type userIDKey struct{} // golang specific trick to avoid memory allocation

func main() {
	bgCtx := context.Background()                             // get a context
	valCtx := context.WithValue(bgCtx, userIDKey{}, "usr123") // HL
	price := getPrice(valCtx, "prod45")
	addr := getAddr(valCtx, "123456")
	reportedIn := 0
	checkInCh := make(chan struct{}, 2)
	// 25 OMIT
main:
	for {
		select {
		case p := <-price:
			fmt.Printf("price of prod45 is $%v\n", p)
			checkInCh <- struct{}{}
		case a := <-addr:
			fmt.Printf("address of store nearest postcode 123456 is %v\n", a)
			checkInCh <- struct{}{}
		case <-checkInCh: // HL
			reportedIn++
			if reportedIn >= 2 {
				break main
			}
		}
	}
	time.Sleep(100 * time.Millisecond)
}

// 30 OMIT
