package require_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zimmermanncode/go-require"
)

func TestNilPtr(t *testing.T) {
	t.Parallel()
	t.Run("should accept nil pointer & return nil", func(t *testing.T) {
		t.Parallel()
		assert.Nil(t, require.NilPtr("Test pointer", (*int)(nil)))
	})
	t.Run("should panic when pointer is not nil", func(t *testing.T) {
		value := "test"

		t.Parallel()
		assert.PanicsWithValue(t, "assertion failed: Test pointer should be a nil pointer", func() {
			require.NilPtr("Test pointer", &value)
		})
	})
	t.Run("should use given name in panic message", func(t *testing.T) {
		value := 42

		t.Parallel()
		assert.PanicsWithValue(t, "assertion failed: Other value should be a nil pointer", func() {
			require.NilPtr("Other value", &value)
		})
	})
}

func TestNotNilPtr(t *testing.T) {
	t.Parallel()
	t.Run("should accept not-nil pointer & return same pointer", func(t *testing.T) {
		value := 42

		t.Parallel()
		assert.Same(t, &value, require.NotNilPtr("Test pointer", &value))
	})
	t.Run("should panic when pointer is nil", func(t *testing.T) {
		t.Parallel()
		assert.PanicsWithValue(t, "assertion failed: Test pointer should not be a nil pointer", func() {
			require.NotNilPtr("Test pointer", (*string)(nil))
		})
	})
	t.Run("should use given name in panic message", func(t *testing.T) {
		t.Parallel()
		assert.PanicsWithValue(t, "assertion failed: Other value should not be a nil pointer", func() {
			require.NotNilPtr("Other value", (*float64)(nil))
		})
	})
}

func BenchmarkNilPtr(b *testing.B) {
	b.ResetTimer()

	for range b.N {
		require.NilPtr("Benchmark pointer", (*int)(nil))
	}
}

func BenchmarkNotNilPtr(b *testing.B) {
	value := "benchmark"

	b.ResetTimer()

	for range b.N {
		require.NotNilPtr("Benchmark pointer", &value)
	}
}
