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
func problem67() int {
	max := func(x, y int) int {
		if x < y {
			return y
		}
		return x
	}
	type Data []int
	var data = []Data{}
	snl := strings.Split(input67, "\n")
	for i := 1; i < len(snl)-1; i++ {
		sints := strings.Split(snl[i], " ")
		var nums []int
		for _, n := range sints {
			cn, _ := strconv.Atoi(n) // There is nothing but ints in data, ignore err.
			nums = append(nums, cn)
		}
		data = append(data, nums)
	}
	for line := len(data) - 2; line >= 0; line-- {
		for pos := 0; pos < len(data[line]); pos++ {
			e1, e2 := data[line+1][pos], data[line+1][pos+1]
			cur := data[line][pos]
			cur += max(e1, e2)
			data[line][pos] = cur
		}
	}

	return data[0][0]
}
