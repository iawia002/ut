package ut

import (
	"reflect"
)

// Range generate a sequence of numbers by range.
func Range(min, max int) []int {
	items := make([]int, max-min+1)
	for index := range items {
		items[index] = min + index
	}
	return items
}

// ItemInSlice if a item is in the list.
func ItemInSlice(item, list interface{}) bool {
	v1 := reflect.ValueOf(item)
	v2 := reflect.ValueOf(list)
	for i := 0; i < v2.Len(); i++ {
		indexType := v2.Index(i).Type().String()
		if v1.Type().String() != indexType {
			continue
		}
		switch indexType {
		case "int":
			if v1.Int() == v2.Index(i).Int() {
				return true
			}
		case "string":
			if v1.String() == v2.Index(i).String() {
				return true
			}
		}
	}
	return false
}
