package getelementtest

import (
	"day08/ex00/getelement"
	"errors"
	"testing"
)

func TestGetElement(t *testing.T) {
	tests := []struct {
		name string
		idx  int
		arr  []int
		ans  int
		err  error
	}{
		{"1", 4, []int{1, 2, 3, 4, 5, 6, 7, 8}, 5, nil},
		{"2", 7, []int{1, 2, 3, 4, 5, 6, 7, 8}, 8, nil},
		{"3", 2, []int{1, 2, 3, 4, 5, 6, 7, 8}, 3, nil},
		{"nil", 4, nil, 0, errors.New("non-existent slice")},
		{"empty slice", 1, []int{}, 0, errors.New("empty slice")},
		{"negative index", -1, []int{1, 2, 3, 4}, 0, errors.New("negative index")},
		{"index out of range", 30, []int{1, 2, 3, 4, 5}, 0, errors.New("index out of range")},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ans, err := getelement.GetElement(test.arr, test.idx)
			if ans != test.ans && !errors.Is(err, test.err){
				t.Errorf("val, err = %d, %v.  want %d, %v", ans, err, test.ans, test.err)
			}
		})
	}
}
