package benchtest

import (
	mincoins "day07/ex01/mincoins"
	"testing"
)

func BenchmarkMinCoins2(b *testing.B) {
	val := 100
	coins := []int{1, 5, 10}
	for i := 0; i < b.N; i++ {
		mincoins.MinCoins2(val, coins)
	}
}

func BenchmarkMinCoins2Optimized(b *testing.B) {
	val := 100
	coins := []int{1, 5, 10}
	for i := 0; i < b.N; i++ {
		mincoins.MinCoins2Optimized(val, coins)
	}
}
