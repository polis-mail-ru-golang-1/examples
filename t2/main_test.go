package main

import "testing"

var (
	files = []string{"books/hard", "books/noon", "books/prisoners", "books/time"}
)

func BenchmarkParallel(b *testing.B) {
	for i := 0; i < b.N; i++ {
		parallelLoad(files)
	}
}
