package main

import "testing"

func TestTask(t *testing.T) {
	type Input struct {
		values  []int
		weights []int
		cap     int
	}
	for _, tc := range []struct {
		input Input
		want  int
	}{
		{Input{nil, nil, 0}, 0},
		{Input{[]int{}, []int{}, 1}, 0},
		{Input{[]int{1, 2, 3}, []int{4, 5, 1}, 3}, 3},
		{Input{[]int{2, 3, 1}, []int{4, 3, 2}, 5}, 4},
		{Input{[]int{1, 2, 3}, []int{4, 5, 6}, 3}, 0},
		{Input{[]int{2, 7, 4, 6, 5}, []int{3, 2, 5, 2, 3}, 8}, 18},
		{Input{[]int{2, 2, 4, 5, 3}, []int{3, 1, 3, 4, 2}, 7}, 10},
	} {
		if got := knapsack(tc.input.cap, tc.input.values, tc.input.weights); got != tc.want {
			t.Errorf("got = %v, want = %v", got, tc.want)
		}
	}

}
