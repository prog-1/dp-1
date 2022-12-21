package main

import (
	"testing"
)

func TestKnapsack(t *testing.T) {
	for _, tc := range []struct {
		name   string
		cap    int
		vs, ws []int
		want   int
	}{
		{"1", 4, []int{1, 2, 3}, []int{4, 5, 1}, 3},
		{"2", 3, []int{1, 2, 3}, []int{4, 5, 6}, 0},
		{"3", 5, []int{2, 7, 4, 6, 5}, []int{3, 2, 5, 2, 3}, 11},
		{"4", 8, []int{2, 7, 4, 6, 5}, []int{3, 2, 5, 2, 3}, 18},
		{"5", 1, []int{}, []int{}, 0},
		{"6", 0, nil, nil, 0},
	} {
		t.Run(tc.name, func(t *testing.T) {
			got := Knapsack(tc.vs, tc.ws, tc.cap)
			if got != tc.want {
				t.Errorf("got = %v, want = %v", got, tc.want)
			}
		})
	}
}
