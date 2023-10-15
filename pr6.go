
package main

import (
	"fmt"
)

func main() {

	// The difference between the sum of the squares (1^2 + 2^2...) and the square of sums (1 +2+..)^2
	// for {1..100}

	fmt.Println("Sum Square problem 100")

	sumOfSquare := 0
	tmpSum := 0
	for i:=1;i<=100;i++{
		sumOfSquare += i*i
		tmpSum += i
	}
	squareOfSum := tmpSum * tmpSum

	fmt.Printf("sumOfSquare = %d, squareOfSum = %d, difference = %d", sumOfSquare, squareOfSum, squareOfSum-sumOfSquare)

	//fmt.Printf("\ntotal = %0.10f", total)
}

//func threeConsecutive(n int) bool {
//	return true
//}

