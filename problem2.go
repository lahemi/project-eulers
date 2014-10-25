package main

/*
Even Fibonacci numbers

By considering the terms in the Fibonacci sequence whose values do not
exceed four million, find the sum of the even-valued terms.
*/
func problem2() int {
	var fib func(int, int, int) int
	fib = func(prev, cur, sum int) int {
		if cur < 4000000 {
			if cur%2 == 0 {
				sum += cur
			}
			return fib(cur, prev+cur, sum)
		}
		return sum
	}
	return fib(1, 2, 0)
}
