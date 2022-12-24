package main

import (
	"fmt"
)

type Key struct {
	cap int
	i   int
}

func main() {
	vs := []int{2, 7, 4, 6, 5} //item values
	ws := []int{3, 2, 5, 2, 3} //item weights
	cap := 5                   //knapsack capacity (w)
	var cache map[Key]int      //declaring cache map
	fmt.Println(knapsack(vs, ws, cap, 0, cache))
}

//returns max value of items, which weight is <= cap
func knapsack(vs, ws []int, cap, i int, cache map[Key]int) (res int) {
	if res, ok := cache[Key{cap, i}]; ok {
		return res
	}
	if i >= len(ws) || cap <= 0 { //base case
		return 0
	}
	if ws[i] > cap { //if w of i item > cap -> skip i item
		res = knapsack(vs, ws, cap, i+1, cache) //calculate choices with i item excluded
	} else {
		incl := vs[i] + knapsack(vs, ws, cap-ws[i], i+1, cache) //calculate choices with i item included
		excl := knapsack(vs, ws, cap, i+1, cache)               //calculate choices with i item excluded
		if incl > excl {                                        //choose max from two
			res = incl
		} else {
			res = excl
		}
	}
	cache[Key{cap, i}] = res
	return res
}
