package main

import "fmt"

func main() {
	startApp(false)
}

func endApp() {
	message := recover() // recover menangkap error dari panic
	if message != nil {
		fmt.Println("error dengan message", message)
	}
	fmt.Println("app berhenti")
}

func startApp(error bool) {
	defer endApp() // defer di eksekusi setelah sebuah blok kode selesai ter eksekusi, error maupun running well
	if error {
		panic("ERROR BOSKU") // mengirimkan panic message
	}
	fmt.Println("Aplikasi aman jaya")
}
