package main

import "fmt"

func main() {
	var name string = "sulaiman baik hati"
	gender := "semendawai" // inisiasi awal, deklarasi variable menggunakan :=

	var (
		firstName = "sulaiman"
		lastName  = "semendawai"
	)

	// var no use = error
	// const no use is ok
	
	const tempatLahir string = "malang"
	const tanggalLahir = 35

	fmt.Println(name)
	fmt.Println(gender)
	fmt.Println(firstName)
	fmt.Println(lastName)
}
