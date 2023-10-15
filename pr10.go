
package main

import (
	"fmt"
)




// attempt 2 using sieve of eratosthenes
func main() {
	// sum of all primes below 2,000,000

	fmt.Println("Sum of all primes below 2,000,000")

	// make an array of numbers from 1..2000000
	nums := make([]int, 2000000)
	// setting i to i is faster than i = i-1. 2000000 isnt a prime anyway. and im going to sum the array so it doesnt matter that 0 is in there
	for i:=0;i<2000000;i++{
		nums[i] = i
	}

//	fmt.Println("nums array in the beginning:")
//	fmt.Printf("\n\n%v", nums)
	// set index=1 (val 1) to 0 bc not a prime
	nums[1] = 0

	max :=2000000
	l := 1000000
	for i:=2;i<l;i++{
		//nums[i+i, i+2i, i+3i...]
		j :=1
		for true {
			ni := i + (i*j)
			if ni >= max {
				break
			}
			nums[ni] = 0
			j++
		}
	}
	//fmt.Println("nums array after seive of eratosthenes:")
	//fmt.Printf("\n\n%v", nums)

	// holy shit... that was almost instantaneous... the algorithm below took over a minute to get to 0.7 percent of all numbers... wild....

	// ok wait.. the answer was 142913828923 which was wrong... 
	// now, we sum all the numbers
	sum := 0
	for i := range nums {
		sum += nums[i]
	}





	fmt.Println("Answer is: ", sum)
}







//func main() {
//
//	// sum of all primes below 2,000,000
//
//	fmt.Println("Sum of all primes below 2,000,000")
//
//	t := 2
//	n := 2
//
//	max := float64(2000000)
//	for true {
//		n++
//		if isPrime(n) {
//			t += n
//			//fmt.Printf("\n\tprime #%d: %d", i, n)
//		}
//		if n == 2000000 {
//			break
//		}
//		if n %  200000 == 0{
//			fmt.Printf("\n %0.2f percent complete", float64(n)/max)
//		}
//	}
//
//	fmt.Println("\nAnswer: %d", t)
//
//}

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
