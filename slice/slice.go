package main

import "fmt"

// type data slice, mirip array,
// pointer, len, capacity
// array[low, high]
// array[low:]
// array[:high]
// array[:]

// func in slice :
// len(slice) = panjang
// cap(slice) = kapasitas
// append(slice, data) = menambah data ke posisi terakhir slice, jika kapasitas penuh maka akan membuat array baru
// make([]TypeData, length, capacity) = membuat slice baru
// copy(destination, source) = menyalin slice, dari source ke destination

func main() {

	hari := [...]string{
		"senin",
		"selasa",
		"rabu",
		"kamis",
		"jum'at",
		"sabtu",
		"minggu",
	}
	slice1 := hari[:5]
	fmt.Println(slice1)

	iniArray := [...]int{1, 2, 3, 4, 5}
	iniSlice := []int{1, 2, 3, 4, 5}

	fmt.Println(iniArray)
	fmt.Println(iniSlice)
}
