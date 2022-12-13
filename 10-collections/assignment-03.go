/*
write a function "genPrimes" that generates the prime numbers from the given start to end and returns the primeNos
	the main function receives the prime nos and prints them
*/

package main

import "fmt"

func main() {
	primes := genPrimes(3, 100)
	fmt.Println(primes)
}

func genPrimes(start, end int) []int {
	primes := []int{}
	for no := start; no <= end; no++ {
		if isPrime(no) {
			primes = append(primes, no)
		}
	}
	return primes
}

func isPrime(no int) bool {
	for i := 2; i <= (no / 2); i++ {
		if no%i == 0 {
			return false
		}
	}
	return true
}
