package main

import "fmt"

type Tfun struct {
	name string
	fn   func() int
	ret  int
}

func resetColours()  { fmt.Print("\033[0m") }
func passedColours() { fmt.Print("\033[32m") }
func failedColours() { fmt.Print("\033[31;1m") }

var problems = []Tfun{
	Tfun{"problem1", problem1, 233168},
	Tfun{"problem2", problem2, 4613732},
	//Tfun{"problem3", problem3, 6857},
	Tfun{"problem4", problem4, 906609},
	//Tfun{"problem5", problem5, 232792560},
	//Tfun{"problem6", problem6, 25164150},
	//Tfun{"problem7", problem7, 104743},
	//Tfun{"problem8", problem8, 40824},
	Tfun{"problem9", problem9, 31875000},
	//Tfun{"problem10", problem10, 142913828922},
	//Tfun{"problem11", problem11, 70600674},
	//Tfun{"problem12", problem12, 76576500},
	//Tfun{"problem13", problem13, 5537376230},
	//Tfun{"problem14", problem14, 837799},
	//Tfun{"problem15", problem15, 137846528820},
	//Tfun{"problem16", problem16, 1366},
	//Tfun{"problem17", problem17, 21124},
	Tfun{"problem18", problem18, 1074},
	//Tfun{"problem19", problem19, 171},
	//Tfun{"problem20", problem20, 648},
	//Tfun{"problem21", problem21, 31626},
	//Tfun{"problem22", problem22, 871198282},
	//Tfun{"problem23", problem23, 4179871},
	//Tfun{"problem24", problem24, 2783915460},
	//Tfun{"problem25", problem25, 4782},
	//Tfun{"problem26", problem26, 983},
	//Tfun{"problem27", problem27, -59231},
	//Tfun{"problem28", problem28, 669171001},
	//Tfun{"problem29", problem29, 9183},
	//Tfun{"problem30", problem30, 443839},
	//Tfun{"problem31", problem31, 73682},
	//Tfun{"problem32", problem32, 45228},
	//Tfun{"problem33", problem33, 100},
	//Tfun{"problem34", problem34, 40730},
	//Tfun{"problem35", problem35, 55},
	//Tfun{"problem36", problem36, 872187},
	//Tfun{"problem37", problem37, 748317},
	//Tfun{"problem38", problem38, 932718654},
	//Tfun{"problem39", problem39, 840},
	//Tfun{"problem40", problem40, 210},
	//Tfun{"problem41", problem41, 7652413},
	//Tfun{"problem42", problem42, 162},
	//Tfun{"problem43", problem43, 16695334890},
	//Tfun{"problem44", problem44, 5482660},
	//Tfun{"problem45", problem45, 1533776805},
	//Tfun{"problem46", problem46, 5777},
	//Tfun{"problem47", problem47, 134043},
	//Tfun{"problem48", problem48, 9110846700},
	//Tfun{"problem49", problem49, 296962999629},
	//Tfun{"problem50", problem50, 997651},
	Tfun{"problem67", problem67, 7273},
}

func main() {
	for _, t := range problems {
		switch out := t.fn(); out {
		case t.ret:
			passedColours()
			fmt.Printf("Passed %s with %d\n", t.name, t.ret)
		default:
			failedColours()
			fmt.Printf("%s → %d, but expected %d\n", t.name, out, t.ret)
		}
        resetColours()
	}
}
