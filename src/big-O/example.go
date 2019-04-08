package main

import (
	"fmt"
)

func main() {
	n := []int{1, 2, 3, 4, 5}
	m := []int{6, 7, 8, 9, 10}

	s1, r1 := getSum(n)
	fmt.Printf("Sum: %d / Reps: %d\n", s1, r1)

	s2, r2 := getSum2(n)
	fmt.Printf("Sum: %d / Reps: %d\n", s2, r2)

	s3, r3 := getSum3(n)
	fmt.Printf("Sum: %d / Reps: %d\n", s3, r3)

	s4, r4 := getSum4(n, m)
	fmt.Printf("Sum: %d / Reps: %d\n", s4, r4)
}

// Time complixity: O(N)
func getSum(n []int) (int, int) {
	sum := 0
	reps := 0
	cnt := len(n)
	for i := 0; i < cnt; i++ {
		sum += n[i]
		reps++
	}

	for j := 0; j < cnt; j++ {
		sum += n[j]
		reps++
	}

	return sum, reps
}

// Time complixity: O(N^2)
func getSum2(n []int) (int, int) {
	sum := 0
	reps := 0
	cnt := len(n)
	for i := 0; i < cnt; i++ {
		sum += n[i]
		reps++
		for j := 0; j < cnt; j++ {
			sum += n[j]
			reps++
		}
	}

	return sum, reps
}

// Time complixity: O(N^2)
func getSum3(n []int) (int, int) {
	sum := 0
	reps := 0
	cnt := len(n)
	for i := 0; i < cnt; i++ {
		sum += n[i]
		reps++
		for j := i + 1; j < cnt; j++ {
			sum += n[j]
			reps++
		}
	}

	return sum, reps
}

// Time complixity: O(nm)
func getSum4(n []int, m []int) (int, int) {
	sum := 0
	reps := 0
	aCnt := len(n)
	bCnt := len(m)
	for i := 0; i < aCnt; i++ {
		sum += n[i]
		reps++
		for j := 0; j < bCnt; j++ {
			sum += m[j]
			reps++
		}
	}

	return sum, reps
}
