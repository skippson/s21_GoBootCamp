package mincoinstest

import (
	 "day07/ex00/mincoins"
	. "day07/ex00/mincoins_test/sliceEqual"
	"testing"
)

func TestMinCoins(t *testing.T) {
	tests := []struct {
		name     string
		val      int
		coins    []int
		expected []int
	}{
		{"simple", 13, []int{1, 5, 10}, []int{1, 1, 1, 10}},
		{"nil coins", 13, nil, nil},
		{"edge coins", 13, []int{}, []int{}},
		{"val = 0", 0, []int{1, 5, 10}, []int{}},
		{"duplicate coins", 13, []int{1, 5, 5, 10}, []int{1, 1, 1, 10}},
		{"unsorted coins", 13, []int{10, 5, 1}, []int{10, 1, 1, 1}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := mincoins.MinCoins(test.val, test.coins)
			if !SliceEqual(actual, test.expected) {
				t.Errorf("minCoins(%d, %v) = %v, want %v", test.val, test.coins, actual, test.expected)
			}
		})
	}
}
