package my_benchmark

import (
	"fmt"
	"testing"
)

//計算 A+B
func Add(a, b int) int {
	return a + b
}

func BenchmarkAdd(b *testing.B) {
	fmt.Println("6. test Benchmark")
	for i := 0; i < b.N; i++ {
		Add(1, 2)
	}
}
