
package main

import (
	"fmt"
)

func main() {

	// smallest positive number that is evenly divisible by all numbers from 1 to 20

	fmt.Println("Smallest positive number divisible by {1..20}")

	n := 20

	for true {
		winner := true
		for i:=2; i<=20; i++ {
			if n%i !=0 {
				winner = false
				break
			}
		}
		if winner == true {
			fmt.Printf("\nWinning number: %d", n)
			break
		}
		n++
	}


	//fmt.Printf("\ntotal = %0.10f", total)
}

//func threeConsecutive(n int) bool {
//	return true
//}

