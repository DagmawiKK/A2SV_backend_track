package main

import (
	"math"
	"testing"
)

func TestCalculateAverage(t *testing.T) {
	tests := []struct {
		name     string
		grades   []float64
		expected float64
	}{
		{"All 100s", []float64{100, 100, 100}, 100},
		{"All 0s", []float64{0, 0, 0}, 0},
		{"Mixed grades", []float64{80, 90, 100}, 90},
		{"Single subject", []float64{73}, 73},
		{"Decimal grades", []float64{87.5, 91.3}, (87.5 + 91.3) / 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := calculateAverage(tt.grades)
			if math.Abs(result-tt.expected) > 0.0001 {
				t.Errorf("calculateAverage(%v) = %v; want %v", tt.grades, result, tt.expected)
			}
		})
	}
}
