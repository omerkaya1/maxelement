package comparator

import (
	"fmt"
	"sort"
)

func CustomSliceSortWrapper(data interface{}) interface{} {
	switch slice := data.(type) {
	case []int:
		if !sort.IntsAreSorted(slice) {
			sort.Ints(slice)
		}
		return slice[len(slice)-1]
	case []float64:
		if !sort.Float64sAreSorted(slice) {
			sort.Float64s(slice)
		}
		return slice[len(slice)-1]
	case []string:
		if !sort.StringsAreSorted(slice) {
			sort.Strings(slice)
		}
		return slice[len(slice)-1]
	// TODO: implement me, please!
	case []struct {
		string
		int
	}:
		fmt.Printf("a slice of struct{string;int}\n")
		break
	// TODO: implement me, please!
	case []struct {
		int
		string
	}:
		fmt.Printf("a slice of struct{int;string}\n")
		break
	default:
		fmt.Printf("not a slice, or an unknown type\n")
		break
	}
	return nil
}
