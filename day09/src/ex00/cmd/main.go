package main

import (
	. "day09/ex00/sleep_sort"
	"fmt"
)

func main(){
	slice := []int{5,4,3,2,1}
	ch := SleepSort(slice)
	for val := range ch{
		fmt.Print(val)
	}
}