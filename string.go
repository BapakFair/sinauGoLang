package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var name string = "sulaiman baik hati"
	fmt.Println(name)

	tesString := strings.Contains(name, "iman") // check apakah mengandung string
	fmt.Println(tesString)

	testString := "100"
	stringToInteger, err := strconv.ParseInt(testString, 10, 64) // base = 10 for decimal, 2 for binary, 8 for octa, 16 for hexa
	if err == nil {
		fmt.Println(stringToInteger)
	} else {
		fmt.Println(err.Error())
	}

	fmt.Println(strconv.Atoi("20"))
	fmt.Println(strconv.Itoa(333))
}

// noted :
// 		len = length of string
