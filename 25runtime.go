package main

import (
	"fmt"
	"math"
)

func main() {

	var numinputs, num int

	fmt.Scan(&numinputs)

	for j := 0; j < numinputs; j++ {
		fmt.Scan(&num)

		var squareroot, i int
		prime := true
		squareroot = int(math.Sqrt(float64(num)))

		for i = 2; i <= squareroot; i++ {
			if num%i == 0 {
				prime = false
				break
			}
		}

		if prime && num != 1 {
			fmt.Println("Prime")
		} else {
			fmt.Println("Not prime")
		}

	}

}
