package main

import (
	"testing"
)

func BenchmarkWC(b *testing.B) {
	for i := 0; i < b.N; i++ {
		wc()
	}
}