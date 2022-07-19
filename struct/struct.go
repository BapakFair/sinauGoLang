package main

import "fmt"

func main() {
	fair := Customer{
		Name:    "Fair",
		Address: "malang",
		Age:     33,
	}

	fmt.Println(fair)
}

type Customer struct { // uppercase di awal, budaya untuk struct
	Name, Address string
	Age           int
}
