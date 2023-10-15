
package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	// Largest palindrome that is the product of 2 3digit numbers

	fmt.Println("Largest palindrome that is product of 2 3-digit numbers")

	fi := 1
	fj := 1
	fpal := 0
	for i := 100; i < 1000; i++ {
		for j := 100; j < 1000; j++ {
			product := i*j
			str := strconv.Itoa(product)
			sa := strings.Split(str, "")
			pal := true
			for k:=0;k<len(sa);k++{
				// determine if palindrome
				if sa[k] != sa[len(sa)-k-1] {
					pal = false
					break
				}
			}
			if pal == true && product > fpal{
				fmt.Printf("\n (larger) palindrome detected %s", str)
				fmt.Printf("\ti = %d, j = %d", i, j)
				fi = i
				fj = j
				fpal = product
			}
		}

	}

	fmt.Println("\nAnswer: %d, %d, %d", fi, fj, fpal)

	//fmt.Printf("\ntotal = %0.10f", total)
}

//func threeConsecutive(n int) bool {
//	return true
//}

