package fizzbuzz

import (
	"fmt"
	"math/rand"
)

func Say(n int) string {
	return "1"
}

func FizzBuzz(n int) string {
	if n%3 == 0 && n%5 == 0 {
		return "FizzBuzz"
	}

	if n%3 == 0 {
		return "Fizz"
	}

	if n%5 == 0 {
		return "Buzz"
	}

	return fmt.Sprintf("%d", n)
}

func RandomFizzBuzz(r *rand.Rand) string {
	return fmt.Sprintf("%s%s%s%s", FizzBuzz(r.Intn(8)+1), FizzBuzz(r.Intn(8)+1), FizzBuzz(r.Intn(8)+1), FizzBuzz(r.Intn(8)+1))
}
