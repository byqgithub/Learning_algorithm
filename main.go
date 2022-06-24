package main

import (
	"fmt"
)

func rpow(x float64, n int64) float64  {
	fmt.Printf("x: %v, n: %v\n", x, n)
	if n == 0 {
		return 1
	}

	t := rpow(x, n / 2)
	fmt.Printf("t %v, n %v, n%%2 %v\n", t, n, n%2)
	if n % 2 > 0 {
		//fmt.Printf("n %% 2\n")
		return x * t * t
	}

	return t * t
}

func pow(x float64, n int64) float64 {
	var ans float64 = 1
	for {
		if n & 1 > 0 {
			ans = ans * x
		}
		x = x * x
		n = n >> 1
		if n <= 0 {
			break
		}
	}
	return ans
}

func main() {
	x := float64(7)
	n := int64(1)
	result := rpow(x, n)
	fmt.Printf("recursion pow result: %v\n", result)

	result = pow(x, n)
	fmt.Printf("non-recursion pow result: %v\n", result)
}
