package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	loopTest()

	t := time.Now()
	elapsed := t.Sub(start)
	fmt.Println(elapsed)

}

func loopTest() {
	for counter := 0; counter <= 10; counter++ {
		fmt.Println(counter)
	}

	slice := []string{"senin", "selasa", "rabu", "kamis", "jum'at", "sabtu", "minggu"}
	for i, hari := range slice {
		fmt.Println("index ke :", i, "adalah hari : ", hari)
	}

	// gunakan underscore untuk declare bahwa variable tidak dibutuhkan
	for _, hari := range slice {
		fmt.Println("hari : ", hari)
	}
	// =================================================================
}
