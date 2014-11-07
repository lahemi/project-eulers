package main

import "sort"

/*
 * Special Pythagorean triplet
 *
 * A Pythagorean triplet is a set of three natural numbers, a < b < c, for
 * which, a^2 + b^2 = c^2
 *
 * For example, 3^2 + 4^2 = 9 + 16 = 25 = 5^2.
 *
 * There exists exactly one Pythagorean triplet for which a + b + c = 1000.
 * Find the product abc.

Dickson's method

a² + b² = c²
a < b < c

a = r + s
b = r + t
c = r + s + t

c must be odd
r must be even
*/
func problem9() int {
	// Perhaps a bit naive implementation.
	factor := func(n int) (ret []int) {
		for i := 1; i <= n; i++ {
			if n%i == 0 {
				ret = append(ret, i)
			}
		}
		return
	}
	type Triplet struct {
		a, b, c int
	}
	prodTriplet := func(r, s, t int) Triplet {
		return Triplet{r + s, r + t, r + s + t}
	}
	genAllTriplets := func(r int, sl, tl []int) (ret []Triplet) {
		for i := 0; i < len(tl); i++ {
			tr := prodTriplet(r, sl[i], tl[i])
			if tr.c%2 == 1 && (tr.a < tr.b && tr.a < tr.c) && (tr.b < tr.c) {
				ret = append(ret, tr)
			}
		}
		return
	}
	// Because Dickson's method!
	dickson := func(r int) int { return (r * r) / 2 }

	for r := 2; ; r += 2 {
		fl := factor(dickson(r))
		// Grouping the smallest and largest values to be used together.
		ss, ts := fl[:len(fl)/2], fl[len(fl)/2:]
		sort.Sort(sort.Reverse(sort.IntSlice(ts)))
		all := genAllTriplets(r, ss, ts)
		for i := 0; i < len(all); i++ {
			tr := all[i]
			if tr.a+tr.b+tr.c == 1000 {
				return tr.a * tr.b * tr.c
			}
		}
	}
	// No value found...
	return 0
}
