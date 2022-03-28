package my_test

import (
	"fmt"
	"testing"
)

func TestSomething(t *testing.T) {
	fmt.Println("1. write some test")
}

func TestFail(t *testing.T) {
	fmt.Println("2. test fail")
	t.Fail()
}

func TestError(t *testing.T) {
	fmt.Println("3. test error")
	t.Error("something wrong")
}

//計算 A+B
func Add(a, b int) int {
	return a + b
}
func TestAdd(t *testing.T) {
	fmt.Println("4. test add")
	if Add(1, 2) == 3 {
		t.Log("my_test.Add PASS")
	} else {
		t.Error("my_test.Add FAIL")
	}
}

func TestSkip1(t *testing.T) {
	fmt.Println("5. test skip pass")
	t.Skip()
	fmt.Println("5. test skip fail")
}

func BenchmarkAdd(b *testing.B) {
	fmt.Println("6. test Benchmark")
	for i := 0; i < b.N; i++ {
		Add(1, 2)
	}
}
