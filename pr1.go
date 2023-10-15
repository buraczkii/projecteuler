
// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
//	"strconv"
//	"strings"
)

func main() {

	// Sum of all multiples of 3 & 5 below 1000

	sum := 0
	for i:= 1; i < 1000; i++ {
		if i %3 == 0 || i % 5 == 0 {
			sum += i
		}
	}

	fmt.Println("Answer: %d", sum)

	//fmt.Printf("\ntotal = %0.10f", total)
}

//func threeConsecutive(n int) bool {
//	return true
//}

