package main

import "fmt"

func main() {
	vs := []int{2, 7, 4, 6, 5} //item values
	ws := []int{3, 2, 5, 2, 3} //item weights
	cap := 5                   //knapsack capacity (w)
	fmt.Println(Knapsack(vs, ws, cap))
}

//returns sum of   max subset of val, where wt <= cap
func Knapsack(vs, ws []int, cap int) (max int) {
	var cs, cw int      //cur sum & cur weight
	for i := range vs { //for each element
		for j := i; j < len(vs); j++ { //for each el from the right of i
			if cw+ws[j] > cap { //if we went out of capacity with new el
				if max < cs { //compare results
					max = cs
				}
				cs, cw = 0, 0 //reset counters
			} else { // if cw + new el <= cap
				cs += vs[j] //ad new el to cur sum
				cw += ws[j] //ad new el weight to cur weight
			}
		}
	}
	return max
}
