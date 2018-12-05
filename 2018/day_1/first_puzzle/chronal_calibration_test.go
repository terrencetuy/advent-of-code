package main

import (
	"testing"
)

func TestInputFileExists(t *testing.T) {
	_, err := GetChronalCalibration("test_data/test_empty.txt")
	if err == nil {
		t.Errorf("Program should error on non-existant file")
	}
}

func TestGetResultingFrequencyIsSum(t *testing.T) {
	table := []struct {
		filename string
		expected float64
	}{
		{"test_data/test.txt", 6},
		{"test_data/test_negative.txt", -1},
		{"test_data/test_float.txt", 3.14},
	}

	for _, test := range table {
		result, _ := GetChronalCalibration(test.filename)
		if result != test.expected {
			t.Errorf("Expected %f, got %f", test.expected, result)
		}
	}
}
