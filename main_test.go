package main

import "testing"

// valid:   9663412039983 | 4974754173435
// invalid: 3012702219781 | 2492419244542

func TestCalculateChecksum(t *testing.T) {
	tcs := []struct {
		expectedChecksum int
		input            []int
	}{
		{
			expectedChecksum: 5,
			input:            []int{0, 8, 7, 9, 1, 2, 2, 0, 7, 2, 1, 0, 3},
		},
		{
			expectedChecksum: 3,
			input:            []int{0, 4, 5, 4, 4, 2, 9, 1, 4, 2, 9, 4, 2},
		},
		{
			expectedChecksum: 5,
			input:            []int{0, 3, 4, 3, 7, 1, 4, 5, 7, 4, 7, 9, 4},
		},
		{
			expectedChecksum: 3,
			input:            []int{0, 8, 9, 9, 3, 0, 2, 1, 4, 3, 6, 6, 9},
		},
	}

	for _, tc := range tcs {
		actualChecksum := CalculateChecksum(tc.input)

		if tc.expectedChecksum != actualChecksum {
			t.Error("expect", tc.expectedChecksum, "actual", actualChecksum, "input", tc.input)
		}
	}
}
