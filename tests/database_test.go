package tests

import (
	"testing"

	"github.com/tof4/yrc/server"
)

func BenchmarkBackwardFileReader10(b *testing.B) {
	server.BackwardFileRead("test.txt", 10)
}

func BenchmarkBackwardFileReader100(b *testing.B) {
	server.BackwardFileRead("test.txt", 100)
}

func BenchmarkBackwardFileReader1000(b *testing.B) {
	server.BackwardFileRead("test.txt", 1000)
}

func BenchmarkBackwardFileReader10000(b *testing.B) {
	server.BackwardFileRead("test.txt", 10000)
}
