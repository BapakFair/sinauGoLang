package main

import "fmt"

func main() {
	var integer32 int32 = 300123
	var integer64 int64 = int64(integer32)
	var integer8 int8 = int8(integer32)

	fmt.Println(integer32)
	fmt.Println(integer64)
	fmt.Println(integer8)
}
