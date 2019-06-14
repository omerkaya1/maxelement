package main

import (
	"fmt"
	"github.com/omerkaya1/maxelement/comparator"
)

var (
	testCase1 = []int{4, 7, 1, 2, 5, 6}
	testCase4 = []float64{1, 2, 3, 4, 5, 6}
	testCase5 = []string{"r", "jdf", "skd", "klslke", "sfrws", "wfr", "a"}
	testCase6 = []struct {
		string
		int
	}{{"r", 1}, {"jdf", 2}, {"skd", 3}, {"klslke", 4}, {"sfrws", 7}, {"wfr", 6}}
	testCase7 = []struct {
		int
		string
	}{{1, "ters"}, {2, "sfjh"}}
)

func main() {
	//test := comparator.GetUnknownSlice()
	//test.CustomSliceSortWrapper(testCase1)
	//test.CustomSliceSortWrapper(testCase2)
	//test.CustomSliceSortWrapper(testCase3)
	//test.CustomSliceSortWrapper(testCase4)
	//test.CustomSliceSortWrapper(testCase5)
	//test.CustomSliceSortWrapper(testCase6)
	//test.CustomSliceSortWrapper(testCase7)
	//test.CustomSliceSortWrapper(453)
	try := comparator.CustomSliceSortWrapper(testCase5)
	fmt.Printf("try: %v\n", try)
	try1 := comparator.CustomSliceSortWrapper(testCase1)
	fmt.Printf("try: %v\n", try1)
	try2 := comparator.CustomSliceSortWrapper(testCase4)
	fmt.Printf("try: %v\n", try2)
	//comparator.CustomSliceSortWrapper(testCase2)
	//comparator.CustomSliceSortWrapper(testCase3)
	//comparator.CustomSliceSortWrapper(testCase4)
	//comparator.CustomSliceSortWrapper(testCase5)
	//comparator.CustomSliceSortWrapper(testCase6)
	//comparator.CustomSliceSortWrapper(testCase7)
}
