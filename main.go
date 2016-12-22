package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	argStrings := os.Args[1:]
	numArgs := len(argStrings)
	argInts := make([]int, numArgs)
	for i := 0; i < numArgs; i++ {
		num, err := strconv.Atoi(argStrings[i])
		if err != nil {
			panic(err)
		}
		argInts[i] = num
	}
	sorted := mergesort(argInts)
	fmt.Println(sorted)
}

func mergesort(arr []int) []int {
	if len(arr) == 0 || len(arr) == 1 {
		return arr
	}
	halfway := len(arr) / 2
	return merge(mergesort(arr[:halfway]), mergesort(arr[halfway:]))
}

func merge(a []int, b []int) []int {
	result := make([]int, len(a)+len(b))
	i := 0
	j := 0
	for k := 0; k < len(a)+len(b); k++ {
		// used up all of a
		if i == len(a) {
			result = append(result[:k], b[j:]...)
			break
		}
		// used up all of b
		if j == len(b) {
			result = append(result[:k], a[i:]...)
			break
		}
		if a[i] < b[j] {
			result[k] = a[i]
			i++
		} else {
			result[k] = b[j]
			j++
		}
	}
	return result
}
