
package main

import (
	"fmt"
//	"strconv"
//	"strings"
)

func main() {

	// What is the 10001st prime number?

	fmt.Println("Largest 10001st prime number")

	i := 1 // index for nth prime number
	lp := 2 // largest prime number so far
	n := 2

	for true {
		n++
		if isPrime(n) {
			i++
			lp = n
			fmt.Printf("\n\tprime #%d: %d", i, n)
			if i == 10001 {
				break
			}
		}
	}

	fmt.Println("\nAnswer: %d", lp)

	//fmt.Printf("\ntotal = %0.10f", total)
}

//func threeConsecutive(n int) bool {
//	return true
//}


func isPrime(n int) bool {
	for i:=2; i <n; i++ {
		if n % i == 0 {
			return false
		}
	}
	return true
}
