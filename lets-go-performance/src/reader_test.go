package bufs

import (
	"bytes"
	"strings"
	"testing"
)

func BenchmarkReadByte(b *testing.B) {
	r := bytes.NewReader([]byte(strings.Repeat("1234", 50000)))
	b.ResetTimer()
	for range 4 * 50000 {
		r.ReadByte()
	}
}

func BenchmarkReadRune(b *testing.B) {
	r := bytes.NewReader([]byte(strings.Repeat("1234", 50000)))
	b.ResetTimer()
	for range 4 * 50000 {
		r.ReadRune()
	}
}
