package main

import "testing"

// 3.46 ns/op
func BenchmarkSliceAndReservingMemory(b *testing.B) {
	slice := make([]int, 0, b.N)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		slice = append(slice, i*3)
	}
}

// 39.3 ns/op
func BenchmarkSliceWithoutReservingMemory(b *testing.B) {
	var slice []int
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		slice = append(slice, i*3)
	}
}

const SIZE = 10

// 427.588 ns/op
func BenchmarkArrayMultiple(b *testing.B) {
	for i := 0; i < b.N; i++ {

		var array [SIZE]int
		for j := 0; j < SIZE; j++ {
			array[j] = j * 3
		}

	}
}

// 3.137.440 ns/op
func BenchmarkSliceAndReservingMemoryMultiple(b *testing.B) {
	for i := 0; i < b.N; i++ {
		slice := make([]int, 0, SIZE)
		for j := 0; j < SIZE; j++ {
			slice = append(slice, j*3)
		}

	}
}

// 13.286.325 ns/op
func BenchmarkSliceWithoutReservingMemoryMultiple(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var slice []int
		for j := 0; j < SIZE; j++ {
			slice = append(slice, j*3)
		}
	}
}
