package main

import (
	"math"
	"reflect"
	"testing"
)

func TestIsValidNumber(t *testing.T) {
	tests := []struct {
		name     string
		number   int64
		expected bool
	}{
		{"positive number", 5, true},
		{"negative number", -5, true},
		{"zero", 0, true},
		{"large positive number", 1000000, true},
		{"large negative number", -1000000, true},
		{"max safe positive", math.MaxInt64 / 10, true},
		{"min safe negative", math.MinInt64 / 10, true},
		{"overflow positive", math.MaxInt64/10 + 1, false},
		{"overflow negative", math.MinInt64/10 - 1, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsValidNumber(tt.number)
			if result != tt.expected {
				t.Errorf("IsValidNumber(%d) = %v, expected %v", tt.number, result, tt.expected)
			}
		})
	}
}

func TestMultiply(t *testing.T) {
	tests := []struct {
		name     string
		a        int64
		b        int64
		expected int64
	}{
		{"positive numbers", 5, 3, 15},
		{"negative and positive", -5, 3, -15},
		{"both negative", -5, -3, 15},
		{"multiply by zero", 5, 0, 0},
		{"multiply by one", 5, 1, 5},
		{"large numbers", 1000, 1000, 1000000},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Multiply(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("Multiply(%d, %d) = %d, expected %d", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

func TestGenerateMultiplicationTable(t *testing.T) {
	tests := []struct {
		name     string
		number   int64
		expected []int64
	}{
		{
			name:     "table of 5",
			number:   5,
			expected: []int64{5, 10, 15, 20, 25, 30, 35, 40, 45, 50},
		},
		{
			name:     "table of 1",
			number:   1,
			expected: []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
		{
			name:     "table of 0",
			number:   0,
			expected: []int64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
		{
			name:     "table of negative number",
			number:   -3,
			expected: []int64{-3, -6, -9, -12, -15, -18, -21, -24, -27, -30},
		},
		{
			name:     "table of 12",
			number:   12,
			expected: []int64{12, 24, 36, 48, 60, 72, 84, 96, 108, 120},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := GenerateMultiplicationTable(tt.number)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("GenerateMultiplicationTable(%d) = %v, expected %v", tt.number, result, tt.expected)
			}
		})
	}
}

func TestGenerateMultiplicationTableLength(t *testing.T) {
	numbers := []int64{1, 5, 10, 100, -50}
	for _, num := range numbers {
		result := GenerateMultiplicationTable(num)
		if len(result) != 10 {
			t.Errorf("GenerateMultiplicationTable(%d) returned %d elements, expected 10", num, len(result))
		}
	}
}

func BenchmarkGenerateMultiplicationTable(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GenerateMultiplicationTable(12)
	}
}

func BenchmarkMultiply(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Multiply(12, 10)
	}
}
