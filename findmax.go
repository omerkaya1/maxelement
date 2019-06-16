package maxelement

import (
	"reflect"
)

type Person struct {
	Name string
	Age  int
}

type PersonByAge []Person
type PersonByName []Person

/**
 * Use FindMaxWrapper() to deal with basic types like int, int32, int64, float32, float64, string etc.
 * It should not be used while dealing with complex or composite data types, such as structs or maps
 * Use FindMaxReflection() function instead.
 *
 * Returns the max element in the data interface.
 * If the data type is unknown, returns nil.
 */
func FindMaxWrapper(data interface{}) interface{} {
	switch slice := data.(type) {
	case []int:
		return FindMaxReflection(slice, func(i, j int) bool { return slice[i] > slice[j] })
	case []int32:
		return FindMaxReflection(slice, func(i, j int) bool { return slice[i] > slice[j] })
	case []int64:
		return FindMaxReflection(slice, func(i, j int) bool { return slice[i] > slice[j] })
	case []float32:
		return FindMaxReflection(slice, func(i, j int) bool { return slice[i] > slice[j] })
	case []float64:
		return FindMaxReflection(slice, func(i, j int) bool { return slice[i] > slice[j] })
	case []string:
		return FindMaxReflection(slice, func(i, j int) bool { return slice[i] > slice[j] })
	case PersonByAge:
		return FindMaxReflection(slice, func(i, j int) bool { return slice[i].Age > slice[j].Age })
	case PersonByName:
		return FindMaxReflection(slice, func(i, j int) bool { return slice[i].Name > slice[j].Name })
	default:
		return nil
	}
}

/**
 * Use FindMaxReflection() for more complex or composite data types.
 * Returns the max element in the data interface or nil, if things go south.
 */
func FindMaxReflection(data interface{}, comparator func(i, j int) bool) interface{} {
	rt := reflect.ValueOf(data)
	if rt.Kind() != reflect.Slice || rt.Len() == 0 {
		return nil
	}
	// Convert reflected value to a slice of interfaces
	interfaceSlice := make([]interface{}, rt.Len())
	for i := 0; i < rt.Len(); i++ {
		interfaceSlice[i] = rt.Index(i).Interface()
	}
	// Leverage the use of a custom comparator to define the max element in the slice
	maxIndex := 0
	for i := range interfaceSlice {
		if comparator(i, maxIndex) {
			maxIndex = i
		}
	}
	return interfaceSlice[maxIndex]
}
