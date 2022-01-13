package main

import (
	"fmt"
)

func main() {
	belajarGo := true

	if semangat := 81; belajarGo && (semangat > 80) {
		fmt.Println("masa depan cerah")
	} else if belajarGo && (semangat <= 80) {
		fmt.Println("masa depan rodok cerah")
	} else {
		fmt.Println("masa depan pancet koyok biasane")
	}

	switch semangat := 90; belajarGo && (semangat > 80) {
	case true:
		{
			fmt.Println("masa depan cerah sekali")
		}
	case false:
		{
			fmt.Println("masa depan apik apik ae")
		}
	}
}
