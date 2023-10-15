
package main

import (
	"fmt"
//	"strconv"
//	"strings"
)

func main() {

	// Find largest prime factor of 600851475143

	n := 600851475143
	fmt.Println("Prime factors of :", n, "\n")
	largest:=71
	for i:=largest;i<(n/70);i++ {
		if n % i == 0{
			if isPrime(i){
				fmt.Printf(" %d", i)
				largest = i
			}
		}
	}
	fmt.Println("\nAnswer: %d", largest)

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
