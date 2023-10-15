
package main

import (
	"fmt"
//	"strconv"
//	"strings"
)

func main() {

	// Sum of all even numbers in fibonacci sequence upto 4,000,000

	sum := 2 // add j now
	i := 1
	j := 2
	next := 0
	fmt.Println("Fibonacci sequence")
	fmt.Printf("1 1 2")
	for true {
		next = i + j
		fmt.Printf(" %d", next)
		if next > 4000000 {
			break
		}
		if next % 2 == 0 {
			sum += next
		}
		// reset i, j
		i = j
		j = next
	}
	
	fmt.Println("\nAnswer: %d", sum)


	//fmt.Printf("\ntotal = %0.10f", total)
}

//func threeConsecutive(n int) bool {
//	return true
//}

