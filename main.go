package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var N, W int
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscan(reader, &N, &W)
	values := make([]int, N)
	weight := make([]int, N)
	for i := 0; i < N; i++ {
		var tmp int
		fmt.Fscan(reader, &tmp)
		values[i] = tmp
	}
	for i := 0; i < N; i++ {
		var tmp int
		fmt.Fscan(reader, &tmp)
		weight[i] = tmp
	}
	fmt.Println(maxValue(values, weight, W))
}

func maxValue(values, weight []int, W int) int {
	var max int
	for i := range values {
		if weight[i] < W {
			a := values[i] + maxValue(append(values[:i], values[i+1:]...), append(weight[:i], weight[i+1:]...), W-weight[i])
			if a > max {
				max = a
			}
		}
	}
	return max
}
