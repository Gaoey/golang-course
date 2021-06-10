package main

import (
	"fmt"
	"strconv"
)

func main() {
	result := squareArea(2)
	println("result: ", result)
	println(multiple(1, 2))

	// v, err := isEven("aa")
	// if err != nil {
	// 	fmt.Printf("err %s", err.Error())
	// }

	// fmt.Printf("isEven result: %v", v)
	// primes(100)
	println(power(2, 3))

}

func squareArea(a int) int {
	return a * a
}

func multiple(a, b int) (int, int) {
	return b, a + b
}

func isEven(a string) (bool, error) {
	b, err := strconv.Atoi(a)
	if err != nil {
		return false, err
	}

	return b%2 == 0, nil
}

func primes(a int) {
	for i := 1; i <= a; i++ {
		count := 0
		for j := 1; j <= i; j++ {
			if i%j == 0 {
				count++
			}
		}
		if count == 2 {
			fmt.Println(i)
		}
	}
}

func power(base, exponent int) int {
	result := 1
	for i := 1; i <= exponent; i++ {
		result = result * base
	}
	return result
}
