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
		{Input{[]int{1, 2, 3}, []int{4, 5, 1}, 3}, 3},
		{Input{[]int{2, 3, 1}, []int{4, 3, 2}, 5}, 4},
		{Input{[]int{1, 2, 3}, []int{4, 5, 6}, 3}, 0},
	} {
		if got := solveTask(tc.input.values, tc.input.weights, tc.input.cap); got != tc.want {
			t.Errorf("got = %v, want = %v", got, tc.want)
		}
	}

}
