package maxelement

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	testCase0 = []int{}
	testCase1 = []int{4, 7, 1, 2, 5, 6}
	testCase2 = []int32{4, 7, 1, 2, 5, 6}
	testCase3 = []int64{4, 7, 1, 2, 5, 6}
	testCase4 = []float32{1, 2, 3, 4, 5, 6}
	testCase5 = []float64{1, 2, 3, 4, 5, 6}
	testCase6 = []string{"r", "jdf", "skd", "klslke", "sfrws", "wfr", "a", "z"}
	testCase7 = []Person{
		{"Rob Pike", 1},
		{"Ken Thompson", 2},
		{"Brian Kernighan", 3},
		{"Dennis Ritchie", 4},
		{"Alan Turing", 11},
		{"Isaac Newton", 13},
		{"Doug McIlroy", 3},
		{"Bjarne Stroustrup", 2}}
	testCase8 = []struct {
		int
		string
	}{{1, "ters"}, {2, "sfjh"}}
)

func TestFindMaxWrapper(t *testing.T) {
	assert.Equal(t, nil, FindMaxWrapper(testCase0))
	assert.Equal(t, 7, FindMaxWrapper(testCase1))
	assert.Equal(t, int32(7), FindMaxWrapper(testCase2))
	assert.Equal(t, int64(7), FindMaxWrapper(testCase3))
	assert.Equal(t, float32(6), FindMaxWrapper(testCase4))
	assert.Equal(t, float64(6), FindMaxWrapper(testCase5))
	assert.Equal(t, "z", FindMaxWrapper(testCase6))
	assert.Equal(t, Person{"Isaac Newton", 13}, FindMaxWrapper(PersonByAge(testCase7)))
	assert.Equal(t, nil, FindMaxWrapper(testCase8))
}

func TestFindMaxReflection(t *testing.T) {
	assert.Equal(t, Person{"Rob Pike", 1}, FindMaxReflection(testCase7, func(i, j int) bool { return testCase7[i].Name > testCase7[j].Name }))
	assert.Equal(t, Person{"Isaac Newton", 13}, FindMaxReflection(testCase7, func(i, j int) bool { return testCase7[i].Age > testCase7[j].Age }))
}
