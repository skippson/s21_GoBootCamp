package main

import (
	"day09/ex02/octopus"
	"fmt"
	"math"
	"sync"
)

type testing struct {
	str    string
	status bool
}

func mkCh(data any) <-chan interface{} {
	ch := make(chan interface{})

	go func() {
		for i := 0; i < 3; i++ {
			ch <- data
		}
		close(ch)
	}()

	return ch
}

func test1() {
	fmt.Println("\n\033[32mTEST 1\033[0m")

	ch1 := mkCh("daryl dixon")
	ch2 := mkCh("school 21")
	ch3 := mkCh(1337)
	ch4 := mkCh(2.28)
	ch5 := mkCh('E')
	ch6 := mkCh('4')
	t := testing{"carl grimes", true}
	ch7 := mkCh(t)

	out := octopus.Multiplaxe(ch1, ch2, ch3, ch4, ch5, ch6, ch7)
	for val := range out {
		fmt.Println(val)
	}
}

func test2() {
	fmt.Println("\n\033[32mTEST 2\033[0m")

	pow := []float64{1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0, 9, 0, 10.0}
	sqrt := []float64{322,72,32,41,69}
	ch1 := make(chan any)
	ch2 := make(chan any)

	wg := sync.WaitGroup{}
	wg.Add(len(pow) + len(sqrt))

	for _, val := range pow {
		go func() {
			defer wg.Done()
			ch1 <- math.Pow(val, 2.0)
		}()
	}

	for _, val := range sqrt {
		go func() {
			defer wg.Done()
			ch2 <- math.Sqrt(val)
		}()
	}

	go func() {
		wg.Wait()
		close(ch1)
		close(ch2)
	}()

	out := octopus.Multiplaxe(ch1, ch2)
	for val := range out {
		fmt.Println(val)
	}
}

func main() {
	test1()
	test2()
}
