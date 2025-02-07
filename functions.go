package main

import (
	"math"
	"math/big"
)

// isPrime checks if a number is prime.
//
// A number is prime if it is only divisble by 1 and itself.
//
// Retuns a bool.
func isPrime(n int) bool {
	return big.NewInt(int64(n)).ProbablyPrime(0)
}

// isPerfect checks if a number is perfect.
//
// A number is perfect if the sum of its positive proper divisor excluding itself.
//
// Retuns a bool.
func isPerfect(n int) bool {
	if n <= 1 {
		return false
	}
	total := 0
	for x := 1; x < n; x++ {
		if n%x == 0 {
			total += x
		}
	}
	return total == n
}

// isArmstrong checks if a number is an armstrong number.
//
// A number is regarded as an armstrong number if the sum of its
// digits each raised to the power of the number of digits equals itself.
//
// returns a bool.
func isArmstrong(n int) bool {
	original, total := n, 0
	digitCount := int(math.Log10(float64(n)) + 1)
	for temp := n; temp > 0; temp /= 10 {
		digit := temp % 10
		total += int(math.Pow(float64(digit), float64(digitCount)))
	}
	return original == total
}

// isOdd checks if a number is odd.
//
// A number is odd if it is not diviisble by 2 without having remainders.
//
// Retuns a bool.
func isOdd(n int) bool {
	return n%2 != 0
}

// compute properties returns a slice for the json response property field.
//
// Returns a slice of string with allowed values ["armstrong", "odd", "even"]
func computeProperties(isArmstrong bool, isOdd bool) []string {
	var props []string
	if isArmstrong {
		props = append(props, "armstrong")
	}
	if isOdd {
		props = append(props, "odd")
	} else {
		props = append(props, "even")
	}
	return props
}
