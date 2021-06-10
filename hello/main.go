package main

func main() {
	result := squareArea(2)
	println("result: ", result)
	println(multiple(1, 2))
	println(isEven(2))
}

func squareArea(a int) int {
	return a * a
}

func multiple(a, b int) (int, int) {
	return b, a + b
}

func isEven(a int) bool {
	return a%2 == 0
}
