package main

import (
	"fmt"
	"strings"
)

func main() {
	testParam("hai dunia, aku datang.", "fair")

	nama, umur := belajarReturnFunc("Fair", 33)

	testKata := hasilKata("Jancuk", filterKataKotor)
	fmt.Println("test Kata: ", testKata)
	fmt.Println("usia", nama, "saat ini :", umur)

	fmt.Println(tambahSemua(1, 2, 3, 4))
}

func testParam(data string, nama string) {
	fmt.Println(data, nama)
}

// return function, return harus di define type data nya
func belajarReturnFunc(nama string, usia int) (string, int) {
	return nama, usia

}

// variadic function
func tambahSemua(number ...int) int {
	total := 0
	for _, value := range number {
		total += value
	}
	return total
}

// function using on anothers function as parameter
func filterKataKotor(kata string) string {
	if kata == "jancuk" {
		return "hayo bahasanya dijaga"
	} else {
		return kata
	}
}

func hasilKata(kata string, filter func(string) string) string {
	filterasi := filter(strings.ToLower(kata))
	return filterasi
}

// ==================================================
