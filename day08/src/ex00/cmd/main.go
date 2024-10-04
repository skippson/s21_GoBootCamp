package main

import (
	gl "day08/ex00/getelement"
	"fmt"
)

func main() {
	tests := []struct {
		name string
		idx  int
		arr  []int
	}{
		{"1", 4, []int{1, 2, 3, 4, 5, 6, 7, 8}},
		{"2", 7, []int{1, 2, 3, 4, 5, 6, 7, 8}},
		{"3", 2, []int{1, 2, 3, 4, 5, 6, 7, 8}},
		{"nil", 4, nil},
		{"empty slice", 1, []int{}},
		{"negative index", -1, []int{1, 2, 3, 4}},
		{"index out of range", 30, []int{1, 2, 3, 4, 5}},
	}

	for _, test := range tests {
		ans, err := gl.GetElement(test.arr, test.idx)
		fmt.Printf("case %s: arr = %v, idx = %d | getElement(ans, idx) = %d, %v\n", test.name, test.arr, test.idx, ans, err)
	}
}
