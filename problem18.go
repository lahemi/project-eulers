package main

import (
	"strconv"
	"strings"
)

/*
 * Maximum path sum I
 *
 * By starting at the top of the triangle below and moving to adjacent numbers
 * on the row below, the maximum total from top to bottom is 23.
 *
 *      3
 *     7 4
 *    2 4 6
 *   8 5 9 3
 *
 * That is, 3 + 7 + 4 + 9 = 23.
 *
 * Find the maximum total from top to bottom of the given triangle with 15
 * rows:
 */
func problem18() int {
	max := func(x, y int) int {
		if x < y {
			return y
		}
		return x
	}

	var data = [][]int{}

	// input18 can be found in inputdata.go, for convenience.
	snl := strings.Split(input18, "\n")

	// Convert the data and fill the [][]int
	for i := 1; i < len(snl)-1; i++ {
		sints := strings.Split(snl[i], " ")
		var nums []int
		for _, n := range sints {
			cn, _ := strconv.Atoi(n) // There is nothing but ints in data, ignore err.
			nums = append(nums, cn)
		}
		data = append(data, nums)
	}

	// The actual logic and its implementation are pretty simple.
	// We "cheat" a little by not running down to the second last
	// line first, instead starting from there straight away.
	// Effectively the same.
	for line := len(data) - 2; line >= 0; line-- {
		for pos := 0; pos < len(data[line]); pos++ {
			e1, e2 := data[line+1][pos], data[line+1][pos+1]
			cur := data[line][pos]
			// I think the compiler actually inlines max to avoid the func call overhead.
			cur += max(e1, e2)
			data[line][pos] = cur
		}
	}

	return data[0][0]
}
