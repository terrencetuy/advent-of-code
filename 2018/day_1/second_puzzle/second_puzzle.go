package main

import (
	"fmt"
)

func main() {
	freq, err := GetFirstSecondFrequency("input.txt")
	if err == nil {
		fmt.Println(freq)
	} else {
		fmt.Println(err)
	}
}
