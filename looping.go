package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	counter := 0
	for counter <= 10000000 {
		fmt.Println(counter)
		counter++
	}

	t := time.Now()
	elapsed := t.Sub(start)
	fmt.Println(elapsed)
}
