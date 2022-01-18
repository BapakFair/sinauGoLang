package main

import (
	"container/list"
	"fmt"
)

func main() {
	data := list.New()
	data.PushBack("Sulaiman")
	data.PushBack("yo iki")

	fmt.Println(data)
	fmt.Println(data.Back().Value)
	fmt.Println(data.Front().Value)

}
