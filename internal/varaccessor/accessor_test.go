package varaccessor

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVarAccessor(t *testing.T) {
	t.Run("simple", func(t *testing.T) {
		acc1 := New(map[string]any{"a": 1, "b": 2, "c": 3})
		assert.Equal(t, 1, acc1.Int("a"))
		assert.Equal(t, 2, acc1.Int("b"))
		assert.Equal(t, 3, acc1.Int("c"))
		assert.Equal(t, "def", acc1.String("d", "def"))
	})

	t.Run("nested", func(t *testing.T) {
		acc1 := New(map[string]any{"a": 1, "b": 2, "c": map[string]any{"d": 4, "e": 5}})
		assert.Equal(t, 1, acc1.Int("a"))
		assert.Equal(t, 2, acc1.Int("b"))
		assert.Equal(t, 4, acc1.Int("c.d"))
		assert.Equal(t, "5", acc1.String("c.e"))
	})

	t.Run("slice", func(t *testing.T) {
		acc1 := New(map[string]any{"a": 1, "b": 2, "c": []any{4, 5}})
		assert.Equal(t, 1, acc1.Int("a"))
		assert.Equal(t, 2, acc1.Int("b"))
		assert.Equal(t, 4, acc1.Int("c.0"))
		assert.Equal(t, "5", acc1.String("c.1"))
		assert.Equal(t, 5.0, acc1.Float64("c.1"))
		assert.Equal(t, -1., acc1.Float64("c.-1", -1.))
	})

	t.Run("nested.struct", func(t *testing.T) {
		type nested struct {
			D int
			E string
		}
		acc1 := New(map[string]any{"a": 1, "b": 2, "c": nested{4, "5"}})
		assert.Equal(t, 1, acc1.Int("a"))
		assert.Equal(t, 2, acc1.Int("b"))
		assert.Equal(t, 4, acc1.Int("c.D"))
		assert.Equal(t, "5", acc1.String("c.E"))
		assert.Nil(t, acc1.Get("c.Nil"))
	})
}
