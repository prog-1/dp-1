package main

import "testing"

func TestMaxValue(t *testing.T) {
	for _, tc := range []struct {
		W             int
		value, weight []int
		want          int
	}{
		{4, []int{1, 2, 3}, []int{4, 5, 1}, 3},
		{3, []int{1, 2, 3}, []int{4, 5, 6}, 0},
		{5, []int{1, 2, 3}, []int{4, 5, 1}, 4}, // fails
	} {
		got := maxValue(tc.value, tc.weight, tc.W)
		if got != tc.want {
			t.Errorf("maxValue(%v, %v, %v) = %v, want = %v", tc.value, tc.weight, tc.W, got, tc.want)
		}
	}
}
