package main

import (
	"bufio"
	"os"
	"strconv"
)

func GetChronalCalibration(filename string) (float64, error) {
	file, err := os.Open(filename)
	scanner := bufio.NewScanner(file)
	sum := float64(0)
	for scanner.Scan() {
		value, err := strconv.ParseFloat(scanner.Text(), 64)
		if err == nil {
			sum += value
		}
	}
	return sum, err
}
