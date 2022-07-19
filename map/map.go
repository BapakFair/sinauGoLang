package main

import "fmt"

func main() {
	person := map[string]string{
		"name":    "Fair",
		"address": "sinisaja",
	}
	person["ketiga"] = "coba insert value ketiga"

	fmt.Println(person)
	fmt.Println(person["ketiga"])
}
