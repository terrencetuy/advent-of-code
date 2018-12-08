package main

import (
	"os"
	"bufio"
	"strconv"
)

func GetFirstSecondFrequency(filename string) (float64, error) {
	// initialize
	past_frequencies := []float64{0}
	file, err := os.Open(filename)
	if err != nil {
		return -1, err
	}

	for {
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			// Get frequency to check
			frequency_change, err := strconv.ParseFloat(scanner.Text(), 64)
			last_frequency := past_frequencies[len(past_frequencies) - 1]
			current_frequency := last_frequency + frequency_change

			// Check if frequency has been encountered
			for _, freq_being_checked := range past_frequencies {
				if current_frequency == freq_being_checked {
					return current_frequency, err
				}
			}

			past_frequencies = append(past_frequencies, current_frequency)
		}
		// Second frequency still not found, start list over
		file.Seek(0,0)
	}
}
