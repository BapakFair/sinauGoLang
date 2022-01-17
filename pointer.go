package main

import (
	"fmt"
)

type Address struct {
	City, Province, Country string
}

func changeCountryToGerman(address *Address) {
	address.Country = "Germany"
}

// pass by value Vs pass by reference
func main() {
	address1 := Address{"Malang", "Jawa Timur", "Indonesia"}
	address2 := address1  // pass by value
	address3 := &address1 // pass by reference

	address2.City = "Bandung"
	address3.City = "Seoul"

	*address3 = Address{"Surabaya", "Jawa Utara", "Delhi"}

	address4 := new(Address)
	address5 := address4
	address5.City = "Denpasar"

	fmt.Println(address1)
	fmt.Println(address2)
	fmt.Println(address3)
	fmt.Println(address5)

	alamat := Address{
		City:     "Kediri",
		Province: "Jawa Selatan",
		Country:  "Indonesia",
	}
	changeCountryToGerman(&alamat)
	fmt.Println(alamat)
}
