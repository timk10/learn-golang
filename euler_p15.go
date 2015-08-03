package main

import (
	"fmt"
)
/*
Starting in the top left corner of a 2×2 grid, and only being able to move to the right and down, there are exactly
6 routes to the bottom right corner.


How many such routes are there through a 20×20 grid?
 */

func main() {
	const len = 21
	lattice := [len][len]uint64{};

	// init the matrix
	for i := 0; i < len; i++ {
		lattice[i][0] = uint64(1)
		lattice[0][i] = uint64(1)
	}

	for i := 1; i < len; i++ {
		for j := 1; j < len; j++ {
			lattice[i][j] = lattice[i-1][j] + lattice[i][j-1]
		}
	}

	fmt.Printf("%d\n", lattice[20][20])
}
