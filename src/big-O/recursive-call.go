package main

import (
	"fmt"
)

func main() {
	maxDepth := 4
	for depth := 0; depth <= maxDepth; depth++ {
		fmt.Printf("Depth level: %d / Number of nodes: %d\n", depth, f(depth))
	}
}

// Time complexity: O(2^N)
// Space complexity: O(N)
func f(n int) int {
	if n <= 1 {
		return 1
	}

	return f(n-1) + f(n-1)
}
