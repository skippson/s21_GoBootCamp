package main

import (
	"day05/present"
	"fmt"
	"log"
)

func testSlice() []present.Present{
	root := []present.Present{{5, 1}, {4, 5}, {3, 1}, {5, 2}}

	return root
}

func getNCoolestPresents(slice []present.Present, n int) []present.Present{
	if n < 0 {
		log.Fatal("n is a negative number")
	} 

	if n > len(slice){
		log.Fatal("n out of range len(slice)")
	}

	heap := present.HeapInit(slice)
	ans := make([]present.Present, 0)
	for i := 0; i < n; i++{
		ans = append(ans, heap.Pop())
	}

	return ans
}

func main(){
	root := testSlice()
	ans := getNCoolestPresents(root, 2)

	fmt.Println(ans)
}