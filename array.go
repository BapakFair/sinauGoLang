package main

import "fmt"

func main() {
	var nama [3]string // isi jumlah dari array harus di definisikan

	nama[0] = "fair"
	nama[1] = "sulaiman"
	nama[2] = "semendawai"
	//nama[3] = "lebih dari"

	// cara lain inisiasi array sekaligus isi
	nilai := [2]int{
		1,
		2,
	}
	fmt.Println(nama)
	fmt.Println(nilai)
}
