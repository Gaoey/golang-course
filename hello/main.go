package main

import (
	"fmt"
	"strconv"
)

func main() {
	result := squareArea(2)
	println("result: ", result)
	println(multiple(1, 2))

	v, err := isEven("aa")
	if err != nil {
		fmt.Printf("err %s", err.Error())
	}

	fmt.Printf("isEven result: %v", v)

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
