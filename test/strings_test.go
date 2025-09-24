package test

import (
	"strings"
	"testing"
)

func BenchmarkStringPlus(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := ""
		for i := 0; i < 10; i++ {
			s += "abc"
		}
	}
}

func BenchmarkStringBuilder(b *testing.B) {

	for i := 0; i < b.N; i++ {
		build := strings.Builder{}
		for i := 0; i < 10; i++ {
			build.WriteString("abc")
		}
	}
}
