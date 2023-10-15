
package main

import (
	"fmt"
	"math"
)

func main() {

	// find the pythagorean triplet that sums to 1000
	// pythagorean triplet = a^2 + b^2 = c^2
	var a, b, cf int
	fmt.Println("Pythagorean triplet that sums to 1000")
	var w bool
	w = false
	for i:=1; i<1000; i++ {
		for j:=1; j< 1000; j++ {
			// find c
			c2 := i*i + j*j
			c := int64(math.Sqrt(float64(c2)))
		//	fmt.Println("c = ", c, ", c^2 = ", c2, "c*c = ", c*c, " while i = ", i, ", j = ", j)
			if c*c == int64(c2) {
				// found a natural square
				fmt.Println("Found a triplet. A = ", i, ", B = ", j, ", C = ", c)
				fmt.Println("sum = ", int(i) + int(j) + int(c))
				if i + j + int(c) == 1000 {
					fmt.Println("Found the winning combo!")
					a = i
					b = j
					cf = int(c)
					w = true
					break
				}
			}
		}
		if w == true {
			break
		}
	}
	fmt.Println("Answer = ", a*b*cf)


	//fmt.Printf("\ntotal = %0.10f", total)
}

//func threeConsecutive(n int) bool {
//	return true
//}

