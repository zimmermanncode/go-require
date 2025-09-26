package require

import (
	"testing"

	assert "github.com/stretchr/testify/assert"
)

func TestNilPtr(t *testing.T) {
	t.Run("should accept nil pointer & return nil", func(t *testing.T) {
		var nilPtr *int
		assert.Nil(t, NilPtr("Test pointer", nilPtr))
	})
	t.Run("should panic when pointer is not nil", func(t *testing.T) {
		assert.PanicsWithValue(t, "assertion failed: Test pointer should be a nil pointer", func() {
			value := "test"
			NilPtr("Test pointer", &value)
		})
	})
	t.Run("should use given name in panic message", func(t *testing.T) {
		assert.PanicsWithValue(t, "assertion failed: Other pointer should be a nil pointer", func() {
			value := 42
			NilPtr("Other pointer", &value)
		})
	})
}

func TestNotNilPtr(t *testing.T) {
	t.Run("should accept not-nil pointer & return same pointer", func(t *testing.T) {
		value := 42
		assert.Same(t, &value, NotNilPtr("Test pointer", &value))
	})
	t.Run("should panic when pointer is nil", func(t *testing.T) {
		assert.PanicsWithValue(t, "assertion failed: Test pointer should not be a nil pointer", func() {
			var nilPtr *string
			NotNilPtr("Test pointer", nilPtr)
		})
	})
	t.Run("should use given name in panic message", func(t *testing.T) {
		assert.PanicsWithValue(t, "assertion failed: Other pointer should not be a nil pointer", func() {
			var nilPtr *float64
			NotNilPtr("Other pointer", nilPtr)
		})
	})
}

func BenchmarkNilPtr(b *testing.B) {
	var nilPtr *int
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		NilPtr("Benchmark pointer", nilPtr)
	}
}

func BenchmarkNotNilPtr(b *testing.B) {
	value := "benchmark"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		NotNilPtr("Benchmark pointer", &value)
	}
}
