package bufs

import (
	"bytes"
	"strings"
	"testing"
	"unsafe"
)

func BenchmarkBuffer(b *testing.B) {
	buf := bytes.Buffer{}
	for range 10_000 {
		buf.WriteByte('a')
	}
	_ = buf.String()
}

func BenchmarkString(b *testing.B) {
	buf := strings.Builder{}
	for range 10_000 {
		buf.WriteByte('a')
	}
	_ = buf.String()
}

func BenchmarkArray(b *testing.B) {
	buf := make([]byte, 0, 16)
	for range 10_000 {
		buf = append(buf, 'a')
	}
	_ = unsafe.String(unsafe.SliceData(buf), len(buf))
}

func BenchmarkBufferLarge(b *testing.B) {
	buf := bytes.Buffer{}
	for range 1_000_000 {
		buf.WriteByte('a')
	}
	_ = buf.String()
}

func BenchmarkStringLarge(b *testing.B) {
	buf := strings.Builder{}
	for range 1_000_000 {
		buf.WriteByte('a')
	}
	_ = buf.String()
}

func BenchmarkArrayLarge(b *testing.B) {
	buf := make([]byte, 0, 16)
	for range 1_000_000 {
		buf = append(buf, 'a')
	}
	_ = unsafe.String(unsafe.SliceData(buf), len(buf))
}
