package main

import (
	"reflect"
	"testing"
)

func TestFindCommonDivisors(t *testing.T) {
	testTable := []struct {
		nums     []int
		expected []int
	}{
		{
			nums:     []int{42, 12, 18},
			expected: []int{1, 2, 3, 6},
		},
		{
			nums:     []int{5, 7, 9, 35},
			expected: []int{1},
		},
		{
			nums:     []int{8, 12, 16},
			expected: []int{1, 2, 4},
		},
	}

	for _, testCase := range testTable {
		result := FindCommonDivisors(testCase.nums)
		if !reflect.DeepEqual(result, testCase.expected) {
			t.Errorf("Incorrect result for %v. Expected %v, got %v", testCase.nums, testCase.expected, result)
		}
	}
}
