package main

import (
	"fmt"
)

func main() {
	sum, err := GetChronalCalibration("input.txt")
	if err == nil {
		fmt.Println(sum)
	} else {
		fmt.Println(err)
	}
}
