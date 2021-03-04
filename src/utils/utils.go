package utils

import (
	"fmt"
	"strconv"
)

func IntSliceToString(arr []int) string {
	var s string = "["
	for i, v := range arr {
		if i != 0 {
			s += ","
		}
		s += strconv.Itoa(v)
	}
	s += "]"
	return s
}

func StringSliceToString(arr []string) string {
	var s string = "["
	for i, v := range arr {
		if i != 0 {
			s += ","
		}
		s += fmt.Sprintf("\"%s\"", v)
	}
	s += "]"
	return s
}

func FloatSliceToString(arr []float64) string {
	var s string = "["
	for i, v := range arr {
		if i != 0 {
			s += ","
		}
		s += strconv.FormatFloat(v, 'e', 8, 64)
	}
	s += "]"
	return s
}
