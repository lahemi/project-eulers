package main

// Find the sum of all the multiples of 3 or 5 below 1000.
func problem1() (sum int) {
	limit := 1000
	for i := 1; i < limit; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum += i
		}
	}
	return
}
