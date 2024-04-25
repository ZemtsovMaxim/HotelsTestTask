package main

import (
	"reflect"
	"testing"
)

func TestFindPrimeNumbers(t *testing.T) {
	testTable := []struct {
		min      int
		max      int
		expected []int
	}{
		{
			min:      11,
			max:      20,
			expected: []int{11, 13, 17, 19},
		},
		{
			min:      1,
			max:      10,
			expected: []int{2, 3, 5, 7},
		},
		{
			min:      20,
			max:      30,
			expected: []int{23, 29},
		},
		{
			min:      50,
			max:      60,
			expected: []int{53, 59},
		},
		{
			min:      2,
			max:      2,
			expected: []int{2},
		},
	}

	for _, testCase := range testTable {
		result := FindPrimeNumbers(testCase.min, testCase.max)
		if !reflect.DeepEqual(result, testCase.expected) {
			t.Errorf("Incorrect result for [%d, %d]. Expected %v, got %v", testCase.min, testCase.max, testCase.expected, result)
		}
	}
}
