package main

import (
	"os"
	"testing"
)

func BenchmarkMain(b *testing.B) {
originalStdout := os.Stdout
os.Stdout,_ = os.Open(os.DevNull)

for i := 0; i < b.N; i++ {
	main()
}

os.Stdout = originalStdout
}
