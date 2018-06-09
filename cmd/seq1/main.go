package main

import (
	"fmt"
	"time"
)

// 10 OMIT
func costlyOp(i int) int {
	time.Sleep(1 * time.Second)
	return i * 2
}
func main() {
	dat := []int{1, 2, 3}
	for _, v := range dat {
		fmt.Println(costlyOp(v)) // HL
	}
}

// 20 OMIT
