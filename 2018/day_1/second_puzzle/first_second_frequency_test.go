package main

import (
	"testing"
)

func TestInputFileExists(t *testing.T) {
	_, err := GetFirstSecondFrequency("test_data/test_empty.txt")
	if err == nil {
		t.Errorf("Program should error on non-existant file")
	}
}

func TestSecondFrequencyFound(t *testing.T) {
	freq, _ := GetFirstSecondFrequency("test_data/test_basic_frequencies.txt")

	if freq != float64(0) {
		t.Errorf("Expected frequency of 0, got %f", freq)
	}
}

func TestFrequencyFoundAfterMultipleCycles(t *testing.T) {
	table := []struct {
		filename string
		expected float64
	}{
		{"test_data/test_cycle_1.txt", 10},
		{"test_data/test_cycle_2.txt", 5},
		{"test_data/test_cycle_3.txt", 14},
	}

	for _, test := range table {
		freq, _ := GetFirstSecondFrequency(test.filename)
		if  freq != test.expected {
			t.Errorf("Expected frequency of %f, got %f", test.expected, freq)
		}
	}
}
