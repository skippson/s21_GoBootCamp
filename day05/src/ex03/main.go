package main

import (
	"day05/present"
	"fmt"
)

func grabPresents(slice []present.Present, capacity int) []present.Present {
	n := len(slice)
	table := make([][]int, n+1)
	for i := range table {
		table[i] = make([]int, capacity+1)
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= capacity; j++ {
			if slice[i-1].Size > j {
				table[i][j] = table[i-1][j]
			} else {
				table[i][j] = max(table[i-1][j], table[i-1][j-slice[i-1].Size]+slice[i-1].Value)
			}
		}
	}

	ans := make([]present.Present, 0)
	for i, j := len(slice), capacity; i > 0 && j > 0; {
		if table[i][j] != table[i-1][j] {
			ans = append(ans, slice[i-1])
			j -= slice[i-1].Size
		}
		i--	
	}

	return ans
}

func testSlice() []present.Present {
	root := []present.Present{{5, 1}, {4, 5}, {3, 5}, {5, 2}, {4, 1}, {4, 3}}

	return root
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	slice := testSlice()
	ans := grabPresents(slice, 8)
	fmt.Println(ans)
}
