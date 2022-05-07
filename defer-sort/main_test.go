package main

import (
	"testing"
	"time"
)

func BenchmarkReverseDefer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		reverseDefer(arr)
		reverseSimple(arr)
	}
}

func BenchmarkReverseSimple(b *testing.B) {
	for i := 0; i < b.N; i++ {
		reverseSimple(arr)
	}
}

func BenchmarkReverseSimpleSlow(b *testing.B) {
	for i := 0; i < b.N; i++ {
		time.Sleep(200 * time.Nanosecond)
		reverseSimple(arr)
	}
}
