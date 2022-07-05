package main

import "fmt"

func main() {
	fmt.Println("Hello world!")

	var integer int = 1
	fmt.Println(integer)

	var integer2, string = 10, "string"
	fmt.Println(integer2, string)

	booleanVariable := false
	fmt.Println("boolean value :", booleanVariable)

	// pointer
	x := 1
	p := &x
	fmt.Println("pointer value :", p)
	fmt.Println("pointer value :", *p)

	// type
	type fahrenheit float64
	type celsius float64

	var f fahrenheit = 32
	var c celsius = 0
	fmt.Println(f, c)
}
