package main

import (
	"fmt"
	"log"
	"math"
	"net/http"
	"strconv"
	"time"

	"git.wndv.co/workshop/letsgo/hello/fizzbuzz"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// result := sum(primes(100))
		result2 := mapAbs([]float64{-1, -2, -3})
		fmt.Fprint(w, result2)
	})

	r.HandleFunc("/fizzbuzz/{num}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		num := vars["num"]
		n, err := strconv.Atoi(num)
		if err != nil {
			panic(err)
		}

		fmt.Fprint(w, fizzbuzz.FizzBuzz(n))
	})

	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8080",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
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

func primes(a int) []int {
	var list []int
	for i := 1; i <= a; i++ {
		count := 0
		for j := 1; j <= i; j++ {
			if i%j == 0 {
				count++
			}
		}
		if count == 2 {
			list = append(list, i)
		}
	}

	return list
}

func power(base, exponent int) int {
	result := 1
	for i := 1; i <= exponent; i++ {
		result = result * base
	}
	return result
}

func sum(nums []int) int {
	result := 0
	for _, v := range nums {
		result += v
	}
	return result
}

func mapAbs(nums []float64) []float64 {
	var arr []float64
	for _, v := range nums {
		arr = append(arr, math.Abs(v))
	}

	return arr
}
