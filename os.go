package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	fmt.Println(args)

	Name, err := os.Hostname()
	if err == nil {
		fmt.Println("Hostname :", Name)
	} else {
		fmt.Println("Error :", err)
	}

	//os.Getenv() untuk mengambil data .env
	
}
