package main

import "fmt"

type HasName interface {
	GetName() string
}

func sayHello(hasName HasName) {
	fmt.Println("Hello", hasName.GetName())
}

type Nama struct {
	nama string
}

func (person Nama) GetName() string {
	return person.nama
}
func main() {
	var fair Nama
	fair.nama = "FAIR"
	sayHello(fair)
}
