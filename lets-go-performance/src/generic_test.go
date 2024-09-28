package bufs

import (
	"testing"
)

func generic[T any](data any) (T, bool) {
	v, ok := data.(T)
	if !ok {
		var e T
		return e, false
	}
	return v, true
}

func specific(data any) (bool, bool) {
	switch data.(type) {
	case bool:
		return true, true
	default:
		return false, false
	}
}

func BenchmarkGeneric(b *testing.B) {
	var e any
	e = false
	for range 1000000 {
		_, _ = generic[bool](e)
	}
}

func BenchmarkSpecific(b *testing.B) {
	var e any
	e = false
	for range 1000000 {
		_, _ = specific(e)
	}
}
