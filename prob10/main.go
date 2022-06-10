package main

import (
	"fmt"
	"math"
	"math/big"
	"time"
)

//Problem 10
// The sum of the primes below 10 is 2 + 3 + 5 + 7 = 17.
// Find the sum of all the primes below two million.

//This function outputs prime numbers smaller than given number n and sum of them
func PrimeFilter(n int64) ([]int64, *big.Int, bool) {

	if n > 10000000000 {
		fmt.Println("only support to 10^10 for time and memory problem.")
		return nil, big.NewInt(0), true
	}

	//step 1. Create a bool array each value set true -assume that all integers are prime.
	integers := make([]bool, n+1)
	for i := int64(2); i < n+1; i++ {
		integers[i] = true
	}

	//step 2. Eliminate all composite numbers.

	//Get max integer smaller than (root(n))
	maxInt := int64(math.Floor(math.Sqrt(float64(n))))
	// and set its multiples false.
	for p := int64(2); p < maxInt; p++ {
		//only when it was checked as
		if integers[p] == true {
			// Update all multiples of p
			for i := p * 2; i <= n; i += p {
				integers[i] = false
			}
		}
	}
	sum := big.NewInt(0)

	//step 3. return all prime numbers <= n and sum of them.
	var primes []int64
	for p := int64(2); p <= n; p++ {
		if integers[p] == true {
			primes = append(primes, p)
			sum.Add(sum, big.NewInt(p))

		}
	}

	return primes, sum, false
}

func main() {

	input := int64(2000000)
	fmt.Println("Input: ", input)
	t0 := time.Now()
	primes, sum, err := PrimeFilter(input)

	if err {
		fmt.Println("your input is too big, try another value")
		return
	}

	t1 := time.Now()
	fmt.Println("entire benchmark: ", t1.Sub(t0))
	fmt.Println("prime count:", len(primes))
	fmt.Println("sum of all primes:", sum)

}
