package main

import "testing"

var (
	files = []string{"books/hard", "books/noon", "books/prisoners", "books/time"}
)

func BenchmarkReading0(b *testing.B) {
	for i := 0; i < b.N; i++ {
		readingLoad(files, 0)
	}
}

func BenchmarkReading100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		readingLoad(files, 100)
	}
}

func BenchmarkReading100000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		readingLoad(files, 100000)
	}
}
