package main

import (
	"fmt"
	"testing"
)

func TestParseMixedRadix(t *testing.T) {
	var tests = []struct {
		a, b   []int
		expect int
	}{
		{[]int{1, 2, 5}, []int{1, 1, 2}, 17},
		{[]int{7, 24, 60}, []int{0, 5, 32}, 332},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v, %v", tt.a, tt.b)
		t.Run(testname, func(t *testing.T) {
			result := ParseMixedRadix(tt.a, tt.b)
			if result != tt.expect {
				t.Errorf("got %d, expected %d", result, tt.expect)
			}
		})
	}
}

func TestFormatInt(t *testing.T) {
	var tests = []struct {
		radix  []int
		i      int64
		expect []int
	}{
		{[]int{1, 2, 5}, 17, []int{1, 1, 2}},
		{[]int{7, 24, 60}, 332, []int{0, 5, 32}},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v, %v", tt.radix, tt.i)
		t.Run(testname, func(t *testing.T) {
			result := FormatInt(tt.radix, tt.i)
			if !sliceEq(result, tt.expect) {
				t.Errorf("got %d, expected %d", result, tt.expect)
			}
		})
	}
}

func sliceEq(a, b []int) bool {
	if (a == nil) != (b == nil) {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
