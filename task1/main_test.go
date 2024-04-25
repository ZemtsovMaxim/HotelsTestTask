package main

import "testing"

func TestSayItRight(t *testing.T) {
	testTable := []struct {
		number   int
		expected string
	}{
		{
			number:   25,
			expected: "25 компьютеров",
		},
		{
			number:   41,
			expected: "41 компьютер",
		},
		{
			number:   1048,
			expected: "1048 компьютеров",
		},
		{
			number:   100,
			expected: "100 компьютеров",
		},
		{
			number:   3,
			expected: "3 компьютера",
		},
		{
			number:   10,
			expected: "10 компьютеров",
		},
	}

	for _, testCase := range testTable {
		result := SayItRight(testCase.number)
		if result != testCase.expected {
			t.Errorf("Incorrect result for %d. Expected %s, got %s", testCase.number, testCase.expected, result)
		}
	}
}
