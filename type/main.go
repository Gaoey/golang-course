package main

import (
	"fmt"
	"strconv"
)

type Int int

func (i Int) String() string {
	return strconv.Itoa(int(i))
}

func (i *Int) Set(n int) {
	*i = Int(n) // refer type
}

func main() {
	var n Int = 14
	fmt.Println(n.String())
	n.Set(15)
	fmt.Println(n.String())
}
