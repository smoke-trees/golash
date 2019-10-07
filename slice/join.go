package slice

import (
	"fmt"
	"reflect"
)

func Join(sl interface{}, s string) string {
	slice, success := takeArg(sl, reflect.Slice)
	if !success {
		panic("You must pass in a slice value!")
	}

	result := ""

	for i := 0; i < slice.Len(); i++ {
		if i != slice.Len() - 1 {
			result += fmt.Sprint(slice.Index(i).Interface()) + s
		} else {
			result += fmt.Sprint(slice.Index(i).Interface())
		}
	}

	return result
}
