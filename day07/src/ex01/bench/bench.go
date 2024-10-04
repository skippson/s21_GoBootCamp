package main

import (
	"log"
	"os"
	"runtime/pprof"
	"day07/ex01/mincoins"
)


func main() {
    f, err := os.Create("cpu.prof")
    if err != nil {
        log.Fatal(err)
    }
    pprof.StartCPUProfile(f)
    defer pprof.StopCPUProfile()
    for i := 0; i < 1000; i++ {
        mincoins.MinCoins2(100, []int{1, 5, 10, 25})
    }
}