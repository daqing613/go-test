package main

import (
	"fmt"
	"testing"
)

func TestCalculate(t *testing.T) {
	fmt.Println("Test Calculate")
	expected := 4
	result := Calculate(2)
	if expected != result {
		t.Error("Failed")
	}
}

// func BenchmarkCalculate(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		Calculate(2)
// 	}
// }

func benchmarkCalculate(input int, b *testing.B) {
	for i := 0; i < b.N; i++ {
		Calculate(input)
	}
}

func BenchmarkCalculate100(b *testing.B)         { benchmarkCalculate(100, b) }
func BenchmarkCalculateNegative100(b *testing.B) { benchmarkCalculate(-100, b) }
func BenchmarkCalculateNegative1(b *testing.B)   { benchmarkCalculate(-1, b) }

func TestOther(t *testing.T) {
	fmt.Println("Testing something else")
	fmt.Println("This should't run with -run=calc")
}
